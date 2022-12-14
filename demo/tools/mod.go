package tools

import (
	"fmt"

	"github.com/ozgio/strutil"
)

// blog: https://segmentfault.com/a/1190000040179648/
// go mod edit -replace [old git package]@[version]=[new git package]@[version]
// package未修改后项目名称，若是项目的package则不生效
// replace github.com/getsentry/sentry-go v0.13.0 => github.com/UndertaIe/sentry-go v1.0.0   		 √√√√
// replace github.com/getsentry/sentry-go/gin v0.13.0 => github.com/UndertaIe/sentry-go/gin v1.0.0   ××××
// ！！！ go.mod没有依赖处理，当A项目导入B Module(mod使用了replace)，则A跟B一样要使用相同的replace
func (c Tools) ReplaceOfMode() {
	s, _ := strutil.Random("1234567890", 4)
	fmt.Println("Random string", s, 20)
	fakeS := strutil.FakeRandom()
	fmt.Println(fakeS)
}
