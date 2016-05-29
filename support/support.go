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

package support

import (
	"errors"
	"fmt"

	"github.com/gopherjs/gopherjs/js"
)

// Require loads a JS object the Node.js way.
// Note that this requires that the require function is present; if in the browser,
// and not in Node.js, try Browserify.
func Require(path string) (*js.Object, error) {
	require := js.Global.Get("require")

	if require == js.Undefined {
		return nil, errors.New("require() not defined; if this is not Node.js, give Browserify a try.")
	}

	m := require.Invoke(path)

	if m == js.Undefined {
		return nil, fmt.Errorf("Module %q not found", path)
	}

	return m, nil
}
