package main

import (
	"fmt"
	"time"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/examples"
)

var (
	externalGlobalComponent = gr.FromGlobal("ElapserExtGlobal")
	externalModuleComponent = gr.Require("ElapserExtModule")
)

func main() {
	var (
		start = time.Now().Unix()

		// Set it in module exports so it can be used from JavaScript via require().
		// Also set it as a JavaScript global.
		internalComponent = gr.New(new(elapser), gr.Export("Elapser"), gr.Global("ElapserGlobal"))
	)

	gr.RenderLoop(func() {
		props := gr.Props{"title": "Interop from Go Global/Module", "elapsed": (time.Now().Unix() - start)}
		internalComponent.Render("react", props)
	})
}

type elapser int

// Implements the Renderer interface.
func (e elapser) Render(this *gr.This) gr.Component {
	elapsed := this.Props()["elapsed"]
	title := this.Props()["title"].(string) // TODO(bep cleanup props handling

	message := fmt.Sprintf("Go Timer has been successfully running for %v seconds.", elapsed)

	internalCounter := examples.Alert("info", el.Strong(gr.Text(message)))
	externalGlobalCounter := examples.Alert("warning", externalGlobalComponent.CreateElement(this.Props()))
	externalModuleCounter := examples.Alert("danger", externalModuleComponent.CreateElement(this.Props()))

	return examples.Panel(title, internalCounter, externalGlobalCounter, externalModuleCounter)
}

// Implements the ShouldComponentUpdate interface.
func (e elapser) ShouldComponentUpdate(this *gr.This, nextProps gr.Props, nextState gr.State) bool {
	return this.Props().HasChanged(nextProps, "elapsed")
}
