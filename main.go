package main

import (
	"github.com/UndertaIe/go-demo/demo"
	"github.com/UndertaIe/go-demo/demo/core"
	"github.com/UndertaIe/go-demo/demo/sdk/gorm"
	"github.com/UndertaIe/go-demo/demo/sdk/jaeger"
	"github.com/UndertaIe/go-demo/demo/sdk/logger"
	"github.com/UndertaIe/go-demo/demo/sdk/ratelimit"
	"github.com/UndertaIe/go-demo/demo/sdk/sentry"
)

// =========== 加载示例 ===========
// 运行log包的Desc函数
// go run main.go log Desc
// go run main.go core Desc
// go run main.go --help
// go run main.go core --help
func main() {
	defer func() {
		demo.Rdemo.Execute()
	}()

	demo.Fire(demo.Rdemo, new(core.Core))
	demo.Fire(demo.Rdemo, new(logger.L))
	demo.Fire(demo.Rdemo, new(gorm.G))
	demo.Fire(demo.Rdemo, new(sentry.S))
	demo.Fire(demo.Rdemo, new(jaeger.J))
	demo.Fire(demo.Rdemo, new(gorm.G))
	demo.Fire(demo.Rdemo, new(ratelimit.R))

}
