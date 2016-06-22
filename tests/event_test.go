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
	"github.com/bep/gr/evt"
	"github.com/bep/gr/tests/grt"
)

func TestClickableButton(t *testing.T) {

	clickCount := 0
	clickListener := func(event *gr.Event) {
		clickCount++
	}

	button := el.Button(
		gr.Text("Clickable Button"),
		evt.Click(clickListener),
	)

	component := gr.NewSimpleComponent(button)
	elem := component.CreateElement(nil)
	tree := grt.ShallowRender(elem)

	for i := 0; i < 42; i++ {
		tree.CallEventListener("onClick")
	}

	grt.Equal(t, "<button onClick={function()}>Clickable Button</button>", tree.String())
	grt.Equal(t, 42, clickCount)

}
