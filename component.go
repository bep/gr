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

package gr

import (
	"errors"
	"fmt"
	"strings"

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
	CreateElement(props Props) *Element
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
	globalName string

	// Needs to be created by createElement as opposed to standalone React factories.
	// TODO(bep) figure a way to extract that info from the JS object.
	needsCreate bool
}

// FromGlobal loads a React component from JavaScript's global object
// ("window" for browsers and "GLOBAL" for Node.js)
func FromGlobal(path ...string) *ReactComponent {

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
	return &ReactComponent{node: component, needsCreate: true}
}

// Require loads a module the Node.js way.
// Note that this requires that the require function is present; if in the browser,
// and not in Node.js, try Browserify.
func Require(path string) *ReactComponent {
	require := js.Global.Get("require")

	if require == js.Undefined {
		panic("require() not defined; if this is not Node.js, give Browserify a try.")
	}

	m := require.Invoke(path)

	if m == js.Undefined {
		panic(fmt.Sprintf("Module %q not found", path))
	}

	return &ReactComponent{node: m, needsCreate: true}
}

// Export is an option used to mark that the component should be exported to the
// JavaScript world as a Node.js module export.
func Export(name string) func(*ReactComponent) error {
	return func(r *ReactComponent) error {
		if name == "" {
			return errors.New("Must provide export name")
		}
		r.exportName = name
		return nil
	}
}

// Global is an option used to mark that the component should be exported to the
// JavaScript world as a global with the given name.
func Global(name string) func(*ReactComponent) error {
	return func(r *ReactComponent) error {
		if name == "" {
			return errors.New("Must provide global name")
		}
		r.globalName = name
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

type reactClass struct {
	*js.Object

	displayName string `js:"displayName"`

	render                    *js.Object `js:"render"`
	getInitialState           *js.Object `js:"getInitialState"`
	shouldComponentUpdate     *js.Object `js:"shouldComponentUpdate"`
	componentWillUpdate       *js.Object `js:"componentWillUpdate"`
	componentDidUpdate        *js.Object `js:"componentDidUpdate"`
	componentWillReceiveProps *js.Object `js:"componentWillReceiveProps"`
	componentWillMount        *js.Object `js:"componentWillMount"`
	componentDidMount         *js.Object `js:"componentDidMount"`
	componentWillUnmount      *js.Object `js:"componentWillUnmount"`
}

// New creates a new Component given a Renderer and optinal option(s).
// Note that the Renderer is the minimum interface that needs to be implemented,
// but New will perform interface upgrades for other lifecycle interfaces.
func New(r Renderer, options ...func(*ReactComponent) error) *ReactComponent {
	root := &ReactComponent{r: r}

	reactClass := &reactClass{Object: js.Global.Get("Object").New()}

	typ := fmt.Sprintf("%T", r)
	displayName := strings.TrimLeft(typ, "*")
	reactClass.displayName = displayName

	//classProps.Set("getDefaultProps", https://github.com/bep/gr/issues/23
	//	js.MakeFunc(func(this *js.Object, arguments []*js.Object) interface{} { return nil }))
	//classProps.Set("propTypes", make(map[string]interface{}))
	//classProps.Set("mixins", nil) https://github.com/bep/gr/issues/24
	//classProps.Set("statics", nil) https://github.com/bep/gr/issues/25

	// Every component needs to render itself.
	reactClass.render = makeRenderFunc(r.Render)

	// Optional lifecycle implementations below.
	if v, ok := r.(StateInitializer); ok {
		reactClass.getInitialState = makeStateFunc(v.GetInitialState)
	}

	if v, ok := r.(ShouldComponentUpdate); ok {
		reactClass.shouldComponentUpdate = makeComponentUpdateFunc(v.ShouldComponentUpdate)
	}

	if v, ok := r.(ComponentWillUpdate); ok {
		reactClass.componentWillUpdate = makeComponentUpdateVoidFunc(v.ComponentWillUpdate)
	}

	if v, ok := r.(ComponentDidUpdate); ok {
		reactClass.componentDidUpdate = makeComponentUpdateVoidFunc(v.ComponentDidUpdate)
	}

	if v, ok := r.(ComponentWillReceiveProps); ok {
		reactClass.componentWillReceiveProps = makeComponentPropertyReceiverFunc(v.ComponentWillReceiveProps)
	}

	if v, ok := r.(ComponentWillMount); ok {
		reactClass.componentWillMount = makeVoidFunc(v.ComponentWillMount, true)
	}

	if v, ok := r.(ComponentDidMount); ok {
		reactClass.componentDidMount = makeVoidFunc(v.ComponentDidMount, true)
	}

	if v, ok := r.(ComponentWillUnmount); ok {
		reactClass.componentWillUnmount = makeVoidFunc(v.ComponentWillUnmount, true)
	}

	class := react.Call("createClass", reactClass.Object)

	root.node = react.Call("createFactory", class)

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
		return v.CreateElement(nil)
	default:
		return NewPreparedElement(c.Node())
	}
}

// Node implements the Component interface.
func (r *ReactComponent) Node() *js.Object {
	return r.node
}

// CreateElement implements the Factory interface.
func (r *ReactComponent) CreateElement(props Props) *Element {
	var elem *js.Object

	if r.needsCreate {
		elem = react.Call("createElement", r.Node(), props)
	} else {
		elem = r.Node().Invoke(props)
	}

	e := NewPreparedElement(elem)
	return e
}

// Render the Component in the DOM with the given element ID and props.
func (r *ReactComponent) Render(elementID string, props Props) {
	container := js.Global.Get("document").Call("getElementById", elementID)
	elem := r.CreateElement(props)

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
		exports := js.Module.Get("exports")
		if exports == js.Undefined {
			panic("module.exports not present.")
		}
		exports.Set(r.exportName, r.node)
	}
	if r.globalName != "" {
		js.Global.Set(r.globalName, r.node)
	}
}
