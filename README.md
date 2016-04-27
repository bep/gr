# Go React
[GopherJS](https://github.com/gopherjs/gopherjs) bindings for Facebook React. 

**NOTE: Still very early and not production ready.**

## Examples

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

type clickCounter int

// Implements the StateInitializer interface.
func (c clickCounter) GetInitialState(this *gr.This) gr.State {
	return gr.State{"counter": 0}
}

// Implements the Renderer interface.
func (c clickCounter) Render(this *gr.This) gr.Component {
	counter := this.State()["counter"]
	message := fmt.Sprintf("Click me! Number of clicks: %v", counter)

	return el.Div(
		el.Button(
			gr.CSS("btn", "btn-lg", "btn-primary"),
			gr.Text(message),
			evt.Click(c.onClick)))
}

func (c clickCounter) onClick(this *gr.This, event *gr.Event) {
	this.SetState(gr.State{"counter": this.StateInt("counter") + 1})
}

// Implements the ShouldComponentUpdate interface.
func (e clickCounter) ShouldComponentUpdate(
	this *gr.This, nextProps gr.Props, nextState gr.State) bool {

	return this.State().HasChanged(nextState, "counter")
}
```

For help installing GopherJS, please visit [that cool project](https://github.com/gopherjs/gopherjs).

## Inspiration

This project is highly inspired by [Vecty](https://github.com/gopherjs/vecty), a *promising and pure* Go React-like framework. If you're not heavily invested in Facebook's React, take that for a spin.

