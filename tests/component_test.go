package tests

import (
	"testing"

	"time"

	"github.com/bep/gr"
	"github.com/bep/gr/el"
	"github.com/bep/gr/tests/grt"
)

// func NewSimpleRenderer(c Component) Renderer {
func TestNewSimpleRenderer(t *testing.T) {
	button := el.Button(gr.Text("Shiny Button"))
	renderer := gr.NewSimpleRenderer(button)
	grt.Equal(t, button, renderer.Render(&gr.This{}))
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
	r.ReRender(gr.Props{"text": "Updated Button"})

	grt.Equal(t, `<div><button style={{"color": "blue"}}>Updated Button</button></div>`,
		r.String())

	time.Sleep(200 * time.Millisecond)

	// TODO(bep) Verify that this is the expected behavior in this case.
	// TODO(bep) Find a way to check the other methods.
	grt.Equal(t, 7, component.visitCounter)

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
	return true
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

// func NewSimpleComponent(c Component, options ...func(*ReactComponent) error) *ReactComponent {
// func New(r Renderer, options ...func(*ReactComponent) error) *ReactComponent {

// type Component interface {
//type Factory interface {

// func FromJS(path ...string) *ReactComponent {

// func Export(name string) func(*ReactComponent) error {
// func Apply(f func(o *js.Object) *js.Object) func(*ReactComponent) error {

// func CreateIfNeeded(c Component) *Element {

// func (r *ReactComponent) Node() *js.Object {

// func (r *ReactComponent) Create(props Props) *Element {
// func (r *ReactComponent) Render(elementID string, props Props) {

// +++ Life cycle
