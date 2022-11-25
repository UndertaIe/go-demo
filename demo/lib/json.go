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
