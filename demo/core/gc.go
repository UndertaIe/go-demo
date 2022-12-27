package core

import (
	"runtime"
	"time"
)

func (Core) GC() {
	go func() {
		for {
		}
	}()
	time.Sleep(time.Millisecond)
	runtime.GC()
	println("OK")
} // go1.14 不会打印OK
