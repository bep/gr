package gr

import (
	"github.com/gopherjs/gopherjs/js"
)

type Event struct {
	// TODO(bep) having this embedded is convenient, but may be confusing.
	// An alternetive would be to inject it into the event handler somehow.
	*This
	target *js.Object
}

type EventListener struct {
	name           string
	listener       func(*Event)
	preventDefault bool
	delegate       func(jsEvent *js.Object)
}

type Listener func(*Event)

func NewEventListener(name string, listener func(*Event)) *EventListener {
	l := &EventListener{name: name, listener: listener}

	return l
}

func (l *EventListener) Modify(element *Element) {
	element.eventListeners = append(element.eventListeners, l)
}
