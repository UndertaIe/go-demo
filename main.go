package main

import (
	"github.com/UndertaIe/go-demo/demo"
	"github.com/UndertaIe/go-demo/demo/core"
	"github.com/UndertaIe/go-demo/demo/lib"
	"github.com/UndertaIe/go-demo/demo/sdk/convert"
	"github.com/UndertaIe/go-demo/demo/sdk/gorm"
	"github.com/UndertaIe/go-demo/demo/sdk/grpc"
	"github.com/UndertaIe/go-demo/demo/sdk/grpc/helloworld"
	"github.com/UndertaIe/go-demo/demo/sdk/jaeger"
	"github.com/UndertaIe/go-demo/demo/sdk/logger"
	"github.com/UndertaIe/go-demo/demo/sdk/parser"
	"github.com/UndertaIe/go-demo/demo/sdk/prometheus"
	"github.com/UndertaIe/go-demo/demo/sdk/protocolbuffer"
	"github.com/UndertaIe/go-demo/demo/sdk/ratelimit"
	"github.com/UndertaIe/go-demo/demo/sdk/sentry"
	"github.com/UndertaIe/go-demo/demo/sdk/swagger"
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
	demo.Fire(demo.Rdemo, new(lib.Lib))
	demo.Fire(demo.Rdemo, new(logger.L))
	demo.Fire(demo.Rdemo, new(gorm.G))
	demo.Fire(demo.Rdemo, new(sentry.S))
	demo.Fire(demo.Rdemo, new(jaeger.J))
	demo.Fire(demo.Rdemo, new(gorm.G))
	demo.Fire(demo.Rdemo, new(ratelimit.R))
	demo.Fire(demo.Rdemo, new(swagger.S))
	demo.Fire(demo.Rdemo, new(prometheus.P))
	demo.Fire(demo.Rdemo, new(convert.C))
	demo.Fire(demo.Rdemo, new(protocolbuffer.PB))
	demo.Fire(demo.Rdemo, new(grpc.G))
	demo.Fire(demo.Rdemo, new(helloworld.H))
	demo.Fire(demo.Rdemo, new(parser.Parser))
}
