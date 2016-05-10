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
	name           string
	listener       func(*This, *Event)
	preventDefault bool
	delegate       func(jsEvent *js.Object)
}

// PreventDefault prevents the default event behaviour in the browser.
func (l *EventListener) PreventDefault() *EventListener {
	l.preventDefault = true
	return l
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
