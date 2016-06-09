# Go React

[![Build Status](https://travis-ci.org/bep/gr.svg)](https://travis-ci.org/bep/gr)
[![GoDoc](https://godoc.org/github.com/bep/gr?status.svg)](https://godoc.org/github.com/bep/gr)
[![Go Report Card](https://goreportcard.com/badge/github.com/bep/gr)](https://goreportcard.com/report/github.com/bep/gr)

**See Also:**

* https://github.com/bep/grcomponents
* [grouter: react-router bindings](https://github.com/bep/grouter)
		
[GopherJS](https://github.com/gopherjs/gopherjs) bindings for Facebook React. 

**NOTE: Still early and not production ready.**

## Examples

**NOTE: Make sure that your GopherJS is up-to-date before running these: `go get -u github.com/gopherjs/gopherjs`**

For a live demo of the examples below, see [http://bego.io/goreact/examples/basic/](http://bego.io/goreact/examples/basic/)  (may not be up-to-date).

There are some runnable examples in `/examples`. Just navigate to that folder and do a:

```bash
gopherjs serve
```
Then navigate to [http://localhost:8080/github.com/bep/gr/examples/](http://localhost:8080/github.com/bep/gr/examples/).

To get a sense of the API, here is the [click-counter](https://github.com/bep/gr/blob/master/examples/basic-click-counter/main.go) example:

```go
func main() {
	component := gr.New(new(clickCounter))

	gr.RenderLoop(func() {
		component.Render("react", gr.Props{})
	})
}

type clickCounter struct {
	*gr.This
}

// Implements the StateInitializer interface.
func (c clickCounter) GetInitialState() gr.State {
	return gr.State{"counter": 0}
}

// Implements the Renderer interface.
func (c clickCounter) Render() gr.Component {
	counter := c.State()["counter"]
	message := fmt.Sprintf(" Click me! Number of clicks: %v", counter)

	elem := el.Div(
		el.Button(
			gr.CSS("btn", "btn-lg", "btn-primary"),
			gr.Style("color", "orange"),
			gr.Text(message),
			evt.Click(c.onClick)))

	return examples.Example("Click Counter", elem)
}

func (c clickCounter) onClick(event *gr.Event) {
	c.SetState(gr.State{"counter": c.State().Int("counter") + 1})
}

// Implements the ShouldComponentUpdate interface.
func (c clickCounter) ShouldComponentUpdate(next gr.Cops) bool {
	return c.State().HasChanged(next.State, "counter")
}
```

For help installing GopherJS, please visit [that cool project](https://github.com/gopherjs/gopherjs).


## Inspiration

This project is highly inspired by [Vecty](https://github.com/gopherjs/vecty), a *promising and pure* Go React-like framework. If you're not heavily invested in Facebook's React, take that for a spin.

