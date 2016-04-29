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

func Text(i interface{}) Modifier {
	text := toString(i)
	return &textEl{text: text}
}

type cssClasses []string

func CSS(classes ...string) Modifier {
	return cssClasses(classes)
}

func Data(name, val string) Modifier {
	return Prop("data-"+name, val)
}

func Aria(name, val string) Modifier {
	return Prop("aria-"+name, val)
}

type prop struct {
	name  string
	value interface{}
}

func (p *prop) Modify(element *Element) {
	if element.properties == nil {
		element.properties = make(map[string]interface{})
	}
	if _, ok := element.properties[p.name]; ok {
		panic("Duplicate property: " + p.name)
	}
	element.properties[p.name] = p.value
}

func Prop(name string, value interface{}) Modifier {
	return &prop{name: name, value: value}
}

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

func (d discard) Modify(element *Element) {
	// Do nothing!
}

// A Modifier that does nothing.
var Discard = new(discard)

func (s *style) Modify(element *Element) {
	if element.style == nil {
		element.style = make(map[string]interface{})
	}
	element.style[s.name] = s.value
}

func Style(name string, value interface{}) Modifier {
	return &style{name: name, value: value}
}
