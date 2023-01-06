package benchmark

const N = 100000

// go test -bench=.
// output:
// goos: linux
// goarch: amd64
// pkg: github.com/UndertaIe/go-demo/test/bench
// cpu: AMD Ryzen 7 5800X 8-Core Processor
// BenchmarkMakeSliceNoCap-4           1028           1011965 ns/op
// BenchmarkMakeSliceWithCap-4         6537            188055 ns/op
// PASS
// ok      github.com/UndertaIe/go-demo/test/bench 3.304s

// go test -bench=MakeSliceNoCap
func MakeSliceNoCap() []int {
	var sli []int
	for i := 0; i < N; i++ {
		sli = append(sli, i)
	}
	return sli
}

// go test -bench=MakeSliceWithCap
func MakeSliceWithCap() []int {
	sli := make([]int, 0, N)
	for i := 0; i < N; i++ {
		sli = append(sli, i)
	}
	return sli
}
