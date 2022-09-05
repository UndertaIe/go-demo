package demo

import "fmt"

type CacheConfig func() map[string]interface{}

func f(ff CacheConfig) {
	fmt.Println(ff())
}

func FuncType() {
	fff := func() map[string]interface{} {
		return map[string]any{"1": "2"}
	}
	if v := CacheConfig(fff); true {
		f(v)
	}
}
