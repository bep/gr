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
)

type textEl struct {
	text string
}

func (s *textEl) Modify(in *Element) {
	in.text = s.text
}

// Text creates a text element.
func Text(i interface{}) Modifier {
	text := toString(i)
	return &textEl{text: text}
}

type cssClasses []string

// CSS creates a CSS element with the provided classes.
// Note that duplicates are happily accepted.
func CSS(classes ...string) Modifier {
	return cssClasses(classes)
}

// Data creates a data attribute, e.g. data-columns="3"
// See https://developer.mozilla.org/en-US/docs/Web/Guide/HTML/Using_data_attributes
func Data(name, val string) Modifier {
	return Prop("data-"+name, val)
}

// Aria creates an accessibility attributes.
// See https://developer.mozilla.org/en-US/docs/Web/Accessibility/ARIA
func Aria(name, val string) Modifier {
	return Prop("aria-"+name, val)
}

type prop struct {
	name  string
	value interface{}
}

// Modify implements the Modifier interface.
func (p *prop) Modify(element *Element) {
	if element.properties == nil {
		element.properties = make(map[string]interface{})
	}
	if _, ok := element.properties[p.name]; ok {
		panic("Duplicate property: " + p.name)
	}
	element.properties[p.name] = p.value
}

// Prop adds a custom attribute.
func Prop(name string, value interface{}) Modifier {
	return &prop{name: name, value: value}
}

// Modify implements the Modifier interface.
func (m cssClasses) Modify(element *Element) {
	if existing, ok := element.properties["className"]; ok {
		//merge with existing
		element.properties["className"] = existing.(string) + " " + strings.Join(m, " ")
	} else {
		Prop("className", strings.Join(m, " ")).Modify(element)
	}
}

type style struct {
	name  string
	value interface{}
}

type discard int

// Modify implements the Modifier interface with a no-op.
func (d discard) Modify(element *Element) {
	// Do nothing!
}

// Discard is a Modifier that does nothing.
var Discard = new(discard)

// Modify implements the Modifier interface.
func (s *style) Modify(element *Element) {
	if element.style == nil {
		element.style = make(map[string]interface{})
	}
	// last style with a given name wins
	element.style[s.name] = s.value
}

// Style adds a inline CSS style.
func Style(name string, value interface{}) Modifier {
	return &style{name: name, value: value}
}
