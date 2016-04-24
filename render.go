package gr

import (
	"time"
	//"github.com/gopherjs/gopherjs/js"
)

const (
	defaultFramesPerSecond = 24
	defaultWaitTime        = 1000 / defaultFramesPerSecond
)

func RenderLoop(render func()) chan struct{} {

	quit := make(chan struct{})

	go func() {
		for {
			// TODO(bep): Use this in browsers that supports it.
			//ran := js.Global.Get("requestAnimationFrame")
			select {
			case <-time.After(defaultWaitTime * time.Millisecond):
				render()
			case <-quit:
				return
			}
		}
	}()

	return quit

}
