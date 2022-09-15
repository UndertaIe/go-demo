package core

import "fmt"

type Iface interface {
	f()
}

type structer struct{}

func (s structer) f() {
	fmt.Println("s implemented f()")
}

// 显式表示实现了Iface接口的structer结构体：语言方面存在缺陷，当structer没有全部实现接口对应的方法时会编译失败
func (c Core) ExplicitDeclarationImplOfInterface() {
	var _ Iface = (*structer)(nil)
}
