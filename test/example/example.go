package exmaple

import "fmt"

func SayHi() {
	fmt.Println("Hi")
}

func SayHiTo(name string) {
	fmt.Println("Hi")
	fmt.Println(name)
}

func SayHiToAll(names ...string) {
	fmt.Println("Hi")
	for _, n := range names {
		fmt.Println(n)
	}
}

func PrintAll(names ...string) {
	elements := make(map[int]string)
	elements[0] = "A"
	elements[1] = "B"
	elements[2] = "C"
	for _, val := range elements {
		fmt.Println(val)
	}
}
