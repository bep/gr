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

func Text(text string) Modifier {
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
		panic("Duplicate property" + p.name)
	}
	element.properties[p.name] = p.value
}

func Prop(name string, value interface{}) Modifier {
	return &prop{name: name, value: value}
}

func (m cssClasses) Modify(element *Element) {
	Prop("className", strings.Join(m, " ")).Modify(element)
}

type style struct {
	name  string
	value interface{}
}

func (s *style) Modify(element *Element) {
	if element.style == nil {
		element.style = make(map[string]interface{})
	}
	element.style[s.name] = s.value
}

func Style(name string, value interface{}) Modifier {
	return &style{name: name, value: value}
}
