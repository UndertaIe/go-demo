package core

import (
	"fmt"
	"time"
)

// select case 代码块随机 接收发送多个channel，防止饥饿问题
func (c Core) Select() {
	sender1 := make(chan int)
	go func() {
		for range time.NewTicker(time.Millisecond).C {
			sender1 <- 0
		}
	}()
	sender2 := make(chan int)
	go func() {
		time.Sleep(2 * time.Second)
		for range time.NewTicker(time.Millisecond).C {
			sender2 <- 0
		}
	}()

	for {
		select {
		case <-sender1:
			fmt.Println("case 1")
			time.Sleep(time.Second)
		case <-sender2:
			fmt.Println("case 2")
		}
	}
}

func (c Core) SelectNil() {
	var ch chan int
	if ch == nil {
		fmt.Println(ch)
	}
	select {
	case <-ch:
		println(1)
	}
}

func (Core) CloseChannel() {
	ch := make(chan int)
	go func() {
		ch <- 0
		ch <- 0
		close(ch)
	}()
	for range ch {
		fmt.Println(1)
	}
	fmt.Println("step 1 finish")

	ch2 := make(chan int)
	go func() {
		ch2 <- 0
		ch2 <- 0
		close(ch2) //关闭channel，接收channel则不阻塞
	}()
	i0, ok := <-ch2
	fmt.Println(i0, ok)

	i1, ok := <-ch2
	fmt.Println(i1, ok)

	i2, ok := <-ch2
	fmt.Println(i2, ok)
	fmt.Println("step 2 finish")
}
