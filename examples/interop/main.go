package main

import (
	"fmt"
	"time"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/examples"
)

func main() {
	var (
		start             = time.Now().Unix()
		internalComponent = gr.New(new(elapser))
	)

	gr.RenderLoop(func() {
		props := gr.Props{"elapsed": (time.Now().Unix() - start)}
		internalComponent.Render("react", props)
	})
}

type elapser int

var (
	externalComponent = gr.FromJS("ElapserExt")
)

// Implements the Renderer interface.
func (e elapser) Render(this *gr.This) gr.Component {
	elapsed := this.Props()["elapsed"]
	message := fmt.Sprintf("Go Timer has been successfully running for %v seconds.", elapsed)

	internalCounter := examples.Alert("info", el.Strong(gr.Text(message)))
	externalCounter := examples.Alert("warning", externalComponent.CreateElement(this.Props()))

	return examples.Example("Interop", internalCounter, externalCounter)
}

// Implements the ShouldComponentUpdate interface.
func (e elapser) ShouldComponentUpdate(this *gr.This, nextProps gr.Props, nextState gr.State) bool {
	return this.Props().HasChanged(nextProps, "elapsed")
}
