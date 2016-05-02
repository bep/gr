package gr

import (
	"strings"

	"github.com/gopherjs/gopherjs/js"
)

type This struct {
	this *js.Object
}

func (t *This) Props() Props {
	// TODO(bep) cache?
	po := t.this.Get("props")
	keys := js.Keys(po)
	props := Props{}
	for _, key := range keys {
		o := po.Get(key)
		if strings.Contains(o.Get("$$typeof").String(), "react.element") {
			// React elements can contain circular refs that goes into
			// Uncaught RangeError: Maximum call stack size exceeded
			// So, wrap them.
			props[key] = js.MakeWrapper(po.Get(key))
		} else {
			props[key] = o.Interface()
		}
	}

	return props
}

func (t *This) Context() *js.Object {
	return t.this.Get("context")
}

// Component returns a component stored in props by its name.
func (t *This) Component(name string) Modifier {
	props := t.Props()
	if main, ok := props[name]; ok {
		comp := main.(*js.Object).Interface().(*js.Object)
		return NewPreparedElement(comp)
	}
	return Discard
}

func (t *This) _Props() Props {
	props := t.this.Get("props").Interface()
	if props == nil {
		return Props{}
	}
	return props.(map[string]interface{})
}

func (t *This) State() State {
	state := t.this.Get("state").Interface()
	if state == nil {
		return State{}
	}
	return state.(map[string]interface{})
}

func (t *This) StateInt(key string) int {
	if val, ok := t.State()[key]; ok {
		return int(val.(float64))
	}
	return 0
}

func (t *This) SetState(s State) {
	t.this.Call("setState", s)
}

func makeThis(that *js.Object) *This {
	return &This{this: that}
}

type Props map[string]interface{}
type State map[string]interface{}

// HasChanged reports whether the value of the property with any of the given keys has changed.
func (p Props) HasChanged(nextProps Props, keys ...string) bool {
	return hasChanged(p, nextProps, keys...)
}

// HasChanged reports whether the value of the state with any of the given keys has changed.
func (s State) HasChanged(nextState State, keys ...string) bool {
	return hasChanged(s, nextState, keys...)
}

func hasChanged(m1, m2 map[string]interface{}, keys ...string) bool {
	for _, key := range keys {
		if m1[key] != m2[key] {
			return true
		}
	}
	return false
}
