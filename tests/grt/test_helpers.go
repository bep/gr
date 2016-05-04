package grt

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/bep/gr"
	"github.com/gopherjs/gopherjs/js"
)

func Equal(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Assert mismatch:\n%v\n%v", expected, actual)
	}
}

func NotNil(t *testing.T, val interface{}) {
	if isNil(val) {
		Fail(t, fmt.Sprintf("Got <nil> for %T", val))
	}
}

func Fail(t *testing.T, args ...interface{}) {
	t.Fatal(args)
}

func Shallow(c gr.Component) *RenderedTree {
	tree := sd.Call("shallowRender", c.Node())
	return &RenderedTree{Object: tree}
}

// TODO(bep)
// Move this to its own repo maybe when it is more mature.
type RenderedTree struct {
	*js.Object
	// reRender: [Function],
	// getMountedInstance: [Function],
	subTree func(string, map[string]interface{}) *js.Object `js:"subTree"`
	// subTreeLike: [Function],
	// everySubTree: [Function],
	// everySubTreeLike: [Function],
	// dive: [Function],
	// findNode: [Function],
	// textIn: [Function],
	Text func() string `js:"text"`
	// fillField: [Function],
	//findComponent: [Function],
	//findComponentLike: [Function],
	GetRenderOutput func() map[string]interface{} `js:"getRenderOutput"`
	String          func() string                 `js:"toString"`
	Props           Props                         `js:"props"`
	Type            string                        `js:"type"`
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

func (tree *RenderedTree) HasAllProps(t *testing.T, matchers ...Matcher) {
	for _, m := range matchers {
		found := false
		for key, val := range tree.Props {
			if m.key == key && m.value == val {
				found = true
				break
			}
		}
		if !found {
			t.Fatalf("Prop %q/%q not found", m.key, m.value)
		}
	}
}

// Copied from testify, MIT licensed
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
