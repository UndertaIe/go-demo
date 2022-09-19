package core

import "fmt"

func (c Core) RangeNilMap() {
	var m map[string]int
	if m == nil {
		fmt.Println("sli is nil")
	}
	for _, item := range m {
		fmt.Println(item)
	}
}

// map引用类型未初始化，其指针变量未指向实际对应的存储数据结构，向其赋值报错
func (c Core) SetNilMapPanic() {
	var m map[string]int
	if m == nil {
		fmt.Println("sli is nil")
	}
	// m["1"] = 2 // assignment to nil map (SA5000)go-staticcheck

} // panic: assignment to entry in nil map

func (c Core) SetMakeMap() {
	var m map[string]int
	if m == nil {
		fmt.Println("sli is nil")
	}
	m = make(map[string]int)
	m["1"] = 2
	fmt.Println(m)
} // panic: assignment to entry in nil map
