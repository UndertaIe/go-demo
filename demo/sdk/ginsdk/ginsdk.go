package ginsdk

import "fmt"

type Gin struct{}

// 命令描述
func (Gin) Desc() string {
	return "desc"
}

// 子命令
func (Gin) Name() string {
	return "gin"
}

// 文档
func (Gin) Docs() {
	url := ""
	fmt.Println(url)
}
