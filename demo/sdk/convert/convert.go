package convert

import (
	"fmt"

	"github.com/cstockton/go-conv"
)

type C struct{}

func (c C) Name() string {
	return "convert"
}

func (c C) Desc() string {
	return "数据类型转化"
}

func (c C) Docs() {
	var docUrl = ""
	fmt.Println(docUrl)
}

func (c C) Demo() {
	x, err := conv.Int("")
	fmt.Println(x, err)
}
