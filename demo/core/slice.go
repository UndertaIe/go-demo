package core

import (
	"fmt"
)

func (c Core) RangeNilSlice() {
	var sli []int
	if sli == nil {
		fmt.Println("sli is nil")
	}
	for _, item := range sli {
		fmt.Println(item)
	}
} // ok

func (c Core) AppendToNilSlice() {
	var sli []int
	if sli == nil {
		fmt.Println("sli is nil")
	}
	sli = append(sli, 1)
	fmt.Println(sli)
} // ok

func (c Core) RangeKVofSlice() {
	sli := []int{1, 2, 3}
	for k := range sli {
		fmt.Print(k, " ")
	}
	fmt.Println()
	fmt.Println("==========")
	for k, v := range sli {
		fmt.Println(k, " ", v, " ")
	}
	fmt.Println("==========")
	for _, v := range sli {
		fmt.Println(v, " ")
	}
	fmt.Println("==========")
} // ok

func (c Core) BadCopy() {
	src := []int{1, 2, 3}
	var dst []int
	copy(dst, src)
	fmt.Println(src)
	fmt.Println(dst)
} // msg:  Copy returns the number of elements copied, which will be the minimum of len(src) and len(dst).

// [1 2 3]
// []

func (c Core) RightCopy() {
	src := []int{1, 2, 3}
	dst := make([]int, len(src))
	fmt.Println("current src: ", src)
	fmt.Println("current dst: ", dst)
	copy(dst, src)
	fmt.Println("after copy, src: ", src)
	fmt.Println("after copy, dst: ", dst)
} // [1 2 3]
// [1 2 3]

func (Core) Grow() {
	s := make([]int32, 3)
	fmt.Println(s, len(s), cap(s))
	s = append(s, 1, 2, 3, 4, 5, 6)
	fmt.Println(s, len(s), cap(s))

	s2 := make([]int, 3)
	fmt.Println(s2, len(s), cap(s2))
	s2 = append(s2, 1, 2, 3, 4, 5, 6)
	fmt.Println(s2, len(s2), cap(s2))
}

// 切片通过下标访问数组元素时不能下标大小不能大于slice len
func (Core) OutOfIndex() {
	s := make([]int, 3, 5)
	fmt.Println(s, cap(s), len(s))
	s[0] = 2
	fmt.Println(s, cap(s), len(s))
	s[4] = 2
	fmt.Println(s, cap(s), len(s))
} // panic: runtime error: index out of range [4] with length 3

// 通过切片创建的切片在未扩容前使用的是同一块内存区域，也就是指针指向了同一个数组
func (Core) InitWithArray() {
	arr := []int{1, 2, 3, 4, 5, 6}
	fmt.Println("arr: ", arr, cap(arr), len(arr)) //arr:  [1 2 3 4 5 6] 6 6
	s1 := arr[2:4]
	fmt.Println("s1: ", s1, cap(s1), len(s1)) // s1:  [3 4] 4 2
	s1[1] = 9
	fmt.Println("arr:", arr) // arr: [1 2 3 9 5 6]
	fmt.Println("s1: ", s1)  // s1:  [3 9]
	arr = append(arr, 10)    // arr扩容后，数组指针指向了跟s1不同的数组
	fmt.Println("arr: ", arr, cap(arr), len(arr))
	fmt.Println("s1: ", s1, cap(s1), len(s1))
}

// 传入参数的切片发生扩容的变化
func (Core) PassInSliceAfterGrow() {
	arr := []int{1, 2, 3}
	fmt.Println("outer arr before append: ", arr, cap(arr), len(arr)) // [1 2 3] 3 3
	appendx(&arr)
	fmt.Println("outer arr after append: ", arr, cap(arr), len(arr)) // [1 2 3 1 2 3 4 5 6 7] 10 10

	// 传递切片指针
	fmt.Println("=================")
	arr2 := []int{4, 5, 6}
	fmt.Println("outer arr2 before append: ", arr2, cap(arr2), len(arr2)) // [4 5 6] 3 3
	appendxx(arr2)
	fmt.Println("outer arr2 after append: ", arr2, cap(arr2), len(arr2)) // [4 5 6] 3 3

	fmt.Println("=================")
	// 不发生扩容时，调用appendxx前后切片的数组指针指向了同一个数组,
	// 虽然数组发生了变化，但arr3的len不会发生变化，即arr3的长度仍为5，即arr3不会发生变化，
	// 如果想让传入的切片在调用内部append作用于外部则应该使用切片指针
	arr3 := make([]int, 5, 20)
	fmt.Println("outer arr3 before append: ", arr3, cap(arr3), len(arr3)) // [0 0 0 0 0] 10 5
	appendxx(arr3)
	fmt.Println("outer arr3 after append: ", arr3, cap(arr3), len(arr3)) // [0 0 0 0 0 1 1 1] 10 5
	fmt.Println("arr3[0:5] =", arr3[0:5])
	fmt.Println("arr3[0:] =", arr3[0:])
	fmt.Println("arr3[0:cap(arr3)] =", arr3[0:cap(arr3)])
	fmt.Println("arr3[5:20] =", arr3[5:20]) // 可以通过下标重新生成一个切片，去访问全部的数组
	fmt.Println("arr3[:] =", arr3[:])
	fmt.Println("=================")
}

func appendx(s *[]int) {
	fmt.Println("inner appendx: ", s, cap(*s), len(*s)) // &[1 2 3] 3 3
	*s = append(*s, 1, 2, 3, 4, 5, 6, 7)
	fmt.Println("inner appendx: ", s, cap(*s), len(*s)) // &[1 2 3 1 2 3 4 5 6 7] 10 10
}

func appendxx(s []int) { // 复制的是 切片的引用
	fmt.Println("inner appendx: ", s, cap(s), len(s)) // [4 5 6] 3 3
	s = append(s, 1, 2, 3, 4, 5, 6, 7)                // 重新赋值的是新的切片，发生扩容则切片的数组指针字段指向了扩容后的数组内存地址
	fmt.Println("inner appendx: ", s, cap(s), len(s)) // [4 5 6 1 2 3 4 5 6 7] 10 10
}
