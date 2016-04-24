package main

import (
	"fmt"
	"time"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/examples"
)

// Basic example similar to:
// https://github.com/facebook/react/blob/master/examples/basic/index.html

type elapser int

// Implements the Renderer interface.
func (e elapser) Render(this *gr.This) gr.Component {
	elapsed := this.Props["elapsed"]
	message := fmt.Sprintf("React has been successfully running for '%v' seconds.", elapsed)

	return examples.Alert("success", el.Strong(gr.Text(message)))
}

// Implements the ShouldComponentUpdate interface.
func (e elapser) ShouldComponentUpdate(this *gr.This, nextProps, p gr.Props) bool {
	return this.Props.HasChanged(nextProps, "elapsed")
}

func main() {

	start := time.Now().Unix()
	comp := gr.NewRoot(new(elapser))

	gr.RenderLoop(func() {
		props := gr.Props{"elapsed": (time.Now().Unix() - start)}
		comp.Render("react", props)
	})
}
