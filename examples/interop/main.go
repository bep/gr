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

type elapser struct {
	*gr.This
}

// Implements the Renderer interface.
func (e elapser) Render() gr.Component {
	elapsed := e.Props().String("elapsed")
	title := e.Props().String("title")
	message := fmt.Sprintf("Go Timer has been successfully running for %v seconds.", elapsed)

	internalCounter := examples.Alert("info", el.Strong(gr.Text(message)))
	externalGlobalCounter := examples.Alert("warning", externalGlobalComponent.CreateElement(e.Props()))
	externalModuleCounter := examples.Alert("danger", externalModuleComponent.CreateElement(e.Props()))

	return examples.Panel(title, internalCounter, externalGlobalCounter, externalModuleCounter)
}

// Implements the ShouldComponentUpdate interface.
func (e elapser) ShouldComponentUpdate(next gr.Cops) bool {
	return e.Props().HasChanged(next.Props, "elapsed")
}
