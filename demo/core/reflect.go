package core

import (
	"fmt"
	"reflect"
)

// go reflect

func (c Core) DemoOfReflect() {
	var i int
	k := reflect.TypeOf(i).Kind()
	kk := reflect.ValueOf(i).Kind()
	fmt.Println(k)
	fmt.Println(kk)
}
