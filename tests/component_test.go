package tests

import (
	"testing"

	"time"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/tests/grt"
	"github.com/gopherjs/gopherjs/js"
)

const exportedTestComponent = "GrtTest"

func TestNewSimpleRenderer(t *testing.T) {
	button := el.Button(gr.Text("Shiny Button"))
	renderer := gr.NewSimpleRenderer(button)
	grt.Equal(t, button, renderer.Render(&gr.This{}))
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

	grt.Equal(t, "blue", this.State()["color"])

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

	// TODO(bep) Verify that this is the expected behavior in this case.
	// TODO(bep) Find a way to check the other methods.
	grt.Equal(t, 9, component.visitCounter)

}

func TestNewWithExport(t *testing.T) {
	defer resetComponentState()

	var component gr.Lifecycler = createLifecycler()

	gr.New(component, gr.Export(exportedTestComponent))

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

func TestComponentFromJS(t *testing.T) {
	rc := gr.FromJS("Hello")

	grt.NotNil(t, rc)

	elem := rc.CreateElement(gr.Props{"message": "Go Go React!"})
	r := grt.ShallowRender(elem)

	grt.Equal(t, "<h1>Go Go React!</h1>", r.String())
}

func resetComponentState() {
	js.Global.Set(exportedTestComponent, nil)
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

func createLifecycler() *testLifecycler {
	return &testLifecycler{e: el.Button(gr.Text("Shiny Button"))}
}

type testLifecycler struct {
	e            *gr.Element
	visitCounter int
}

func (l *testLifecycler) visited() {
	l.visitCounter++
}

func (l *testLifecycler) Render(this *gr.This) gr.Component {
	l.visited()
	println("Render")
	color := this.State()["color"]
	text := this.Props()["text"]
	elem := el.Div(
		el.Button(gr.Style("color", color),
			gr.Text(text),
		))
	return elem
}

func (l *testLifecycler) GetInitialState(this *gr.This) gr.State {
	l.visited()
	println("GetInitialState")
	return gr.State{"color": "blue"}
}
func (l *testLifecycler) ShouldComponentUpdate(this *gr.This, nextProps gr.Props, nextState gr.State) bool {
	l.visited()
	println("ShouldComponentUpdate")
	return this.Props().HasChanged(nextProps, "text")
}
func (l *testLifecycler) ComponentWillUpdate(this *gr.This, nextProps gr.Props, nextState gr.State) {
	l.visited()
	println("ComponentWillUpdate")
}
func (l *testLifecycler) ComponentWillReceiveProps(this *gr.This, props gr.Props) {
	l.visited()
	println("ComponentWillReceiveProps")
}
func (l *testLifecycler) ComponentDidUpdate(this *gr.This, prevProps gr.Props, prevState gr.State) {
	l.visited()
	println("ComponentDidUpdate")
}
func (l *testLifecycler) ComponentWillMount(this *gr.This) {
	l.visited()
	println("ComponentWillMount")
}
func (l *testLifecycler) ComponentWillUnmount(this *gr.This) {
	l.visited()
	println("ComponentWillUnmount")
}
func (l *testLifecycler) ComponentDidMount(this *gr.This) {
	l.visited()
	println("ComponentDidMount")
	return
}
