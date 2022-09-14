package core

import "fmt"

type Stru struct {
	A string
	// InterfaceA // 注释此行 error: demo/strcut.go:26:14: impossible type assertion: ifa.(*Stru)和接口方法 --> 编译时错误
}

type InterfaceA interface {
	f()
}

type InterfaceB interface {
	f()
}

// 注释f(): panic: runtime error: invalid memory address or nil pointer dereference保留InterfaceA在struct显式声明 --> 运行时错误
func (s Stru) f() {
	fmt.Println("call f(), A= ", s.A)
}

func callF(ifa InterfaceA) {
	fmt.Println("runtime!!!")
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
// 深入到go底层：对于转换为interface类型其中包含两个值：接口类型以及数据指针
func (c Core) TestStruct2AnyAssert() {
	s := Stru{A: "123"}
	callF(s)
	callF(&s)
	fmt.Println("callF(s) 和 callF(&s)表示struct不管是指针还是自身都可以传入参数类型为实现了interface方法的函数中")
	callByAssert(s)

	callByAssertStruct(s)
	callByAssertStructPointer(&s)
	fmt.Println("类型断言只能获得转化之前的类型 如 struct传入对应行参为any的函数中时，类型断言只能为s.(Stru), 当传入struct指针时，类型断言为s.(*Stru)")

	var i interface{}
	var j interface{}
	i = "mikasa"
	fmt.Println(i == j)
	var z InterfaceA
	fmt.Println(j == z)
	passNil(i)
	passNil(j)

	var l InterfaceA
	var k InterfaceB
	fmt.Println(l == k)
}

func passNil(x interface{}) {
	fmt.Println(x == nil)
}

type Duck interface {
	Quack()
}

type Cat struct{}

func (c *Cat) Quack() {
	fmt.Println("meow")
}

func (c Cat) Eat() {
	fmt.Println("eat something")
}

func (c Core) StructPointerTest() {

	var d1 = Cat{}
	var d2 = &Cat{}
	// struct 和 *struct 都可以调用指针接受者和结构体接受者方法
	// 区别在于指针接受者方法（不管调用者是sturct或者*struct），可以改变struct本身，即struct字段对应的数据
	d1.Quack() // struct call (s *struct)
	d2.Quack() // *struct call (s *struct)
	d1.Eat()   // struct call (s struct)
	d2.Eat()   // *struct call (s struct)

	// var c Duck = Cat{} //:编译期error
	// c.Quack()
	// 当结构体(非结构体指针)转化为接口类型时，调用接受者为指针struct的方法时会报编译时error
	(&Cat{}).Quack()
	// (Cat{}).Quack() // 编译期error: cannot call pointer method Quack on Cat   compilerInvalidMethodExpr

	// callQuack(Cat{}) //: 编译期error
	callQuack(&Cat{})
}

func callQuack(d Duck) {
	d.Quack()
}
