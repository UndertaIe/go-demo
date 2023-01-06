package lib

import (
	"encoding/json"
	"fmt"
)

type T struct {
	A string `json:"a,omitempty"`
	B int    `json:"b,omitempty"`
}

func (Lib) Omitempty() {
	var t T
	s, _ := json.Marshal(t)
	fmt.Println(string(s))
	t.B = 2
	s, _ = json.Marshal(t)
	fmt.Println(string(s))
}

type A struct {
	A string `json:"a"`
	B string `json:"b"`
}

func (Lib) JsonSerialization() {
	a := A{}
	obj, _ := json.Marshal(a)
	fmt.Println(string(obj))
}

// json反序列化时会去除结构体不存在的属性
func (Lib) JsonDeserialization() {
	obj := []byte(`{"a":"A","b":"B","c":"C"}`)
	a := A{}
	json.Unmarshal(obj, &a)
	fmt.Println(obj)
	fmt.Println(a)
}
