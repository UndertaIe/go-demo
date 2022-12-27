package core

import "fmt"

/*
	结构体方法如果想要修改结构体字段， 则接收者是结构体指针

*/

type I interface {
	F()
}

type S1 struct {
	i int
}

func (s S1) F() {
	s.i += 1
}

type S2 struct {
	i int
}

func (s *S2) F() {
	s.i += 1
}

func (Core) ReceiverDemo() {
	// 结构体调用(接收者为结构体)
	s1 := S1{}
	s1.F()

	// 结构体指针调用(接收者为结构体)
	s11 := &S1{}
	s11.F()
	fmt.Println("结构体指针调用(接收者为结构体): s11 =", s11)

	// 结构体转化为I接口调用F方法(接收者为结构体)
	var s111 I = S1{}
	s111.F()

	// 结构体指针转化为接口调用F方法(接收者为结构体)
	var s1111 I = &S1{}
	s1111.F()
	fmt.Println("结构体指针转化为接口调用F方法(接收者为结构体): s1111 =", s11)

	// ========================

	// 结构体调用(接收者为结构体指针)
	s2 := S2{}
	s2.F()
	fmt.Println("结构体调用(接收者为结构体指针): s2 =", s2)

	// 结构体指针调用(接收者为结构体指针)
	s22 := &S2{}
	s22.F()

	// 结构体转化为I接口调用F方法（接收者为结构体指针）
	// var s2 I = S2{} // S2 does not implement I (method F has pointer receiver) 指针接收者方法不能使用 结构体转化为接口 的参数

	// 结构体指针转化I接口调用F方法（接收者为结构体指针）
	var s222 I = &S2{}
	s222.F()
}
