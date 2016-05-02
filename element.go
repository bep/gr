package gr

import (
	"github.com/gopherjs/gopherjs/js"
)

// A Modifier modifies an element, adding attributes, style or child elements etc.
type Modifier interface {
	Modify(element *Element)
}

type incrementor struct {
	counter int
}

func (i *incrementor) next() int {
	i.counter++
	return i.counter
}

type elementFactory func(e *Element) *js.Object

var defaultElementFactory = (*Element).createElement

var buildOnceFactory = func(o *js.Object) elementFactory {
	return func(e *Element) *js.Object {
		if e.element != nil {
			return e.element
		}
		e.element = o
		return e.element
	}
}

type Element struct {
	tag            string
	properties     map[string]interface{}
	style          map[string]interface{}
	eventListeners []*EventListener

	// An element is either a text node or a container with children.
	text     string
	children []Component

	elFactory elementFactory

	// This is the actual ReactJS element.
	// ReactElement, ReactText or a ReactFragment
	element *js.Object

	*This
}

func NewElement(tag string) *Element {
	return &Element{tag: tag, properties: Props{}, elFactory: defaultElementFactory}
}

// Create an Element from a ready-to-use React element.
func NewPreparedElement(o *js.Object) *Element {
	return &Element{elFactory: buildOnceFactory(o)}
}

func (e *Element) Node() *js.Object {
	e.element = e.elFactory(e)
	return e.element
}

func (e *Element) Modify(in *Element) {
	in.children = append(in.children, e)
}

type Modifiers []Modifier

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

	if e.text != "" {
		return react.Call("createElement", e.tag, e.properties, e.text)
	}

	if len(e.children) > 0 {
		children := make([]interface{}, len(e.children))

		for _, c := range e.children {
			children = append(children, c.Node())
		}

		return createElement(e.tag, e.properties, children)
	}

	return createElement(e.tag, e.properties)
}

func createElement(tag string, props map[string]interface{}, children ...interface{}) *js.Object {
	if len(children) == 0 {
		return react.Call("createElement", tag, props)
	}

	return react.Call("createElement", tag, props, children)
}
