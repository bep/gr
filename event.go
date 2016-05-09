package gr

import (
	"github.com/gopherjs/gopherjs/js"
)

// Event represents a browser event. See https://developer.mozilla.org/en-US/docs/Web/Events
type Event struct {
	*js.Object
}

// Persist can be used to make sure the event survives Facebook Reac's recycling of
// events. Useful to avoid confusing debugging sessions in the console.
func (e *Event) Persist() {
	e.Call("persist")
}

// Int is a convenience method to get an Event attribute as an Int value, e.g. screenX.
func (e *Event) Int(key string) int {
	return e.Get(key).Int()
}

// An EventListener can be attached to a HTML element to listen for events, mouse clicks etc.
type EventListener struct {
	name     string
	listener func(*This, *Event)
	// TODO(bep) prevent...
	preventDefault bool
	delegate       func(jsEvent *js.Object)
}

// Listener is the signature for the func that needs to be implemented by the
// listener, e.g. the clickHandler etc.
type Listener func(*This, *Event)

// NewEventListener creates a new EventListener. In most cases you will use the
// predefined event listeners in the evt package.
func NewEventListener(name string, listener func(*This, *Event)) *EventListener {
	l := &EventListener{name: name, listener: listener}

	return l
}

// Modify implements the Modifier interface.
func (l *EventListener) Modify(element *Element) {
	element.eventListeners = append(element.eventListeners, l)
}
