package gr

import (
	"github.com/gopherjs/gopherjs/js"
)

type Event struct {
	target *js.Object
}

type EventListener struct {
	name           string
	listener       func(*This, *Event)
	preventDefault bool
	delegate       func(jsEvent *js.Object)
}

type Listener func(*This, *Event)

func NewEventListener(name string, listener func(*This, *Event)) *EventListener {
	l := &EventListener{name: name, listener: listener}

	return l
}

func (l *EventListener) Modify(element *Element) {
	element.eventListeners = append(element.eventListeners, l)
}
