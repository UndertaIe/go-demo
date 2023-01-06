package exmaple

// 程序打印内容应与Example函数中注释的内容相同时PASS
func ExampleSayHi() {
	SayHi()
	// OutPut: Hi
}

func ExampleSayHiTo() {
	SayHiTo("mikasa")
	// OutPut:
	// Hi
	// mikasa
}

func ExampleSayHiToAll() {
	SayHiToAll("cat", "dog")
	// OutPut:
	// Hi
	// cat
	// dog
}

func ExamplePrintAll() {
	PrintAll() // 无序输出
	// Unordered Output:
	// A
	// B
	// C
}
