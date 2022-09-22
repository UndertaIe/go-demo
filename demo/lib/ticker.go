package lib

import (
	"fmt"
	"time"
)

// time trigger
func (l Lib) TickerDemo() {
	tk := time.NewTicker(time.Second * 2)
	start := time.Now()
	for range tk.C {
		fmt.Printf("slaped time: %vs\n", float64(time.Since(start))/float64(time.Second))
		start = time.Now()
	}
}
