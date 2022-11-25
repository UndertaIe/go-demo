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

func (c Core) BadCopy() {
	src := []int{1, 2, 3}
	var dst []int
	copy(dst, src)
	fmt.Println(src)
	fmt.Println(dst)
} // msg:  Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).

// [1 2 3]
// []

func (c Core) RightCopy() {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	fmt.Println("current src: ", src)
	fmt.Println("current dst: ", dst)
	copy(dst, src)
	fmt.Println("after copy, src: ", src)
	fmt.Println("after copy, dst: ", dst)
} // [1 2 3]
// [1 2 3]
