package main

import (
	"fmt"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/evt"
)

// Basic click counter example similar to:
// https://github.com/facebook/react/blob/master/examples/basic-click-counter/index.html

type clickCounter int

// Implements the StateInitializer interface.
func (c clickCounter) GetInitialState(this *gr.This) gr.State {
	return gr.State{"counter": 0}
}

// Implements the Renderer interface.
func (c clickCounter) Render(this *gr.This) gr.Component {
	counter := this.State()["counter"]

	message := fmt.Sprintf(" Click me! Number of clicks: %v", counter)

	return el.Div(
		el.Button(
			gr.CSS("btn", "btn-lg", "btn-primary"),
			gr.Text(message),
			evt.Click(c.onClick)))
}

func (c clickCounter) onClick(this *gr.This, event *gr.Event) {
	this.SetState(gr.State{"counter": this.StateInt("counter") + 1})
}

// Implements the ShouldComponentUpdate interface.
func (e clickCounter) ShouldComponentUpdate(this *gr.This, nextProps gr.Props, nextState gr.State) bool {
	return this.State().HasChanged(nextState, "counter")
}

func main() {
	component := gr.New(new(clickCounter))

	gr.RenderLoop(func() {
		component.Render("react", gr.Props{})
	})
}
