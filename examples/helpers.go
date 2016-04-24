package examples

import (
	"github.com/bep/gr"
	"github.com/bep/gr/el"
)

// Alert creates a Bootstrap alert element.
func Alert(classifier string, body gr.Modifier) *gr.Element {
	e := el.Div(
		gr.CSS("alert", "alert-success"),
		el.Anchor(gr.Prop("href", "#"),
			gr.CSS("close"), gr.Data("dismiss", "alert"), gr.Aria("label", "close"),
			gr.Text("Close")),
		body)
	return e
}
