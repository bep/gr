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
	tree := grt.Shallow(component)

	tree.Props.CallEventListener("onClick")
	tree.Props.CallEventListener("onClick")

	grt.Equal(t, "<button>Clickable Button</button>", tree.String())
	grt.Equal(t, 2, clickCount)

}
