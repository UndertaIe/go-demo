package core

import "fmt"

func Print[T any](s []T) {
	for _, v := range s {
		fmt.Println(v)
	}
}

func (c Core) DemoOfGeneric() {
	integers := []int{1, 2, 3, 4, 5, 6}
	Print(integers)
	floats := []float32{1, 2, 3, 4, 5, 6}
	Print(floats)
}
