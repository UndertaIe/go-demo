package lib

import (
	"errors"
	"fmt"
)

func (l Lib) Print() {
	fmt.Print("1")
	fmt.Print("2")
}

func (l Lib) Errorf() {
	var err interface{} = errors.New("123456")
	err = fmt.Errorf("%v", err)
	fmt.Println(err)

}
