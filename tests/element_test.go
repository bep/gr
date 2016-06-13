/*
Copyright 2016 Bjørn Erik Pedersen <bjorn.erik.pedersen@gmail.com> All rights reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

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

func TestRenderRegularTextLast(t *testing.T) {
	div := el.Div(
		el.Bold(gr.Text("Bold")),
		el.Italic(gr.Text("Italic")),
		gr.Text("Regular"),
	)
	tree := grt.ShallowRender(div)

	grt.Equal(t, "<div><b>Bold</b><i>Italic</i>Regular</div>", tree.String())
}
