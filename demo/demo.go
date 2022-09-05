package demo

import (
	"crypto"
	"fmt"

	"github.com/UndertaIe/passwd/pkg/utils"
)

func Demo() {
	fmt.Println("this is a demo")
}

func Md5Demo() {
	s := utils.Hash(crypto.MD5, "21")
	fmt.Println("md5(21)=", s)
}
