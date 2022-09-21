package jaeger

import "fmt"

type J struct{}

func (j J) Name() string {
	return "jaeger"
}

func (j J) Desc() string {
	return "jaeger 分布式追踪系统"
}

func (j J) Docs() {
	var deployUrl = "https://www.jaegertracing.io/docs/1.37/deployment/"
	fmt.Println(deployUrl)
}
