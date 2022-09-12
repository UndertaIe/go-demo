package main

import (
	"github.com/UndertaIe/go-demo/demo"
	"github.com/UndertaIe/go-demo/demo/sdk/logger"
)

// =========== 加载示例 ===========
func main() {
	defer func() {
		demo.Rdemo.Execute()
	}()

	log := new(logger.L)
	demo.Demo(demo.Rdemo, log)

}
