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

	"github.com/gopherjs/gopherjs/js"
)

// This is named for what it represents: The this context representation from the
// JavaScript side of the fence.
type This struct {
	This *js.Object
}

// InitThis implements the ThisInitializer.
func (t *This) InitThis(that *js.Object) {
	t.This = that
}

// ThisInitializer must be implemented by a component if JavaScript's this is needed.
// The simplest way of doing this is to just embed a This on the component struct.
type ThisInitializer interface {
	InitThis(this *js.Object)
}

// Props returns the properties set; this is what you would expect to find in
// this.props in React.
func (t *This) Props() Props {
	return objectToMap(t.This.Get("props"))
}

// Context returns the context set; what you would expect to find in
// this.context in React.
func (t *This) Context() Context {
	return objectToMap(t.This.Get("context"))
}

// Component returns a component stored in props by its name.
func (t *This) Component(name string) Modifier {
	props := t.Props()
	if main, ok := props[name]; ok {
		return NewPreparedElement(main.(*js.Object))
	}
	return Discard
}

// IsMounted reports whether this component is mounted.
func (t *This) IsMounted() bool {
	return t.This.Call("isMounted").Bool()
}

// State returns the state; what you would expect to find in
// this.properties in React.
func (t *This) State() State {
	return objectToMap(t.This.Get("state"))
}

// Int is convenience method to lookup a int value from state.
func (s State) Int(key string) int {
	if val, ok := s[key]; ok {
		return val.(*js.Object).Int()
	}
	return 0
}

// Bool is convenience method to lookup a bool value from state.
func (s State) Bool(key string) bool {
	if val, ok := s[key]; ok {
		return val.(*js.Object).Bool()
	}
	panic(fmt.Sprintf("State variable %q not found", key))
}

// String is convenience method to lookup a bool value from state.
func (s State) String(key string) string {
	if val, ok := s[key]; ok {
		return val.(*js.Object).String()
	}
	return ""
}

// Interface is a convenience method to lookup an interface value from state.
func (s State) Interface(key string) interface{} {
	if val, ok := s[key]; ok {
		return val.(*js.Object).Interface()
	}
	return nil
}

// SetState sets the state with a map of Go interface{} values.
func (t *This) SetState(s State) {
	t.This.Call("setState", s)
}

// Refs returns the component references.
// See https://facebook.github.io/react/docs/more-about-refs.html
func (t *This) Refs() Refs {
	return objectToMap(t.This.Get("refs"))
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
	t.This.Call("forceUpdate")
}

// NewThis creates a new This based on a JavaScript object representation.
func NewThis(that *js.Object) *This {
	return &This{This: that}
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
		return o.(*js.Object).Interface().(func(args ...interface{}) *js.Object)
	}
	panic(fmt.Sprintf("func %s not found in properties", name))
}

// Interface is a convenience method to lookup an interface value from props.
func (p Props) Interface(key string) interface{} {
	if val, ok := p[key]; ok {
		return val.(*js.Object).Interface()
	}
	return nil
}

// Children represents a component's child component(s).
// This is a fairly complex topic in Facebook's React, and more support may arrive here, eventually.
// See https://facebook.github.io/react/tips/children-props-type.html
type Children struct {
	*js.Object
}

// Children returns this component's children, if any.
func (t *This) Children() *Children {
	o := t.This.Get("props").Get("children")

	if o == js.Undefined {
		return nil
	}

	return &Children{o}
}

// Element returns the children as an Element ready to render.
func (c *Children) Element() *Element {
	return NewPreparedElement(c.Object)
}

// HasChanged reports whether the value of the property with any of the given keys has changed.
func (p Props) HasChanged(nextProps Props, keys ...string) bool {
	return hasChanged(p, nextProps, keys...)
}

// HasChanged reports whether the value of the state with any of the given keys has changed.
// Note: This does only reference checking, so no deep equality.
// See https://github.com/gopherjs/gopherjs/issues/473
func (s State) HasChanged(nextState State, keys ...string) bool {
	return hasChanged(s, nextState, keys...)
}

func hasChanged(m1, m2 map[string]interface{}, keys ...string) bool {
	for _, key := range keys {
		// TODO(bep) DeepEqual
		if m1[key] != m2[key] {
			return true
		}
	}
	return false
}
