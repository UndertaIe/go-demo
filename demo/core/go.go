package core

import (
	"fmt"
	"time"
)

func (Core) Go() {
	go func() {
		fmt.Println("goroutine is runing ")
		time.Sleep(time.Second)
		fmt.Println("goroutine finish")
	}()
	time.Sleep(time.Millisecond)
	fmt.Println("main finish")
}
