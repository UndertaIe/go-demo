package gomicro

import "fmt"

type Gomicro struct{}

// 命令描述
func (Gomicro) Desc() string {
	return "Go微服务框架"
}

// 子命令
func (Gomicro) Name() string {
	return "gomicro"
}

// 文档
func (Gomicro) Docs() {
	url := "https://github.com/go-micro/go-micro"
	fmt.Println(url)
}
