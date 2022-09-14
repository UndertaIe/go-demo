package core

import (
	"fmt"
)

type TestStruct struct{}

func NilOrNot(v interface{}) bool {
	return v == nil
}

type Reader interface {
	string()
}

// go source: nil是一个预先声明的标识符，表示指针、通道、函数、接口、映射或切片类型。
// 总结：nil是指针、映射、切片、通道，函数、接口的零值
func (c Core) NilTest() {
	var s TestStruct
	fmt.Println(NilOrNot(s)) // #=> false
	// fmt.Println(s == nil)    // #=> true

	// var i int
	// fmt.Println(i == nil) // (mismatched types int and untyped nil)

	var r Reader
	fmt.Println(r == nil)
	ToIface(r)

	var m map[string]string
	fmt.Println("m == nil: ", m == nil)
	ToIface(r)

	var slice = []int{1, 2, 3, 4}
	fmt.Println("slice == nil: ", slice == nil)
	ToIface(slice)

	var slice2 []int
	fmt.Println("slice2 == nil: ", slice2 == nil)
	fmt.Print("ToIface: ")
	ToIface(slice2) // 即使切片为零值，但接口中包含了切片[]int的类型信息
	fmt.Print("ToIfaceForSlice: ")
	ToIfaceAssertSlice(slice2)

	var ip *int
	fmt.Println("ip == nil: ", ip == nil)
	i1 := 2
	ip = &i1
	fmt.Println("ip == nil: ", ip == nil)

}

func ToIfaceAssertSlice(i interface{}) {
	x := i.([]int)
	fmt.Println("i.([]int) == nil: ", x == nil)
}

func ToIface(i interface{}) {
	fmt.Println("i == nil: ", i == nil)
}
