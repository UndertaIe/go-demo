package core

import "fmt"

func (Core) ArrayOutOfIndex() {
	var arr [4]int
	fmt.Println("running")
	var x = 5
	fmt.Println(arr[x])
}

func (Core) ArrayOutOfIndex2() {
	fmt.Println("running")
	// var arr [4]int
	// fmt.Println(arr[5]) // demo/core/panic.go:16:18: invalid argument: array index 5 out of bounds [0:4]
}
