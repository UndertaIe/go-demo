package benchmark

import "testing"

func BenchmarkMakeSliceNoCap(b *testing.B) {
	for i := 0; i < b.N; i++ { // N值动态调整，直到计算出程序单位时间的执行次数
		MakeSliceNoCap()
	}
}

func BenchmarkMakeSliceWithCap(b *testing.B) {
	for i := 0; i < b.N; i++ {
		MakeSliceWithCap()
	}
}
