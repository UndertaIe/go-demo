package demo

import "fmt"

type Stru struct {
	A string
	InterfaceA
}

type InterfaceA interface {
	f()
}

func (s Stru) f() {
	fmt.Println("call f(), A= ", s.A)
}

func callF(ifa InterfaceA) {
	ifa.f()
}
func callByAssert(ifa InterfaceA) {

	if s, ok := ifa.(*Stru); ok { // 传入类型为*Stru时断言成功
		fmt.Print("callByAssert(&Stru):) ")
		s.f()
	}
	if s2, ok := ifa.(Stru); ok { // 传入类型为Stru时断言失败
		fmt.Print("callByAssert(Stru):) ")
		s2.f()
	}
} // !!! error

func callByAssertStruct(ifa interface{}) {
	fmt.Print("callByAssertStruct:) ")
	s2 := ifa.(Stru)
	s2.f()
} // right

func callByAssertStructPointer(ifa interface{}) {
	fmt.Print("callByAssertStructPointer:) ")
	s2 := ifa.(*Stru)
	s2.f()
} // right

// 测试struct 和*struct类型断言的区别
// 结论是    类型断言只能断言转换接口之前的类型 传入类型为指针则断言指针，传入为struct则断言struct，
//			对于调用对应的接口方法则无区别，struct指针和struct自身都可以调用，但struct指针可以改变接收者本身，而struct则不能，适用于接收者传入和参数传入
func TestStruct2AnyAssert() {
	s := Stru{A: "123"}
	callF(s)
	callF(&s)
	fmt.Println("callF(s) 和 callF(&s)表示struct不管是指针还是自身都可以传入参数类型为实现了interface方法的函数中")
	callByAssert(s)

	callByAssertStruct(s)
	callByAssertStructPointer(&s)
	fmt.Println("类型断言只能获得转化之前的类型 如 struct传入对应行参为any的函数中时，类型断言只能为s.(Stru), 当传入struct指针时，类型断言为s.(*Stru)")

}
