package jaeger

type J struct{}

func (j J) Name() string {
	return "jaeger"
}

func (j J) Desc() string {
	return "jaeger 分布式追踪系统"
}

func (j J) DemoOfJaeger() {

}
