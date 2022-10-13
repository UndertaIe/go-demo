package core

import "fmt"

func (c Core) StringDemo() {
	s := "联合国人权理事会第五十一届会议期间"
	fmt.Println(len(s))
	for idx, c := range s {
		fmt.Println(idx, c)
	}
	fmt.Println(s[0:3])

	s2 := "123"
	fmt.Println(s2[0:3])
}
