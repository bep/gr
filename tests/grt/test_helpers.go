package grt

import (
	"fmt"
	"reflect"
	"strings"
	"testing"

	"github.com/bep/gr"
	"github.com/gopherjs/gopherjs/js"
)

func Equal(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		// TODO(bep) figure a way to get caller info
		t.Errorf("Assert mismatch:\n%v\n%v", expected, actual)
	}
}

func NotNil(t *testing.T, val interface{}) {
	if isNil(val) {
		Fail(t, fmt.Sprintf("Got <nil> for %T", val))
	}
	if val == js.Undefined {
		Fail(t, fmt.Sprintf("Got undefined for %T", val))
	}
}

func Fail(t *testing.T, args ...interface{}) {
	t.Fatal(args)
}

func ShallowRenderWithContext(c gr.Component, ctx gr.Context) *RenderedTree {
	if _, ok := c.(gr.Factory); ok {
		panic("Cannot render factories, create an Element first")
	}
	tree := sd.Call("shallowRender", c.Node(), ctx)
	return &RenderedTree{Object: tree, context: ctx}
}

func ShallowRender(c gr.Component) *RenderedTree {
	return ShallowRenderWithContext(c, gr.Context{})
}

func (t *RenderedTree) ReRender(props gr.Props) {
	t.reRender(props, t.context)
}

func (t *RenderedTree) Dive(path ...string) *RenderedTree {
	tree := t.dive(path, t.context)
	return &RenderedTree{Object: tree, context: t.context}
}

// TODO(bep)
// Move this to its own repo maybe when it is more mature.
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

type Props map[string]interface{}

type Matcher struct {
	key, value string
}

func (p Props) CallEventListener(name string, args ...interface{}) *js.Object {
	return p[name].(func(...interface{}) *js.Object)(name, args)
}

func NewMatcher(key, value string) Matcher {
	return Matcher{key: key, value: value}
}

var renderStringReplacers = strings.NewReplacer("  ", "", "\n", "")

func (t *RenderedTree) String() string {
	s := t.toString()
	return renderStringReplacers.Replace(s)
}

func (t *RenderedTree) Sub(selector string, matchers ...Matcher) *RenderedTree {

	m := make(map[string]interface{})

	for _, matcher := range matchers {
		m[matcher.key] = matcher.value
	}

	var args []interface{} = []interface{}{selector}

	if len(m) > 0 {
		args = append(args, m)
	}

	subTree := t.Call("subTree", args)

	if !subTree.Bool() {
		return nil
	}

	return &RenderedTree{Object: subTree}
}

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
