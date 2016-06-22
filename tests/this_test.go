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

	"fmt"
	"sync"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/evt"
	"github.com/bep/gr/tests/grt"
)

func TestPropsFunc(t *testing.T) {
	clickVal := 0

	propsFunc := func(val int) {
		clickVal += val
	}

	cb := &testClickableButtons{}
	component := gr.New(cb)
	elem := component.CreateElement(gr.Props{"propsFunc": propsFunc})

	tree := grt.ShallowRender(elem)

	firstButton := tree.Dive("p", "button")

	// 2 x 3
	for i := 0; i < 2; i++ {
		firstButton.CallEventListener("onClick")
	}

	secondButton := tree.Dive("footer", "button")

	// 3 x 8
	for i := 0; i < 3; i++ {
		secondButton.CallEventListener("onClick")
	}

	grt.Equal(t, "<div><p><button onClick={function()}>Clickable Button 1</button></p><footer><button onClick={function()}>Clickable Button 2</button></footer></div>", tree.String())
	grt.Equal(t, 30, clickVal)

}

func TestPropsChildren(t *testing.T) {
	buttons := make([]gr.Component, 3)

	for i := 0; i < len(buttons); i++ {
		buttons[i] = el.Button(gr.Text(fmt.Sprintf("B%d", i)))
	}

	comp := gr.New(&thisCompChildren{})

	elem := comp.CreateElement(nil, buttons...)
	tree := grt.ShallowRender(elem)

	grt.Equal(t, "<div><button>B0</button><button>B1</button><button>B2</button></div>", tree.String())

}

func TestBindThis(t *testing.T) {

	compPtr := gr.New(&thisCompEmbed{})
	compNamed := gr.New(&thisCompNamed{})

	grt.Equal(t, "<div>embed</div>", grt.ShallowRender(compPtr.CreateElement(gr.Props{"id": "embed"})).String())
	grt.Equal(t, "<div>named</div>", grt.ShallowRender(compNamed.CreateElement(gr.Props{"id": "named"})).String())
}

func TestBindThisVariations(t *testing.T) {

	var wg sync.WaitGroup
	tc := &thisCompEmbed{}
	comp := gr.New(tc)

	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func(ii int) {
			defer wg.Done()
			props := gr.Props{}

			if ii%2 == 0 {
				props = gr.Props{"id": fmt.Sprintf("%d", ii), "idFunc": tc.ID}
			}
			if ii > 0 && ii%4 == 0 {
				props = gr.Props{"id": fmt.Sprintf("%d", ii), "foo": "bar"}
			}
			elem := comp.CreateElement(props)
			tree := grt.ShallowRender(elem)
			expect := "<div></div>"
			if ii > 0 && ii%4 == 0 {
				expect = fmt.Sprintf("<div>%dbar</div>", ii)
			} else if ii%2 == 0 {
				expect = fmt.Sprintf("<div>%d%d</div>", ii, ii)
			}

			grt.Equal(t, expect, tree.String())
		}(i)
	}

	wg.Wait()

}

type thisCompEmbed struct {
	*gr.This
}

func (c *thisCompEmbed) Render() gr.Component {
	thisStr := c.Props().String("id")
	thisStr += c.Props().String("foo")
	if _, ok := c.Props()["idFunc"]; ok {
		thisStr += c.Props().Call("idFunc").String()
	}
	return el.Div(gr.Text(thisStr))
}

func (c *thisCompEmbed) ID() string {
	return c.Props().String("id")
}

type thisCompNamed struct {
	That *gr.This
}

func (c *thisCompNamed) Render() gr.Component {
	thisStr := fmt.Sprintf("%v", c.That.Props()["id"])
	return el.Div(gr.Text(thisStr))
}

type thisCompChildren struct {
	*gr.This
}

func (c *thisCompChildren) Render() gr.Component {
	return el.Div(c.This.Children().Element())
}

type testClickableButtons struct {
	*gr.This
}

func (c *testClickableButtons) Render() gr.Component {
	clickListener := func(event *gr.Event) {
		c.Props().Call("propsFunc", 3)

	}

	button1 := el.Button(
		gr.Text("Clickable Button 1"),
		evt.Click(clickListener),
	)

	button2 := el.Button(
		gr.Text("Clickable Button 2"),
		evt.Click(c.createClickListener(4)),
	)

	return el.Div(
		el.Paragraph(button1),
		el.Footer(button2),
	)
}

func (c *testClickableButtons) createClickListener(increment int) gr.Listener {
	return func(event *gr.Event) {
		event.This.Props().Call("propsFunc", increment)
		c.Props().Call("propsFunc", increment)
	}
}
