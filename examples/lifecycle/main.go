package main

import (
	"time"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/examples"
)

// Simple example demonstrating the avilable lifecycle hooks.
// See also https://facebook.github.io/react/docs/component-specs.html

func main() {

	var (
		component = gr.New(new(lifecycle))
		props     = gr.Props{}
		counter   = 1
	)

	quit := gr.RenderLoop(func() {
		// Only update now and then.
		props["prop"] = counter%100 == 0
		counter++
		component.Render("react", props)
	})

	time.Sleep(10 * time.Second)

	// Stop the render loop
	close(quit)

	if !gr.UnmountComponentAtNode("react") {
		panic("Unmount failed")
	}

	time.Sleep(10 * time.Second)
}

type lifecycle int

// Implements the Renderer interface.
func (l lifecycle) Render(this *gr.This) gr.Component {
	println("Render")
	elem := el.Div(el.Header2(
		gr.Text("Look at the lifecycle events in your console."),
		gr.Style("color", this.State()["color"])))

	return examples.Example("Lifecycle", elem)
}

// Implements the StateInitializer interface.
func (l lifecycle) GetInitialState(this *gr.This) gr.State {
	println("GetInitialState")
	return gr.State{"color": "#ffcc00"}
}

// Implements the ShouldComponentUpdate interface
func (l lifecycle) ShouldComponentUpdate(this *gr.This, nextProps gr.Props, nextState gr.State) bool {
	return this.Props().HasChanged(nextProps, "prop")
}

// Implements the ComponentWillUpdate interface
func (l lifecycle) ComponentWillUpdate(this *gr.This, nextProps gr.Props, nextState gr.State) {
	println("ComponentWillUpdate")
}

// Implements the ComponentWillReceiveProps interface
func (l lifecycle) ComponentWillReceiveProps(this *gr.This, p gr.Props) {
	println("ComponentWillReceiveProps")
}

// Implements the ComponentDidUpdate interface
func (l lifecycle) ComponentDidUpdate(this *gr.This, props gr.Props, state gr.State) {
	println("ComponentDidUpdate")
}

// Implements the ComponentWillMount interface
func (l lifecycle) ComponentWillMount(this *gr.This) {
	println("ComponentWillMount")
}

// Implements the ComponentDidMount interface
func (l lifecycle) ComponentDidMount(this *gr.This) {
	println("ComponentDidMount")
}

// Implements the ComponentWillUnmount interface
func (l lifecycle) ComponentWillUnmount(this *gr.This) {
	println("ComponentWillUnmount")
}
