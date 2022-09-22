package lib

// Go标准库

type Lib struct{}

func (l Lib) Desc() string {
	desc := "Go标准库"
	return desc
}

func (l Lib) Name() string {
	return "lib"
}
