package main

import (
	"github.com/UndertaIe/go-demo/demo"
	"github.com/UndertaIe/go-demo/demo/core"
	"github.com/UndertaIe/go-demo/demo/sdk/logger"
)

// =========== 加载示例 ===========
// go run main.go log Desc
// go run main.go core Desc
// go run main.go --help
// go run main.go core --help
func main() {
	defer func() {
		demo.Rdemo.Execute()
	}()

	demo.Demo(demo.Rdemo, new(core.Core))
	demo.Demo(demo.Rdemo, new(logger.L))

}
