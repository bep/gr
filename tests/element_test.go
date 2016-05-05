package tests

import (
	"testing"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/tests/grt"
)

func TestRenderButton(t *testing.T) {
	button := el.Button(gr.Text("Shiny Button"))
	tree := grt.ShallowRender(button)

	grt.Equal(t, "<button>Shiny Button</button>", tree.String())
}

func TestRenderNestedSimple(t *testing.T) {
	div := el.Div(
		el.Button(gr.Text("Button in Div")),
	)
	tree := grt.ShallowRender(div)

	grt.Equal(t, "<div><button>Button in Div</button></div>", tree.String())
}

func TestRenderNestedComplex(t *testing.T) {
	div := el.Div(
		el.Div(
			el.Paragraph(gr.Text("P1")),
		),
		el.Div(
			el.Paragraph(gr.Text("P2")),
			el.Div(
				el.Paragraph(gr.Text("P3")),
				el.Div(),
			),
		),
		el.Div(),
	)
	tree := grt.ShallowRender(div)

	grt.Equal(t,
		"<div><div><p>P1</p></div><div><p>P2</p><div><p>P3</p><div /></div></div><div /></div>",
		tree.String())
}
