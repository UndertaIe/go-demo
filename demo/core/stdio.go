package core

import (
	"os"
	"strings"
)

func (c Core) StdIoOfCore() {
	f := os.Stdout
	f.ReadFrom(strings.NewReader("go world\n"))
	f.Close()

	f2 := os.Stdin
	f2.ReadFrom(strings.NewReader("go world\n"))
	f2.Close()
}
