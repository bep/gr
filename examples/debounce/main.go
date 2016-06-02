package main

import (
	"time"

	"fmt"

	"github.com/bep/debounce"
	"github.com/bep/gr"
	"github.com/bep/gr/attr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/evt"
	"github.com/bep/gr/examples"
)

func main() {
	gr.New(new(mouseTracker)).Render("react", gr.Props{})

}

type mouseTracker struct {
	*gr.This
}

// Implements the Renderer interface.
func (m mouseTracker) Render() gr.Component {
	state := m.State()
	x := state["mouseX"]
	y := state["mouseY"]
	counter := state["counter"]

	message := fmt.Sprintf("X: %v / Y:%v / %v", x, y, counter)

	elem := el.Div(
		gr.Style("backgroundColor", "#3399ff"),
		gr.Style("width", "500px"),
		gr.Style("height", "500px"),
		gr.Style("padding", "30px"),
		el.Header1(
			gr.Style("color", "white"),
			gr.Text(message),
		),
		el.Header2(
			gr.Style("color", "white"),
			gr.Text("Move the mouse around in the blue square, then pause ..."),
		),
		el.Paragraph(el.Anchor(
			attr.HRef("https://davidwalsh.name/javascript-debounce-function"),
			gr.Text("Debounce Function Explained"))),
		evt.MouseMove(debounceMouseListener),
	)

	return examples.Example("Debounce", elem)
}

func (m mouseTracker) GetInitialState() gr.State {
	return gr.State{"mouseX": 0, "mouseY": 0, "counter": 0}
}

func (m mouseTracker) ShouldComponentUpdate(next gr.Cops) bool {
	return m.State().HasChanged(next.State, "mouseX", "mouseY")
}

var (
	// Only update the UI when no new events have been received for >= 200 ms.
	debouncer, _                      = debounce.New(200 * time.Millisecond)
	debounceMouseListener gr.Listener = func(e *gr.Event) {

		// React recycles events - so extract early.
		// For the mouse API, see https://developer.mozilla.org/en-US/docs/Web/API/MouseEvent
		// Note the browser compatibility notes.
		clientX := e.Int("screenX")
		clientY := e.Int("screenY")

		f := func() {
			counter := e.This.State().Int("counter") + 1
			e.This.SetState(gr.State{"mouseX": clientX, "mouseY": clientY, "counter": counter})
		}

		debouncer(f)

	}
)
