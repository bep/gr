package gr

import (
	"errors"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

var (
	react    = js.Global.Get("React")
	reactDOM = js.Global.Get("ReactDOM")
)

// A Component represents a React JS component.
//
// http://facebook.github.io/react/docs/glossary.html#react-nodes for a reference.
//
// A Component can be either a constructed element (analogous to a ReactElement)
// or a factory (a ReactClass or a ReactFactory). Factories are identified by their
// implementation of the Factory interface.
type Component interface {
	Node() *js.Object
}

// A Factory is a Component that can construct Elements (analogous to a ReactClass or a ReactFactory).
type Factory interface {
	Component
	Create(props Props) *Element
}

// ReactComponent wraps a Facebook React component.
// This component can either be constructed from a Go implementation (see New) or
// loaded from JavaScript (see FromJS).
type ReactComponent struct {
	// The React.createClass response.
	node *js.Object

	// The minimum interface needed to display something.
	r Renderer

	// Options
	exportName string
}

// FromJS loads a React component from the JavaScript side of the fence.
// TODO(bep) investigate modules.
func FromJS(path ...string) *ReactComponent {

	var component *js.Object

	for _, p := range path {
		if component != nil {
			component = component.Get(p)
		} else {
			component = js.Global.Get(p)
		}
	}

	if component == nil || component == js.Undefined {
		panic(fmt.Sprintf("JS component in path %v not found", path))
	}

	// TODO(bep): No concept of a Renderer implementation here. Do we need it?
	return &ReactComponent{node: component}
}

// Export is an option used to mark that the component should be exported to the
// JavaScript world as a global with the given name.
// TODO(bep) Again, explore the world of JavaScript modules.
func Export(name string) func(*ReactComponent) error {
	return func(r *ReactComponent) error {
		if name == "" {
			return errors.New("Must provide export name")
		}
		r.exportName = name
		return nil
	}
}

// Apply the func to the newly created React component.
func Apply(f func(o *js.Object) *js.Object) func(*ReactComponent) error {
	return func(r *ReactComponent) error {
		r.node = f(r.node)
		return nil
	}
}

// NewSimpleRenderer can be used for quickly putting together components that only
// need to implement Renderer with no need of the owner (this) argument.
func NewSimpleRenderer(c Component) Renderer {
	return simpleRenderer{c}
}

type simpleRenderer struct {
	c Component
}

// Implements the Renderer interface.
func (s simpleRenderer) Render(this *This) Component {
	return s.c
}

// NewSimpleComponent can be used for quickly putting together components that only
// need to implement Renderer with no need of the owner (this) argument.
// Especially convenient for testing.
func NewSimpleComponent(c Component, options ...func(*ReactComponent) error) *ReactComponent {
	return New(NewSimpleRenderer(c), options...)
}

// New creates a new Component given a Renderer and optinal option(s).
// Note that the Renderer is the minimum interface that needs to be implemented,
// but New will perform interface upgrades for other lifecycle interfaces.
func New(r Renderer, options ...func(*ReactComponent) error) *ReactComponent {
	root := &ReactComponent{r: r}

	classProps := js.Global.Get("Object").New()

	// TODO(bep)
	//classProps.Set("displayName", "TODO")
	//classProps.Set("getDefaultProps",
	//	js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} { return nil }))
	//classProps.Set("propTypes", make(map[string]interface{}))
	//classProps.Set("mixins", nil)
	//classProps.Set("statics", nil)

	// Every component needs to render itself.
	classProps.Set("render", makeRenderFunc(r.Render))

	// Optional lifecycle implementations below.
	if v, ok := r.(StateInitializer); ok {
		classProps.Set("getInitialState", makeStateFunc(v.GetInitialState))
	}

	if v, ok := r.(ShouldComponentUpdate); ok {
		classProps.Set("shouldComponentUpdate", makeComponentUpdateFunc(v.ShouldComponentUpdate))
	}

	if v, ok := r.(ComponentWillUpdate); ok {
		classProps.Set("componentWillUpdate", makeComponentUpdateVoidFunc(v.ComponentWillUpdate))
	}

	if v, ok := r.(ComponentDidUpdate); ok {
		classProps.Set("componentDidUpdate", makeComponentUpdateVoidFunc(v.ComponentDidUpdate))
	}

	if v, ok := r.(ComponentWillReceiveProps); ok {
		classProps.Set("componentWillReceiveProps", makeComponentPropertyReceiverFunc(v.ComponentWillReceiveProps))
	}

	if v, ok := r.(ShouldComponentUpdate); ok {
		classProps.Set("shouldComponentUpdate", makeComponentUpdateFunc(v.ShouldComponentUpdate))
	}

	if v, ok := r.(ComponentWillMount); ok {
		classProps.Set("componentWillMount", makeVoidFunc(v.ComponentWillMount, true))
	}

	if v, ok := r.(ComponentDidMount); ok {
		classProps.Set("componentDidMount", makeVoidFunc(v.ComponentDidMount, true))
	}

	if v, ok := r.(ComponentWillUnmount); ok {
		classProps.Set("componentWillUnmount", makeVoidFunc(v.ComponentWillUnmount, true))
	}

	class := react.Call("createClass", classProps)
	//TODO(bep) factory vs class
	root.node = class // react.Call("createFactory", class)

	for _, opt := range options {
		err := opt(root)
		if err != nil {
			panic(err)
		}
	}

	root.handleOptionsOnCreate()

	return root
}

// CreateIfNeeded evaluates the given Component and returns an Element, creating
// a new instance if needed. This is a convenience method; if you need to pass
// properties, use the factory directly.
func CreateIfNeeded(c Component) *Element {
	switch v := c.(type) {
	case *Element:
		return v
	case Factory:
		return v.Create(nil)
	default:
		return NewPreparedElement(c.Node())
	}
}

// Node implements the Component interface.
func (r *ReactComponent) Node() *js.Object {
	return r.node
}

// Create implements the Factory interface.
func (r *ReactComponent) Create(props Props) *Element {
	elem := react.Call("createElement", r.Node(), props)
	//TODO(bep) factory vs class elem := r.Node().Invoke(props)
	e := NewPreparedElement(elem)
	return e
}

// Render the Component in the DOM with the given element ID and props.
func (r *ReactComponent) Render(elementID string, props Props) {
	container := js.Global.Get("document").Call("getElementById", elementID)
	elem := r.Create(props)

	// TODO(bep) evaluate if the need the "this" returned on render.
	reactDOM.Call("render", elem.Node(), container)
}

// Lifecycle interfaces
//
// See http://facebook.github.io/react/docs/component-specs.html#lifecycle-methods

// TODO(bep) Consider some better names. Move to subpackage?

// Lifecycler contains all the lifecycles. Mostly useful for testing.
type Lifecycler interface {
	Renderer
	StateInitializer
	ShouldComponentUpdate
	ComponentWillUpdate
	ComponentWillReceiveProps
	ComponentDidUpdate
	ComponentWillMount
	ComponentWillUnmount
	ComponentDidMount
}

// Renderer is the core interface used to render a Element.
type Renderer interface {
	Render(this *This) Component
}

// StateInitializer sets up the initial state.
type StateInitializer interface {
	GetInitialState(this *This) State
}

// ShouldComponentUpdate gets invoked before rendering when new props or state are being received.
// This is not called for the initial render or when forceUpdate is used.
type ShouldComponentUpdate interface {
	ShouldComponentUpdate(this *This, nextProps Props, nextState State) bool
}

// ComponentWillUpdate gets invoked immediately before rendering when new props or state are being received.
// This is not called for the initial render.
type ComponentWillUpdate interface {
	ComponentWillUpdate(this *This, nextProps Props, nextState State)
}

// ComponentWillReceiveProps gets invoked when a component is receiving new props.
// This method is not called for the initial render.
type ComponentWillReceiveProps interface {
	ComponentWillReceiveProps(this *This, props Props)
}

// ComponentDidUpdate gets invoked immediately after the component's updates are flushed to the DOM.
// This method is not called for the initial render.
type ComponentDidUpdate interface {
	ComponentDidUpdate(this *This, prevProps Props, prevState State)
}

// ComponentWillMount get invoked once, both on the client and server, immediately before the initial rendering occurs.
type ComponentWillMount interface {
	ComponentWillMount(this *This)
}

// ComponentWillUnmount gets invoked immediately before a component is unmounted from the DOM.
type ComponentWillUnmount interface {
	ComponentWillUnmount(this *This)
}

// ComponentDidMount gets invoked once, only on the client (not on the server),
// immediately after the initial rendering occurs.
type ComponentDidMount interface {
	ComponentDidMount(this *This)
}

func makeComponentUpdateFunc(f func(this *This, props Props, state State) bool) *js.Object {
	return js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		return f(extractComponentUpdateArgs(this, arguments))
	})
}

func makeComponentUpdateVoidFunc(f func(this *This, props Props, state State)) *js.Object {
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

func extractComponentUpdateArgs(this *js.Object, arguments []*js.Object) (*This, Props, State) {
	var (
		props Props
		state State
	)

	if arguments[0] != nil {
		props = arguments[0].Interface().(map[string]interface{})
	}
	if arguments[1] != nil {
		state = arguments[1].Interface().(map[string]interface{})
	}

	that := NewThis(this)

	return that, props, state
}

func makeVoidFunc(f func(this *This), assumeBlocking bool) *js.Object {
	return js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		if assumeBlocking {
			go func() {
				f(NewThis(this))
			}()
		} else {
			f(NewThis(this))
		}
		return nil
	})
}

func makeStateFunc(f func(this *This) State) *js.Object {
	return js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		return f(NewThis(this))
	})
}

func makeRenderFunc(f func(this *This) Component) *js.Object {
	return js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} {
		that := NewThis(this)
		comp := f(that)
		idFactory := &incrementor{}
		addMissingKeys(comp, idFactory)
		// TODO(bep) refactor
		if e, ok := comp.(*Element); ok {
			e.This = that
			addEventListeners(comp, that)
		}
		if _, ok := comp.(Factory); ok {
			panic("Render should return a ready-to-use Element.")
		}
		return comp.Node()
	})
}

func addEventListeners(c Component, that *This) {
	if e, ok := c.(*Element); ok {
		for _, l := range e.eventListeners {
			l.delegate = func(event *js.Object) {
				if l.preventDefault {
					event.Call("preventDefault")
				}
				l.listener(that, &Event{Object: event})
			}

			e.properties[l.name] = l.delegate

		}
		for _, child := range e.children {
			addEventListeners(child, that)
		}

	}
}

func (r *ReactComponent) handleOptionsOnCreate() {
	if r.exportName != "" {
		js.Global.Set(r.exportName, r.node)
	}
}

func addMissingKeys(c Component, id *incrementor) {

	if e, ok := c.(*Element); ok {
		if e.properties == nil {
			e.properties = make(map[string]interface{})
		}
		if _, ok := e.properties["key"]; !ok {
			e.properties["key"] = id.next()
		}
		for _, c2 := range e.children {
			addMissingKeys(c2, id)
		}
	}
}
