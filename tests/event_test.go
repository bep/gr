package tests

import (
	"testing"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/evt"
	"github.com/bep/gr/tests/grt"
)

func TestClickableButton(t *testing.T) {

	clickCount := 0
	clickListener := func(this *gr.This, event *gr.Event) {
		clickCount++
	}

	button := el.Button(
		gr.Text("Clickable Button"),
		evt.Click(clickListener),
	)

	component := gr.NewSimpleComponent(button)
	elem := component.Create(nil)
	tree := grt.ShallowRender(elem)

	for i := 0; i < 42; i++ {
		tree.Props.CallEventListener("onClick")
	}

	grt.Equal(t, "<button onClick={function()}>Clickable Button</button>", tree.String())
	grt.Equal(t, 42, clickCount)

}
