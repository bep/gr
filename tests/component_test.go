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
	"fmt"
	"sort"
	"testing"

	"time"

	"github.com/bep/gr"
	"github.com/bep/gr/attr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/tests/grt"
	"github.com/gopherjs/gopherjs/js"
)

const exportedTestComponent = "GrtTest"

func TestNewSimpleRenderer(t *testing.T) {
	button := el.Button(gr.Text("Shiny Button"))
	renderer := gr.NewSimpleRenderer(button)
	grt.Equal(t, button, renderer.Render())
}

func TestNewSimpleComponent(t *testing.T) {
	button := el.Button(gr.Text("Simple Button"))
	rc := gr.NewSimpleComponent(button)
	elem := rc.CreateElement(nil)
	r := grt.ShallowRender(elem)
	grt.Equal(t, "<button>Simple Button</button>", r.String())
}

func TestCreateIfNeeded(t *testing.T) {

	// 3 variants:
	button := el.Button(gr.Text("Shiny Button"))
	rc := gr.NewSimpleComponent(button)
	custom := newTestCustomComponent()

	buttonEl := gr.CreateIfNeeded(button)
	rcEl := gr.CreateIfNeeded(rc)
	customEl := gr.CreateIfNeeded(custom)

	grt.Equal(t, button, buttonEl)
	grt.NotNil(t, rcEl.Node())
	grt.Equal(t, custom.Node(), customEl.Node())

	r := grt.ShallowRender(rcEl)

	grt.Equal(t, "<button>Shiny Button</button>", r.String())

}

func TestNew(t *testing.T) {
	component := createLifecycler()

	reactComponent := gr.New(component)
	elem := reactComponent.CreateElement(gr.Props{"text": "Initial Button"})

	r := grt.ShallowRender(elem)

	// We have no desire to test Facebook's React, those people can be trusted
	// to do the right thing.
	// But we use React's behavior to verify that we have set up the components correctly.
	//

	grt.Equal(t, `<div><button style={{"color": "blue"}}>Initial Button</button></div>`,
		r.String())

	this := r.This()

	grt.Equal(t, "blue", this.State().String("color"))

	// Rerender with new props
	newProps := gr.Props{"text": "Updated Button"}
	r.ReRender(newProps)

	grt.Equal(t, `<div><button style={{"color": "blue"}}>Updated Button</button></div>`,
		r.String())

	// Try to rerender with same props. No rerender should happen.
	// Will trigger:
	// 1) ComponentWillReceiveProps
	// 2) ShouldComponentUpdate
	r.ReRender(newProps)

	time.Sleep(200 * time.Millisecond)
	//component.printVisits()
	// TODO(bep) Verify that this is the expected behavior in this case.
	// TODO(bep) Find a way to check the other methods.
	grt.Equal(t, 12, component.totalVisits())

}

func TestCloneElement(t *testing.T) {
	c := gr.New(&testTwoButtons{})

	firstElement := c.CreateElement(gr.Props{"b1": "b1-1", "b2": "b2-1"})

	r := grt.ShallowRender(firstElement)
	grt.Equal(t, "<div><button>b1-1</button><button>b2-1</button></div>", r.String())

	for i := 0; i < 3; i++ {
		b2New := fmt.Sprintf("b2-1%d", i)
		expect := fmt.Sprintf("<div><button>b1-1</button><button>b2-1%d</button></div>", i)

		clonedElement := c.CloneElement(gr.Props{"b2": b2New})

		r = grt.ShallowRender(clonedElement)
		grt.Equal(t, expect, r.String())
	}
}

func TestCompositeComponents(t *testing.T) {
	c := newTestCompositeComponent()
	rc := gr.New(c)

	elem := rc.CreateElement(gr.Props{"text": "Composite Test"})
	r := grt.ShallowRender(elem)

	grt.Equal(t, `<div><h1>Composite Test</h1><tests.testLifecycler text="c1"></tests.testLifecycler><tests.testLifecycler text="c2"></tests.testLifecycler><tests.testLifecycler text="c3"></tests.testLifecycler></div>`,
		r.String())

	// dive into the first subcomponent
	firstSub := r.Dive("tests.testLifecycler")

	grt.Equal(t, `<div><button style={{"color": "red"}}>c1</button></div>`, firstSub.String())

}

func TestNewWithExport(t *testing.T) {
	defer resetComponentState()

	var component gr.Lifecycler = createLifecycler()

	gr.New(component, gr.Export(exportedTestComponent))

	reloaded := js.Module.Get("exports").Get(exportedTestComponent)

	grt.NotNil(t, reloaded)
}

func TestNewWithGlobal(t *testing.T) {
	defer resetComponentState()

	var component gr.Lifecycler = createLifecycler()

	gr.New(component, gr.Global(exportedTestComponent))

	reloaded := js.Global.Get(exportedTestComponent)

	grt.NotNil(t, reloaded)
}

func TestNewWithApply(t *testing.T) {
	applied := js.Global.Get("Object").New()

	applier := func(o *js.Object) *js.Object {
		return applied
	}

	component := createLifecycler()

	reactComponent := gr.New(component, gr.Apply(applier))

	grt.Equal(t, applied, reactComponent.Node())
}

func TestComponentFromGlobal(t *testing.T) {
	rc := gr.FromGlobal("Hello")

	grt.NotNil(t, rc)

	elem := rc.CreateElement(gr.Props{"message": "Go Go React!"})
	r := grt.ShallowRender(elem)

	grt.Equal(t, "<h1>Go Go React!</h1>", r.String())
}

func TestRequire(t *testing.T) {
	rc := gr.Require("react-bootstrap/lib/Alert")

	grt.NotNil(t, rc.Node())

	elem := rc.CreateElement(nil)

	r := grt.ShallowRender(elem)

	grt.Equal(t, `<div role="alert" className="alert alert-info"></div>`, r.String())

}

func TestForceUpdate(t *testing.T) {
	component := createLifecycler()

	reactComponent := gr.New(component)
	elem := reactComponent.CreateElement(gr.Props{"text": "Initial Button"})

	r := grt.ShallowRender(elem)

	grt.Equal(t, `<div><button style={{"color": "blue"}}>Initial Button</button></div>`,
		r.String())

	this := r.This()
	this.SetState(gr.State{"color": "indigo"})

	this.ForceUpdate()

	time.Sleep(200 * time.Millisecond)

	grt.Equal(t, `<div><button style={{"color": "indigo"}}>Initial Button</button></div>`,
		r.String())

	//component.printVisits()
	grt.Equal(t, 10, component.totalVisits())
	grt.Equal(t, 1, component.visitCounter("ComponentWillUpdate"))
	grt.Equal(t, 2, component.visitCounter("Render"))
	grt.Equal(t, 1, component.visitCounter("ComponentDidUpdate"))
	grt.Equal(t, 2, component.visitCounter("GetChildContext"))
	grt.Equal(t, 1, component.visitCounter("GetChildContext-init"))

}

func TestContext(t *testing.T) {

	ctx := gr.Context{"color": "green", "id": 42}

	config := gr.ComponentConfig{ContextTypesTemplate: ctx}

	rc := gr.New(&testParentWithContext{child: gr.New(new(testChildWithContext), gr.WithConfig(config))})

	elem := rc.CreateElement(nil)

	r := grt.ShallowRenderWithContext(elem, ctx)

	c := r.Dive("tests.testChildWithContext")

	grt.Equal(t, `<button id={42} style={{"color": "green"}} />`, c.String())
}

func resetComponentState() {
	js.Global.Set(exportedTestComponent, nil)
	js.Module.Get("exports").Set(exportedTestComponent, nil)
}

func newTestCustomComponent() *testCustomComponent {
	return &testCustomComponent{Object: js.Global.Get("Object").New()}
}

type testCustomComponent struct {
	*js.Object
}

func (c *testCustomComponent) Node() *js.Object {
	return c.Object
}

type testCompositeComponent struct {
	*gr.This
	c1, c2, c3 gr.Factory
}

func newTestCompositeComponent() *testCompositeComponent {
	return &testCompositeComponent{
		c1: createLifecyclerWithColor("red"),
		c2: createLifecyclerWithColor("green"),
		c3: createLifecyclerWithColor("blue"),
	}
}

func (c *testCompositeComponent) Render() gr.Component {
	return el.Div(
		el.Header1(gr.Text(c.This.Props()["text"])),
		c.c1.CreateElement(gr.Props{"text": "c1"}),
		c.c2.CreateElement(gr.Props{"text": "c2"}),
		c.c3.CreateElement(gr.Props{"text": "c3"}),
	)
}

func createLifecycler() *testLifecycler {
	return &testLifecycler{color: "blue", visits: make(map[string]int)}
}

func createLifecyclerWithColor(color string) gr.Factory {
	return gr.New(&testLifecycler{color: color, visits: make(map[string]int)})
}

type testParentWithContext struct {
	child gr.Factory
}

func (l *testParentWithContext) Render() gr.Component {
	return el.Div(
		l.child.CreateElement(nil),
	)
}

func (l *testParentWithContext) GetChildContext() gr.Context {
	return gr.Context{"color": "blue", "id": 62}
}

type testChildWithContext struct {
	*gr.This
}

func (l *testChildWithContext) Render() gr.Component {
	return el.Button(
		gr.Style("color", l.This.Context()["color"]),
		attr.ID(l.This.Context()["id"]),
	)
}

type testLifecycler struct {
	*gr.This
	color  string
	visits map[string]int
}

func (l *testLifecycler) visited(m string) {
	if v, ok := l.visits[m]; ok {
		l.visits[m] = v + 1
	} else {
		l.visits[m] = 1
	}
}

func (l *testLifecycler) totalVisits() int {
	c := 0

	for _, v := range l.visits {
		c += v
	}
	return c
}

func (l *testLifecycler) printVisits() {
	var keys []string
	for k := range l.visits {
		keys = append(keys, k)
	}

	sort.Strings(keys)
	for _, key := range keys {
		println(fmt.Sprintf("%s: %d", key, l.visitCounter(key)))
	}

}

func (l *testLifecycler) visitCounter(m string) int {
	if v, ok := l.visits[m]; ok {
		return v
	}
	return 0
}

func (l *testLifecycler) Render() gr.Component {
	l.visited("Render")
	this := l.This
	color := this.State()["color"]
	text := this.Props()["text"]
	elem := el.Div(
		el.Button(gr.Style("color", color),
			gr.Text(text),
		))
	return elem
}

func (l *testLifecycler) GetInitialState() gr.State {
	l.visited("GetInitialState")
	return gr.State{"color": l.color}
}

func (l *testLifecycler) GetChildContext() gr.Context {
	label := "GetChildContext"
	if l.This.This == nil {
		label += "-init"
	}
	l.visited(label)
	return gr.Context{"color": l.color}
}

func (l *testLifecycler) ShouldComponentUpdate(next gr.Cops) bool {
	l.visited("ShouldComponentUpdate")
	return l.This.Props().HasChanged(next.Props, "text")
}
func (l *testLifecycler) ComponentWillUpdate(next gr.Cops) {
	l.visited("ComponentWillUpdate")
}
func (l *testLifecycler) ComponentWillReceiveProps(data gr.Cops) {
	l.visited("ComponentWillReceiveProps")
}
func (l *testLifecycler) ComponentDidUpdate(prev gr.Cops) {
	l.visited("ComponentDidUpdate")
}
func (l *testLifecycler) ComponentWillMount() {
	l.visited("ComponentWillMount")
}
func (l *testLifecycler) ComponentWillUnmount() {
	l.visited("ComponentWillUnmount")
}
func (l *testLifecycler) ComponentDidMount() {
	l.visited("ComponentDidMount")
	return
}

type testTwoButtons struct {
	*gr.This
}

func (c *testTwoButtons) Render() gr.Component {
	b1Text := c.Props().String("b1")
	b2Text := c.Props().String("b2")
	return el.Div(
		el.Button(gr.Text(b1Text)),
		el.Button(gr.Text(b2Text)),
	)
}
