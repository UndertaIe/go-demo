package core

import "fmt"

func (c Core) RangeNilSlice() {
	var sli []int
	if sli == nil {
		fmt.Println("sli is nil")
	}
	for _, item := range sli {
		fmt.Println(item)
	}
} // ok

func (c Core) AppendToNilSlice() {
	var sli []int
	if sli == nil {
		fmt.Println("sli is nil")
	}
	sli = append(sli, 1)
	fmt.Println(sli)
} // ok
