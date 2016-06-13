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
	"github.com/gopherjs/gopherjs/js"
)

// A Modifier modifies an element, adding attributes, style or child elements etc.
type Modifier interface {
	Modify(element *Element)
}

type elementFactory func(e *Element) *js.Object

var (
	defaultElementFactory = (*Element).createElement
	returnStoredElement   = func(e *Element) *js.Object {
		return e.element
	}
)

// Element represents a builder for a ReactElement.
// An Element can be a simple text node or a HTML element with children, attributes etc.
type Element struct {
	tag string

	properties     map[string]interface{}
	style          map[string]interface{}
	eventListeners []*EventListener

	children []Component

	elFactory elementFactory

	// We assume that this element is static, and as such we help by
	// adding auto generated keys when missing, to work around
	// annyoing warnings in the console.
	// This can be switched with the Dynamic modifier.
	dynamic bool

	// This is the actual ReactJS element.
	// ReactElement, ReactText or a ReactFragment
	element *js.Object
}

// NewElement creates a new Element with the given tag.
func NewElement(tag string) *Element {
	return &Element{tag: tag, properties: Props{}, elFactory: defaultElementFactory}
}

// NewPreparedElement creates an Element from a ready-to-use React element.
func NewPreparedElement(o *js.Object) *Element {
	return &Element{element: o, elFactory: returnStoredElement}
}

// Node returns the resulting ReactElement.
func (e *Element) Node() *js.Object {
	if e.element == nil {
		e.element = e.elFactory(e)
	}
	return e.element
}

// Modify implements the Modifier interface.
func (e *Element) Modify(in *Element) {
	in.children = append(in.children, e)
}

// Modifiers is used to Modify a list of elements (children).
type Modifiers []Modifier

// Modify implements the Modifier interface.
func (mods Modifiers) Modify(e *Element) {
	for _, m := range mods {
		if m != nil {
			m.Modify(e)
		}
	}
}

func (e *Element) createElement() *js.Object {
	if e.properties == nil {
		e.properties = make(map[string]interface{})
	}

	if len(e.style) != 0 {
		// TODO(bep) merge with existing?
		e.properties["style"] = e.style
	}

	var args []interface{}

	if len(e.children) > 0 {
		for _, c := range e.children {
			args = append(args, c.Node())
		}
	}

	return createElement(e.tag, e.properties, args)

}

func createElement(tag string, props map[string]interface{}, args []interface{}) *js.Object {
	if len(args) == 0 {
		return react.Call("createElement", tag, props)
	}

	return react.Call("createElement", tag, props, args)
}
