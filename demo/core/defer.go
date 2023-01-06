package core

import (
	"fmt"
	"time"
)

// defer 参数复制时机：是在使用defer关键字时就创建好了的
func (c Core) Defer() {
	now := time.Now()
	defer fmt.Println(time.Since(now))
	time.Sleep(time.Second)
} // 5.8μs

// defer 和 return 的先后顺序:
//		return分两步，第一步复制返回值存入栈中；
// 		第二步执行return跳转；
//		而在函数中的所有的defer函数是在第一步和第二步之间执行 来源: GO专家编程
// 类似于defer传参，参数在执行之前就已经写入栈中，之后再修改参数也不会影响到defer函数的参数
func (c Core) DeferReturnSequence() {
	i := _DeferReturnSequence()
	fmt.Println("i: ", i)

	j := _DeferReturnSequence2()
	fmt.Println("j: ", *j)
} // defer block i:  2
// i:  1
// defer block j:  2
// j:  2

func _DeferReturnSequence() int {
	i := 1
	defer func() {
		i++
		fmt.Println("defer block i: ", i)
	}()
	return i
}

func _DeferReturnSequence2() *int {
	i := 1
	var j *int = &i
	defer func() {
		*j += 1
		fmt.Println("defer block j: ", *j)
	}()
	return j
}
