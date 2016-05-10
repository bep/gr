// Package grt contains utilities used to test React components.
// TODO(bep)
// Move this to its own repo maybe when it is more mature.
package grt

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/bep/gr"
	"github.com/gopherjs/gopherjs/js"
)

// Equal is a test assertion used to check equality.
func Equal(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		// TODO(bep) figure a way to get caller info
		t.Errorf("Assert mismatch:\n%v\n%v", expected, actual)
	}
}

// NotNil is a test assertion that checks for both nil values and js.Undefined.
func NotNil(t *testing.T, val interface{}) {
	if isNil(val) {
		Fail(t, fmt.Sprintf("Got <nil> for %T", val))
	}
	if val == js.Undefined {
		Fail(t, fmt.Sprintf("Got undefined for %T", val))
	}
}

// Fail fails the test with the given message.
func Fail(t *testing.T, args ...interface{}) {
	t.Fatal(args)
}

// ShallowRenderWithContext performs a shallow render with the given context.
func ShallowRenderWithContext(c gr.Component, ctx gr.Context) *RenderedTree {
	if _, ok := c.(gr.Factory); ok {
		panic("Cannot render factories, create an Element first")
	}
	tree := sd.Call("shallowRender", c.Node(), ctx)
	return &RenderedTree{Object: tree, context: ctx}
}

// ShallowRender performs a shallow render of the given component.
func ShallowRender(c gr.Component) *RenderedTree {
	return ShallowRenderWithContext(c, gr.Context{})
}

// ReRender rerenders a component with new properties.
func (t *RenderedTree) ReRender(props gr.Props) {
	t.reRender(props, t.context)
}

// Dive can be used to render a sub-component.
func (t *RenderedTree) Dive(path ...string) *RenderedTree {
	tree := t.dive(path, t.context)
	return &RenderedTree{Object: tree, context: t.context}
}

// RenderedTree represents a shallow render.
type RenderedTree struct {
	*js.Object
	reRender           func(gr.Props, gr.Context)                      `js:"reRender"`
	GetMountedInstance func() *js.Object                               `js:"getMountedInstance"`
	subTree            func(string, map[string]interface{}) *js.Object `js:"subTree"`
	dive               func([]string, gr.Context) *js.Object           `js:"dive"`
	Text               func() string                                   `js:"text"`
	// fillField: [Function],
	//findComponent: [Function],
	//findComponentLike: [Function],
	GetRenderOutput func() map[string]interface{} `js:"getRenderOutput"`
	toString        func() string                 `js:"toString"`
	Props           Props                         `js:"props"`
	Type            string                        `js:"type"`

	context gr.Context
}

// Props represents the properties set on a React element.
// This is the same as gr.Props, but reimplemented here so we can add
// test methods to it.
type Props map[string]interface{}

// Matcher used to find a component in the component tree.
type Matcher struct {
	key, value string
}

// CallEventListener is a convenience func to simulate button clicks etc.
// by calling the listener methods by name.
func (p Props) CallEventListener(name string, args ...interface{}) *js.Object {
	return p[name].(func(...interface{}) *js.Object)(name, args)
}

// NewMatcher creates a new matcher.
func NewMatcher(key, value string) Matcher {
	return Matcher{key: key, value: value}
}

var renderStringReplacers = strings.NewReplacer("  ", "", "\n", "")

// String represents a shallow rendered component.
// This is in heavy use in assertions, so we try to keep it stable.
func (t *RenderedTree) String() string {
	s := t.toString()
	return renderStringReplacers.Replace(s)
}

// Sub returns a sub component matching the matchers.
func (t *RenderedTree) Sub(selector string, matchers ...Matcher) *RenderedTree {

	m := make(map[string]interface{})

	for _, matcher := range matchers {
		m[matcher.key] = matcher.value
	}

	var args = []interface{}{selector}

	if len(m) > 0 {
		args = append(args, m)
	}

	subTree := t.Call("subTree", args)

	if !subTree.Bool() {
		return nil
	}

	return &RenderedTree{Object: subTree}
}

// This returns the this context of a shallow render.
func (t *RenderedTree) This() *gr.This {
	return gr.NewThis(t.GetMountedInstance())
}

var (
	react *js.Object
	sd    *js.Object
)

func init() {

	react = js.Global.Get("React")

	if react == js.Undefined {
		panic("Facebook React not found, make sure it is loaded.")
	}

	// Skin deep
	sd = js.Global.Get("sd")

	if sd == js.Undefined {
		panic("Skin deep not found, make sure it is loaded.")
	}

}

// Code below copied from testify, MIT licensed
// Copyright (c) 2012 - 2013 Mat Ryer and Tyler Bunnell
// https://github.com/stretchr/testify/blob/master/LICENSE
func isNil(object interface{}) bool {
	if object == nil {
		return true
	}

	value := reflect.ValueOf(object)
	kind := value.Kind()
	if kind >= reflect.Chan && kind <= reflect.Slice && value.IsNil() {
		return true
	}

	return false
}
