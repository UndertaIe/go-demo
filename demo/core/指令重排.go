package core

import "runtime"

var a, b int

func fu() {
	a = 1
	b = 2
}

func g() {
	print(b)
	print(a)
}
func (Core) CommandRearrangement() {
	go fu()
	g()
	runtime.GC()
}
