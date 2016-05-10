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
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

type This struct {
	this *js.Object
}

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

func (t *This) State() State {
	state := t.this.Get("state").Interface()
	if state == nil {
		return State{}
	}
	return state.(map[string]interface{})
}

func (s State) Int(key string) int {
	if val, ok := s[key]; ok {
		return int(val.(float64))
	}
	return 0
}

func (t *This) SetState(s State) {
	t.this.Call("setState", s)
}

func NewThis(that *js.Object) *This {
	return &This{this: that}
}

type Context map[string]interface{}
type Props map[string]interface{}
type State map[string]interface{}

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
		if m1[key] != m2[key] {
			return true
		}
	}
	return false
}
