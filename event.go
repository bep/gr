package gr

import (
	"github.com/gopherjs/gopherjs/js"
)

type Event struct {
	*js.Object
}

func (e *Event) Persist() {
	e.Call("persist")
}

func (e *Event) Int(key string) int {
	return e.Get(key).Int()
}

type EventListener struct {
	name     string
	listener func(*This, *Event)
	// TODO(bep) prevent...
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
