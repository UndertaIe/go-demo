package lib

import (
	"fmt"
	"sort"
)

type SortSturct struct {
	I int
}

type sortarr []SortSturct

var _ sort.Interface = (*sortarr)(nil)

func (s sortarr) Len() int {
	return len(s)
}

// >表示左大右小， <表示左小右大
func (s sortarr) Less(i, j int) bool {
	return s[i].I >= s[j].I
}

func (s sortarr) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func (l Lib) SortDemo() {
	var arr sortarr
	arr = append(arr, SortSturct{I: 4}, SortSturct{I: 8}, SortSturct{I: 6})
	sort.Sort(arr)
	fmt.Println(arr)
}
