package core

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

func (c Core) NestedRecover() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("NestedRecover")
		}
	}()
	fmt.Println(1)
	nestedFunc()
} // nestedFunc panic: interface conversion: io.Reader is nil, not *os.File

func nestedFunc() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("nestedFunc panic:", err)
		}
	}()
	fmt.Println(2)
	var reader io.Reader
	fd := reader.(*os.File) // panic
	fmt.Println(fd)
}

func (c Core) NestedRecoverFromRePanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("NestedRecover")
			fmt.Println(reflect.ValueOf(err))
			fmt.Println(err)
		}
	}()
	fmt.Println(1)
	nestedFuncFromRePanic()
} //nestedFunc panic: interface conversion: io.Reader is nil, not *os.File
// NestedRecover

func nestedFuncFromRePanic() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("nestedFunc panic:", err)
			panic(err)
		}
	}()
	fmt.Println(2)
	var reader io.Reader
	fd := reader.(*os.File) // panic
	fmt.Println(fd)
}
