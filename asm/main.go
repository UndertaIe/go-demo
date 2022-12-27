package main

func f(a int, b int) (int, int) {
	return a + b, a - b
}

func main() {
	f(11, 22)
}
