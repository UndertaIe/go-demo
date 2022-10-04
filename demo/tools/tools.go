package tools

// Go核心基础及底层原理

type Tools struct{}

func (l Tools) Desc() string {
	desc := "Go工具链使用"
	return desc
}

func (l Tools) Name() string {
	return "tools"
}
