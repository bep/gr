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

	clickListener := func(this *gr.This, event *gr.Event) {
		this.Props().Call("propsFunc", 2)
		this.Props().Func("propsFunc")(2)

	}

	button := el.Button(
		gr.Text("Clickable Button"),
		evt.Click(clickListener),
	)

	component := gr.NewSimpleComponent(button)
	elem := component.CreateElement(gr.Props{"propsFunc": propsFunc})
	tree := grt.ShallowRender(elem)

	for i := 0; i < 2; i++ {
		tree.Props.CallEventListener("onClick")
	}

	grt.Equal(t, "<button onClick={function()}>Clickable Button</button>", tree.String())
	grt.Equal(t, 8, clickVal)

}

func TestPropsChildren(t *testing.T) {
	buttons := make([]gr.Component, 3)

	for i := 0; i < len(buttons); i++ {
		buttons[i] = el.Button(gr.Text(fmt.Sprintf("B%d", i)))
	}

	comp := gr.New(gr.NewRenderer(func(this *gr.This) gr.Component {
		return el.Div(this.Children().Element())
	}))

	elem := comp.CreateElement(nil, buttons...)
	tree := grt.ShallowRender(elem)

	grt.Equal(t, "<div><button>B0</button><button>B1</button><button>B2</button></div>", tree.String())

}
