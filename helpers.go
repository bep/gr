package gr

import (
	"github.com/gopherjs/gopherjs/js"
)

//TODO make this work so the examples can be made without an index.html
/*func AddReactJS(version string) {
	for _, lib := range []string{"react-dom.js", "react.js"} {
		link := js.Global.Get("document").Call("createElement", "script")
		link.Set("src", fmt.Sprintf("//cdnjs.cloudflare.com/ajax/libs/react/%s/%s", version, lib))
		//js.Global.Get("document").Get("head").Call("appendChild", link)
		head := js.Global.Get("document").Get("head")
		head.Call("insertBefore", link, head.Get("firstChild"))
	}
}*/

func UnmountComponentAtNode(elementID string) bool {
	// TODO(bep) maybe incorporate this DOM element into the component
	container := js.Global.Get("document").Call("getElementById", elementID)
	return reactDOM.Call("unmountComponentAtNode", container).Bool()
}
