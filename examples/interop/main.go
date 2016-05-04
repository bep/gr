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
		start = time.Now().Unix()

		// Export it so it can be used from JavaScript.
		internalComponent = gr.New(new(elapser), gr.Export("Elapser"))
	)

	gr.RenderLoop(func() {
		props := gr.Props{"title": "Interop from Go", "elapsed": (time.Now().Unix() - start)}
		internalComponent.Render("react", props)
	})
}

type elapser int

var externalComponent = gr.FromJS("ElapserExt")

// Implements the Renderer interface.
func (e elapser) Render(this *gr.This) gr.Component {
	elapsed := this.Props()["elapsed"]
	title := this.Props()["title"].(string) // TODO(bep cleanup props handling

	message := fmt.Sprintf("Go Timer has been successfully running for %v seconds.", elapsed)

	internalCounter := examples.Alert("info", el.Strong(gr.Text(message)))
	externalCounter := examples.Alert("warning", externalComponent.Create(this.Props()))

	return examples.Example(title, internalCounter, externalCounter)
}

// Implements the ShouldComponentUpdate interface.
func (e elapser) ShouldComponentUpdate(this *gr.This, nextProps gr.Props, nextState gr.State) bool {
	return this.Props().HasChanged(nextProps, "elapsed")
}
