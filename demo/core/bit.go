package core

import "fmt"

func (Core) BitOperation() {
	i := 1
	x := i << 2
	fmt.Println("i(1) << 2 =", x)
	x >>= 1
	fmt.Println("x(4) >>= 1 =", x)
}
