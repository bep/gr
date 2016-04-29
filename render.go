package gr

import (
	"time"
	//"github.com/gopherjs/gopherjs/js"
)

const (
	defaultFramesPerSecond = 3 // 24
	defaultWaitTime        = 1000 / defaultFramesPerSecond
)

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
