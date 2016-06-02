package main

import (
	"fmt"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/evt"
	"github.com/bep/gr/examples"
)

// Basic click counter example similar to:
// https://github.com/facebook/react/blob/master/examples/basic-click-counter/index.html

func main() {
	component := gr.New(new(clickCounter))

	gr.RenderLoop(func() {
		component.Render("react", gr.Props{})
	})
}

type clickCounter struct {
	*gr.This
}

// Implements the StateInitializer interface.
func (c clickCounter) GetInitialState() gr.State {
	return gr.State{"counter": 0}
}

// Implements the Renderer interface.
func (c clickCounter) Render() gr.Component {
	counter := c.State()["counter"]
	message := fmt.Sprintf(" Click me! Number of clicks: %v", counter)

	elem := el.Div(
		el.Button(
			gr.CSS("btn", "btn-lg", "btn-primary"),
			gr.Style("color", "orange"),
			gr.Text(message),
			evt.Click(c.onClick)))

	return examples.Example("Click Counter", elem)
}

func (c clickCounter) onClick(event *gr.Event) {
	c.SetState(gr.State{"counter": c.State().Int("counter") + 1})
}

// Implements the ShouldComponentUpdate interface.
func (c clickCounter) ShouldComponentUpdate(next gr.Cops) bool {
	return c.State().HasChanged(next.State, "counter")
}
