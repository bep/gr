package tests

import (
	"testing"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/gopherjs/gopherjs/js"
)

func TestRenderButton(t *testing.T) {
	button := el.Button(gr.Text("Shiny Button"))
	tree := shallowRender(button)
	assertEqual(t, "button", tree.Type)
	assertEqual(t, "Shiny Button", tree.Text())
}

//
// Utility helpers below
//

func assertEqual(t *testing.T, expected, actual interface{}) {
	if expected != actual {
		t.Errorf("Assert mismatch:\n%v\n%v", expected, actual)
	}
}

func shallowRender(e *gr.Element) *RenderedTree {
	tree := sd.Call("shallowRender", e.Node())
	println(tree)
	return &RenderedTree{Object: tree}
}

// TODO(bep)
// Move this to its own repo maybe when it is more mature.
type RenderedTree struct {
	*js.Object
	// reRender: [Function],
	// getMountedInstance: [Function],
	// subTree: [Function],
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
	//toString: [Function],
	//props: [Getter],
	Type string `js:"type"`
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
