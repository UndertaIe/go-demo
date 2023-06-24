package core

import (
	"fmt"
	"time"
)

// select的四种用法

// 阻塞当前goroutine，让出P使用权，无法被唤醒
func (Core) Select1() {
	fmt.Println("start to select")
	select {}
	// fmt.Println("never reaches the line!")
}

// 阻塞当前goroutine, 发送单一channel后继续执行
func (Core) Select2() {
	fmt.Println("start to select")
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		ch <- 999
	}()
	var i int
	select {
	case i = <-ch:
		fmt.Println("received data: ", i)
	}
	fmt.Println("end")
}

// 阻塞当前goroutine, 接收单一channel后继续执行
func (Core) Select3() {
	fmt.Println("start to select")
	ch := make(chan int)
	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("received data: ", <-ch)
	}()
	var i int = 999
	select {
	case ch <- i: // send data to channel
		fmt.Println("send data: ", i)
	}
	fmt.Println("end")
}

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

func (Core) BlankSelect() {
	select {}
}
