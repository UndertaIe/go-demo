package demo

import (
	"fmt"
	"reflect"
)

func Desc() {
	fmt.Println("demo desc")
}

type D int

func R() {
	t := reflect.ValueOf(new(D))
	v := reflect.TypeOf(new(D))
	for i := 0; i < t.NumMethod(); i++ {
		mv := t.Method(i)
		fmt.Println(v.Method(i).Name)
		mv.Call(make([]reflect.Value, 0))

	}
}
