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
	"time"
	//"github.com/gopherjs/gopherjs/js"
)

const (
	defaultFramesPerSecond = 3 // 24
	defaultWaitTime        = 1000 / defaultFramesPerSecond
)

// RenderLoop runs the given render func in a loop at the given interval.
// It can be stopped by closing the returned channel.
func RenderLoop(render func(), interval ...time.Duration) chan struct{} {

	renderInterval := defaultWaitTime * time.Millisecond

	if len(interval) > 0 {
		renderInterval = interval[0]
	}

	quit := make(chan struct{})

	go func() {
		for {
			// TODO(bep): Use this in browsers that supports it.
			//ran := js.Global.Get("requestAnimationFrame")
			select {
			case <-time.After(renderInterval):
				render()
			case <-quit:
				return
			}
		}
	}()

	return quit
}
