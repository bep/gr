package examples

import (
	"fmt"
	"strings"

	"github.com/bep/gr"
	"github.com/bep/gr/attr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/evt"
)

// Example is a wrapper for the examples
func Example(title string, mods ...gr.Modifier) *gr.Element {

	links := el.Div(gr.CSS("list-group"),
		exampleListItem(title, "basic", "Basic"),
		exampleListItem(title, "basic-click-counter", "Basic Click-counter"),
		exampleListItem(title, "composition", "Component Composition"),
		exampleListItem(title, "lifecycle", "Lifecycle"),
	)

	elem := el.Div(gr.CSS("panel", "panel-primary"),
		el.Div(gr.CSS("panel-heading"), el.Header1(gr.Text(title))),
		el.Div(append(mods, gr.CSS("panel-body"),
			el.Header3(gr.Text("More examples")), links)...),
		el.Div(gr.CSS("panel-footer"),
			el.Div(
				el.Emphasis(gr.Text("Facebook React in Go: ")),
				el.Anchor(attr.HRef("https://github.com/bep/gr/"),
					gr.Text("https://github.com/bep/gr/")))))

	return elem
}

func exampleListItem(title, href, text string) gr.Modifier {
	var (
		itemStatus gr.Modifier = gr.Discard
		loc                    = gr.Location()
	)

	if !strings.HasSuffix(href, "/") {
		href += "/"
	}

	if strings.HasSuffix(loc.Path, href) {
		itemStatus = gr.CSS("active")
	}

	href = "../" + href

	return el.Anchor(gr.CSS("list-group-item"), itemStatus, attr.HRef(href), gr.Text(text))

}

// Alert creates a Bootstrap alert element.
func Alert(classifier string, body gr.Modifier) *gr.Element {
	e := el.Div(
		gr.CSS("alert", "alert-"+classifier),
		el.Anchor(attr.HRef("#"),
			gr.CSS("close"), gr.Data("dismiss", "alert"), gr.Aria("label", "close"),
			gr.Text("Close")),
		body)
	return e
}

// Some reusable components to use in composition examples.
// This is just copy-paste from the click counter example. Consider making something else.
type ClickCounter int

// Implements the StateInitializer interface.
func (c ClickCounter) GetInitialState(this *gr.This) gr.State {
	return gr.State{"counter": 0}
}

// Implements the Renderer interface.
func (c ClickCounter) Render(this *gr.This) gr.Component {
	counter := this.State()["counter"]
	message := fmt.Sprintf(" Click me! Number of clicks: %v", counter)

	return el.Div(
		el.Button(
			gr.CSS("btn", "btn-lg", "btn-primary"),
			gr.Text(message),
			evt.Click(c.onClick)))
}

func (c ClickCounter) onClick(this *gr.This, event *gr.Event) {
	this.SetState(gr.State{"counter": this.StateInt("counter") + 1})
}

func (c ClickCounter) ShouldComponentUpdate(
	this *gr.This, nextProps gr.Props, nextState gr.State) bool {

	return this.State().HasChanged(nextState, "counter")
}

func (c ClickCounter) ComponentDidMount(this *gr.This) {
	println("ClickCounter: ComponentDidMount")
}
