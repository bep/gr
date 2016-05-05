package tests

import (
	"testing"

	"github.com/bep/gr"
	"github.com/bep/gr/attr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/tests/grt"
)

func TestRenderWithAttribute(t *testing.T) {
	img := el.Image(
		attr.Alt("Shiny Image"),
		attr.Src("image.jpg"),
	)

	tree := grt.ShallowRender(img)

	grt.Equal(t, `<img alt="Shiny Image" src="image.jpg" />`, tree.String())
}

func TestDiscard(t *testing.T) {
	img := el.Image(
		gr.Discard,
		attr.Src("image.jpg"),
		gr.Discard,
	)

	tree := grt.ShallowRender(img)

	grt.Equal(t, `<img src="image.jpg" />`, tree.String())
}

func TestRenderWithCSS(t *testing.T) {

	for _, test := range []struct {
		element *gr.Element
		expect  string
	}{
		{el.Paragraph(gr.Text("GopherJS is pretty cool!"), gr.CSS("important")),
			`<p className="important">GopherJS is pretty cool!</p>`},
		{el.Header1(gr.CSS("important", "headline")),
			`<h1 className="important headline" />`},
		{el.Header2(gr.CSS("c1", "c2"), gr.CSS("c3"), gr.CSS("c2")),
			`<h2 className="c1 c2 c3 c2" />`},
		{el.Header2(gr.CSS("c1", "c2"), gr.CSS("c3 c4")),
			`<h2 className="c1 c2 c3 c4" />`},
	} {
		tree := grt.ShallowRender(test.element)
		grt.Equal(t, test.expect, tree.String())
	}
}

func TestRenderWithStyle(t *testing.T) {

	for _, test := range []struct {
		element *gr.Element
		expect  string
	}{
		{el.Header2(gr.Style("color", "blue")),
			`<h2 style={{"color": "blue"}} />`},
		{el.Header2(gr.Style("color", "blue"), gr.Style("color", "red")),
			`<h2 style={{"color": "red"}} />`},
		{el.Paragraph(gr.Style("color", "red"), gr.Style("fontWeight", "bold")),
			`<p style={{"color": "red", "fontWeight": "bold"}} />`},
	} {
		tree := grt.ShallowRender(test.element)
		grt.Equal(t, test.expect, tree.String())
	}
}

func TestRenderWithAria(t *testing.T) {
	for _, test := range []struct {
		element *gr.Element
		expect  string
	}{
		{el.Header2(gr.Aria("a1", "a1-v")),
			`<h2 aria-a1="a1-v" />`},
		{el.Header2(gr.Aria("a1", "a1-v"), gr.Aria("a2", "a2-v"), gr.Aria("a3", "a3-v")),
			`<h2 aria-a1="a1-v" aria-a2="a2-v" aria-a3="a3-v" />`},
	} {
		tree := grt.ShallowRender(test.element)
		grt.Equal(t, test.expect, tree.String())
	}
}

func TestRenderWithData(t *testing.T) {
	for _, test := range []struct {
		element *gr.Element
		expect  string
	}{
		{el.Header2(gr.Data("d1", "d1-v")),
			`<h2 data-d1="d1-v" />`},
		{el.Header2(gr.Data("d1", "d1-v"), gr.Data("d2", "d2-v"), gr.Data("d3", "d3-v")),
			`<h2 data-d1="d1-v" data-d2="d2-v" data-d3="d3-v" />`},
	} {
		tree := grt.ShallowRender(test.element)
		grt.Equal(t, test.expect, tree.String())
	}
}

func TestRenderWithProp(t *testing.T) {
	img := el.Image(
		attr.Src("image.jpg"),
		gr.Prop("cp1", "p1"),
		gr.Prop("cp2", "p2"),
	)

	tree := grt.ShallowRender(img)

	grt.Equal(t, `<img src="image.jpg" cp1="p1" cp2="p2" />`, tree.String())

}
