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

// Lifecycle interfaces
//
// See http://facebook.github.io/react/docs/component-specs.html#lifecycle-methods

// Lifecycler contains all the lifecycle callback interfaces. Mostly useful for testing.
type Lifecycler interface {
	Renderer
	StateInitializer
	ChildContextProvider
	ShouldComponentUpdate
	ComponentWillUpdate
	ComponentWillReceiveProps
	ComponentDidUpdate
	ComponentWillMount
	ComponentWillUnmount
	ComponentDidMount
}

// Cops holds COntext, Props and State received in the lifecycle methods.
// Note that any of these can be nil, depending on the context.
type Cops struct {
	Context Context
	Props   Props
	State   State
}

// Renderer is the core interface used to render a Element.
type Renderer interface {
	Render() Component
}

// StateInitializer sets up the initial state.
type StateInitializer interface {
	GetInitialState() State
}

// ChildContextProvider provides the context for the children.
//
// The GetChildContext function will be called when the state or props changes.
// In order to update data in the context, trigger a local state update with this.SetState.
// This will trigger a new context and changes will be received by the children.
//
// GetChildContext will also be called once in the init phase, to determine the types for
// the context properties. The this will be nil in this single invocation, and there is no need to return
// real data as long as the types are real (in cases where this is an expensive operation).
//
// See https://facebook.github.io/react/docs/context.html
type ChildContextProvider interface {
	GetChildContext() Context
}

// ShouldComponentUpdate gets invoked before rendering when new props or state are being received.
// This is not called for the initial render or when forceUpdate is used.
type ShouldComponentUpdate interface {
	ShouldComponentUpdate(next Cops) bool
}

// ComponentWillUpdate gets invoked immediately before rendering when new props or state are being received.
// This is not called for the initial render.
type ComponentWillUpdate interface {
	ComponentWillUpdate(next Cops)
}

// ComponentWillReceiveProps gets invoked when a component is receiving new props.
// This method is not called for the initial render.
type ComponentWillReceiveProps interface {
	ComponentWillReceiveProps(next Cops)
}

// ComponentDidUpdate gets invoked immediately after the component's updates are flushed to the DOM.
// This method is not called for the initial render.
type ComponentDidUpdate interface {
	ComponentDidUpdate(prev Cops)
}

// ComponentWillMount get invoked once, both on the client and server, immediately before the initial rendering occurs.
type ComponentWillMount interface {
	ComponentWillMount()
}

// ComponentWillUnmount gets invoked immediately before a component is unmounted from the DOM.
type ComponentWillUnmount interface {
	ComponentWillUnmount()
}

// ComponentDidMount gets invoked once, only on the client (not on the server),
// immediately after the initial rendering occurs.
type ComponentDidMount interface {
	ComponentDidMount()
}
