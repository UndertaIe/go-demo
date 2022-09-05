package demo

import (
	"fmt"
	"strings"
)

func TestSliceJoin() {
	var s []string

	if s != nil {
		fmt.Println(strings.Join(s, ","))
	} else {
		fmt.Println("space")
	}

	s = append(s, "1")
	if s != nil {
		fmt.Println(strings.Join(s, ","))
	} else {
		fmt.Println("space")
	}

}
