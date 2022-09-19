package core

import "fmt"

type Iface interface {
	f()
}

type structer struct{}

func (s structer) f() {
	fmt.Println("s implemented f()")
}

// Compile-time check that structer implements Iface.
var _ Iface = (*structer)(nil)

func (c Core) ExplicitDeclarationImplOfInterface() {
	var _ Iface = (*structer)(nil)
}
