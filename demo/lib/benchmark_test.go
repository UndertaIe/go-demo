package lib

import (
	"strings"
	"testing"
)

// go test -bench=.
//		run all bench test
// go test -bench=BenchmarkString
//   	result in running BenchmarkString and BenchmarkStringBuilder

// 每个基准程序执行三次，持续5s
// go test -bench=BenchmarkFoo -benchtime=5s -count=3

// go test -bench=. -benchmem -cpuprofile profile.out
// go test -bench=. -benchmem -memprofile memprofile.out
// go tool pprof profile.out
// go tool pprof memprofile.out

func BenchmarkString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f()
	}
}

func BenchmarkStringBuilder(b *testing.B) {
	for i := 0; i < b.N; i++ {
		f2()
	}
}

func f() {
	var s string
	for i := 0; i < 10000; i++ {
		s += "1"
	}
}

func f2() {
	var s strings.Builder
	for i := 0; i < 10000; i++ {
		s.WriteString("1")
	}
}
