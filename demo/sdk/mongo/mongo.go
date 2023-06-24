package mongo

import (
	"fmt"
)

type Mongo struct{}

// 命令描述
func (Mongo) Desc() string {
	return "do driver for mongodb"
}

// 子命令
func (Mongo) Name() string {
	return "mongo"
}

// 文档
func (Mongo) Docs() {
	url := "https://www.mongodb.com/docs/drivers/go/current/quick-start/"
	fmt.Println(url)
}
