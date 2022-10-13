package parser

import "fmt"

type Parser struct{}

// 命令描述
func (Parser) Desc() string {
	return "解析器"
}

// 子命令
func (Parser) Name() string {
	return "parser"
}

// 文档
func (Parser) Docs() {
	url := ""
	fmt.Println(url)
}
