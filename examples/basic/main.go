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

func main() {
	var (
		start     = time.Now().Unix()
		component = gr.New(new(elapser))
	)

	gr.RenderLoop(func() {
		props := gr.Props{"elapsed": (time.Now().Unix() - start)}
		component.Render("react", props)
	})
}

type elapser int

// Implements the Renderer interface.
func (e elapser) Render(this *gr.This) gr.Component {
	elapsed := this.Props()["elapsed"]
	message := fmt.Sprintf("React has been successfully running for '%v' seconds.", elapsed)

	elem := examples.Alert("info", el.Strong(gr.Text(message)))

	return examples.Example("Basic", elem)
}

// Implements the ShouldComponentUpdate interface.
func (e elapser) ShouldComponentUpdate(this *gr.This, nextProps gr.Props, nextState gr.State) bool {
	return this.Props().HasChanged(nextProps, "elapsed")
}
