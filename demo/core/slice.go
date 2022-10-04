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

func (c Core) RangeKVofSlice() {
	sli := []int{1, 2, 3}
	for k := range sli {
		fmt.Print(k, " ")
	}
	fmt.Println()
	fmt.Println("==========")
	for k, v := range sli {
		fmt.Println(k, " ", v, " ")
	}
	fmt.Println("==========")
	for _, v := range sli {
		fmt.Println(v, " ")
	}
	fmt.Println("==========")
} // ok
