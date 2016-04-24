package gr

import (
	"github.com/gopherjs/gopherjs/js"
)

var (
	react    = js.Global.Get("React")
	reactDOM = js.Global.Get("ReactDOM")
)

type Component interface {
	Node() *js.Object
}

// The root component.
type Root struct {
	node *js.Object

	// The minimum interface needed to display something.
	renderer Renderer
}

func NewRoot(r Renderer) *Root {
	root := &Root{renderer: r}

	classProps := make(map[string]*js.Object)

	// Every component needs to render itself.
	classProps["render"] = makeNodeFunc(r.Render)

	// Optional lifecycle implementations below.
	if v, ok := r.(ShouldComponentUpdate); ok {
		classProps["shouldComponentUpdate"] = makeComponentUpdateFunc(v.ShouldComponentUpdate)
	}

	if v, ok := r.(ComponentWillUpdate); ok {
		classProps["componentWillUpdate"] = makeComponentUpdateVoidFunc(v.ComponentWillUpdate)
	}

	if v, ok := r.(ComponentDidUpdate); ok {
		classProps["componentDidUpdate"] = makeComponentUpdateVoidFunc(v.ComponentDidUpdate)
	}

	if v, ok := r.(ComponentWillReceiveProps); ok {
		classProps["componentWillReceiveProps"] = makeComponentPropertyReceiverFunc(v.ComponentWillReceiveProps)
	}

	if v, ok := r.(ShouldComponentUpdate); ok {
		classProps["shouldComponentUpdate"] = makeComponentUpdateFunc(v.ShouldComponentUpdate)
	}

	if v, ok := r.(ComponentWillMount); ok {
		classProps["componentWillMount"] = makeVoidFunc(v.ComponentWillMount)
	}

	if v, ok := r.(ComponentDidMount); ok {
		classProps["componentDidMount"] = makeVoidFunc(v.ComponentDidMount)
	}

	if v, ok := r.(ComponentWillUnmount); ok {
		classProps["componentWillUnmount"] = makeVoidFunc(v.ComponentWillUnmount)
	}

	root.node = react.Call("createClass", classProps)

	return root
}

// Implements the Component interface.
func (r *Root) Node() *js.Object {
	return r.node
}

func (r *Root) Render(elementID string, props Props) {
	container := js.Global.Get("document").Call("getElementById", elementID)
	elm := react.Call("createElement", r.Node(), props)
	Render(elm, container)
}

func Render(comp *js.Object, container *js.Object) {
	reactDOM.Call("render", comp, container)
}

type Props map[string]interface{}

// HasChanged reports whether the value of the property with any of the given keys has changed.
func (p Props) HasChanged(nextProps Props, keys ...string) bool {
	for _, key := range keys {
		if p[key] != nextProps[key] {
			return true
		}
	}
	return false
}

type This struct {
	Props Props
}

// Lifecycle interfaces
//
// See http://facebook.github.io/react/docs/component-specs.html#lifecycle-methods

// TODO(bep) Consider some better names. Move to subpackage?

// Renderer is the core interface used to ... render a Component.
type Renderer interface {
	Render(this *This) Component
}

// Invoked before rendering when new props or state are being received.
// This is not called for the initial render or when forceUpdate is used.
type ShouldComponentUpdate interface {
	ShouldComponentUpdate(this *This, nextProps, nextState Props) bool
}

// Invoked immediately before rendering when new props or state are being received.
// This is not called for the initial render.
type ComponentWillUpdate interface {
	ComponentWillUpdate(this *This, nextProps, nextState Props)
}

// Invoked when a component is receiving new props.
// This method is not called for the initial render.
type ComponentWillReceiveProps interface {
	ComponentWillReceiveProps(this *This, props Props)
}

// Invoked immediately after the component's updates are flushed to the DOM.
// This method is not called for the initial render.
type ComponentDidUpdate interface {
	ComponentDidUpdate(this *This, prevProps, prevState Props)
}

// Invoked once, both on the client and server, immediately before the initial rendering occurs.
type ComponentWillMount interface {
	ComponentWillMount(this *This)
}

// Invoked immediately before a component is unmounted from the DOM.
type ComponentWillUnmount interface {
	ComponentWillUnmount(this *This)
}

// Invoked once, only on the client (not on the server),
// immediately after the initial rendering occurs.
type ComponentDidMount interface {
	ComponentDidMount(this *This)
}

func makeComponentUpdateFunc(f func(this *This, props, state Props) bool) *js.Object {
	return js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		return f(extractComponentUpdateArgs(this, arguments))
	})
}

func makeComponentUpdateVoidFunc(f func(this *This, props, state Props)) *js.Object {
	return js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		f(extractComponentUpdateArgs(this, arguments))
		return nil
	})
}

func makeComponentPropertyReceiverFunc(f func(this *This, props Props)) *js.Object {
	return js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		that, props, _ := extractComponentUpdateArgs(this, arguments)
		f(that, props)
		return nil
	})
}

func extractComponentUpdateArgs(this *js.Object, arguments []*js.Object) (*This, Props, Props) {
	var props, state Props

	if arguments[0] != nil {
		props = arguments[0].Interface().(map[string]interface{})
	}
	if arguments[1] != nil {
		state = arguments[1].Interface().(map[string]interface{})
	}

	that := makeThis(this)

	return that, props, state
}

func makeVoidFunc(f func(this *This)) *js.Object {
	return js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		f(makeThis(this))
		return nil
	})
}

func makeNodeFunc(f func(this *This) Component) *js.Object {
	return js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		return f(makeThis(this)).Node()
	})
}

func makeThis(that *js.Object) *This {
	props := that.Get("props").Interface().(map[string]interface{})
	return &This{Props: props}
}
