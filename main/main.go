package main

import "fmt"

// 逃逸分析 escape analysis

func Escape() {
	Escape0()
	Escape1()
	Escape2()
	Escape3()
	Escape4()
}

type S struct {
	A string
}

// 不存在外部引用不发生逃逸
func Escape0() S {
	s := S{
		"A",
	}
	return s
} // ./main.go:10:2: moved to heap: s

// 存在外部引用发生逃逸
func Escape1() *S {
	s := S{
		"A",
	}
	return &s
} // ./main.go:10:2: moved to heap: s

// 小对象不存在外部引用不会发生逃逸
func Escape2() {
	s := make([]int, 1000)
	for idx := range s {
		s[idx] = idx
	}
} // ./main.go:19:11: make([]int, 1000) does not escape

// 大对象发生逃逸:栈空间不足发生逃逸
func Escape3() {
	s := make([]int, 10000)
	for idx := range s {
		s[idx] = idx
	}
} // ./main.go:26:11: make([]int, 10000) escapes to heap

// 参数为interface{}发生逃逸
func Escape4() {
	str := "mikasa"
	fmt.Println(str)
} // str escapes to heap

func main() {
	Escape()
}
