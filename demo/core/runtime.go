package core

import (
	"fmt"
	"reflect"
	"runtime"
)

func f() {
	fmt.Println("f()")
}

func (c Core) DemoOfRuntime() {
	fn := runtime.FuncForPC(reflect.ValueOf(f).Pointer()).Name()
	fmt.Println(fn)
}
