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
	"fmt"
	"reflect"
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

// This is named for what it represents: The this context representation from the
// JavaScript side of the fence.
type This struct {
	this *js.Object
}

// Props returns the properties set; what you would expect to find in
// this.properties in React.
func (t *This) Props() Props {

	// TODO(bep) cache?
	po := t.this.Get("props")
	keys := js.Keys(po)
	props := Props{}
	for _, key := range keys {
		o := po.Get(key)
		if o == js.Undefined {
			panic("Got undefined object in props: " + key)
		}
		if o == nil {
			continue
		}
		if strings.Contains(o.Get("$$typeof").String(), "react.element") {
			// React elements can contain circular refs that goes into
			// Uncaught RangeError: Maximum call stack size exceeded
			// So, wrap them.
			props[key] = js.MakeWrapper(po.Get(key))
		} else {
			props[key] = o.Interface()
		}
	}

	return props
}

// Context returns the context set; what you would expect to find in
// this.context in React.
func (t *This) Context() Context {
	c := t.this.Get("context")

	if c == js.Undefined {
		panic("No context found")
	}

	return c.Interface().(map[string]interface{})
}

// Component returns a component stored in props by its name.
func (t *This) Component(name string) Modifier {
	props := t.Props()
	if main, ok := props[name]; ok {
		comp := main.(*js.Object).Interface().(*js.Object)
		return NewPreparedElement(comp)
	}
	return Discard
}

// State returns the state; what you would expect to find in
// this.properties in React.
func (t *This) State() State {
	state := t.this.Get("state").Interface()
	if state == nil {
		return State{}
	}
	return state.(map[string]interface{})
}

// Int is concenience method to lookup a int value from state.
func (s State) Int(key string) int {
	if val, ok := s[key]; ok {
		return int(val.(float64))
	}
	return 0
}

// Bool is concenience method to lookup a bool value from state.
func (s State) Bool(key string) bool {
	if val, ok := s[key]; ok {
		return val.(bool)
	}
	panic(fmt.Sprintf("State variable %q not found", key))
}

// SetState is a way of setting the state.
func (t *This) SetState(s State) {
	t.this.Call("setState", s)
}

// Refs returns the component references.
// See https://facebook.github.io/react/docs/more-about-refs.html
func (t *This) Refs() Refs {
	refs := t.this.Get("refs").Interface()
	return refs.(map[string]interface{})
}

// GetDOMNode returns the component from refs if it has been mounted into the DOM.
func (r Refs) GetDOMNode(key string) *js.Object {
	if o, ok := r[key]; ok {
		return reactDOM.Call("findDOMNode", o)
	}
	return nil
}

// ForceUpdate forces a re-render of the component.
func (t *This) ForceUpdate() {
	t.this.Call("forceUpdate")
}

// NewThis creates a new This based on a JavaScript object representation.
func NewThis(that *js.Object) *This {
	return &This{this: that}
}

// Context holds the React context.
type Context map[string]interface{}

// Props holds the React properties.
type Props map[string]interface{}

// State holds the React state.
type State map[string]interface{}

// Refs holds a reference to component references.
// See https://facebook.github.io/react/docs/more-about-refs.html
type Refs map[string]interface{}

// Call calls a func with the given name in Props with the given args.
func (p Props) Call(name string, args ...interface{}) *js.Object {
	f := p.Func(name)
	return f(args...)
}

// Func returns the func with the given name in Props.
func (p Props) Func(name string) func(args ...interface{}) *js.Object {
	if o, ok := p[name]; ok {
		return o.(func(args ...interface{}) *js.Object)
	}
	panic(fmt.Sprintf("func %s not found in properties", name))
}

// Children represents a component's child component(s).
// This is a fairly complex topic in Facebook's React, and more support may arrive here, eventually.
// See https://facebook.github.io/react/tips/children-props-type.html
type Children struct {
	*js.Object
}

// Children returns this component's children, if any.
func (t *This) Children() *Children {
	o := t.this.Get("props").Get("children")

	if o == js.Undefined {
		return nil
	}

	return &Children{o}
}

// Element returns the children as an Element ready to render.
func (c Children) Element() *Element {
	return NewPreparedElement(c.Object)
}

// HasChanged reports whether the value of the property with any of the given keys has changed.
func (p Props) HasChanged(nextProps Props, keys ...string) bool {
	return hasChanged(p, nextProps, keys...)
}

// HasChanged reports whether the value of the state with any of the given keys has changed.
func (s State) HasChanged(nextState State, keys ...string) bool {
	return hasChanged(s, nextState, keys...)
}

func hasChanged(m1, m2 map[string]interface{}, keys ...string) bool {
	for _, key := range keys {
		if !reflect.DeepEqual(m1[key], m2[key]) {
			return true
		}
	}
	return false
}
