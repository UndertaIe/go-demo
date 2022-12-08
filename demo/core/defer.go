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
