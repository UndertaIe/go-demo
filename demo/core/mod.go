package core

import (
	"fmt"

	"github.com/ozgio/strutil"
)

// blog: https://segmentfault.com/a/1190000040179648/
// go mod edit -replace [old git package]@[version]=[new git package]@[version]
func (c Core) ReplaceOfMode() {
	s, _ := strutil.Random("1234567890", 4)
	fmt.Println("Random string", s, 20)
	fakeS := strutil.FakeRandom()
	fmt.Println(fakeS)
}
