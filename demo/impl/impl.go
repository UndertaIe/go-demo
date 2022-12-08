package impl

type Impl struct{}

func (Impl) Desc() string {
	desc := "Go语言设计与实现，相关代码"
	return desc
}

func (Impl) Name() string {
	return "impl"
}
