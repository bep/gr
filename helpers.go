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
	"fmt"

	"strconv"

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

// UnmountComponentAtNode unmounts the DOM element at the given ID.
func UnmountComponentAtNode(elementID string) bool {
	// TODO(bep) maybe incorporate this DOM element into the component
	container := js.Global.Get("document").Call("getElementById", elementID)
	return reactDOM.Call("unmountComponentAtNode", container).Bool()
}

// TODO(bep)
func toString(i interface{}) string {
	switch v := i.(type) {
	case string:
		return i.(string)
	case *js.Object:
		if v == js.Undefined {
			return ""
		}
		// TODO(bep) v.Interface() fails with a infinite loop here for strings returned from React.
		// TODO(bep)  Try to make a standalone test that fails.
		o := v.Get("object")
		return o.String()
	}
	return fmt.Sprint(i)
}

func toInt(i interface{}) int {
	switch v := i.(type) {
	case int:
		return v
	case float32:
		return int(v)
	case float64:
		return int(v)
	case string:
		if v == "" {
			return 0
		}
		iv, err := strconv.ParseInt(v, 0, 0)
		if err == nil {
			return int(iv)
		}
		panic(err)
	default:
		panic(fmt.Sprintf("Unhandled number type: %T", v))
	}
}

func objectToMap(o *js.Object) map[string]interface{} {

	m := make(map[string]interface{})

	if o == js.Undefined {
		return m
	}

	for _, k := range js.Keys(o) {
		kv := o.Get(k)

		m[k] = kv

	}

	return m
}



// TODO(bep) get rid of all below.
// HostInfo represents the location info from the browser window.
type HostInfo struct {
	Path     string
	Port     int
	Host     string
	Href     string
	Protocol string
	Origin   string
}

// Location returns info about the current browser location.
func Location() HostInfo {
	l := js.Global.Get("window").Get("location").Interface().(map[string]interface{})
	loc := HostInfo{
		Path:     toString(l["pathname"]),
		Port:     toInt(l["port"]),
		Host:     toString(l["hostname"]),
		Href:     toString(l["href"]),
		Protocol: toString(l["protocol"]),
		Origin:   toString(l["origin"])}

	return loc
}