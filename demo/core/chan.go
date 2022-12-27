package core

import (
	"fmt"
	"time"
)

func (Core) SelectNilChannel() {
	var ch chan int
	select {
	case <-ch:
		{
			fmt.Println("never reach this code")
		}
	}
}

// 拥有缓冲区的chan，发送者可以在发送n个数据后阻塞goroutine，直到有接收者消费后，发送者才可以继续发送直到n个继续阻塞goroutine
func (Core) NoBlockChannel() {
	ch := make(chan int, 1)
	ch <- 1
	fmt.Println("send data") // send data
	ch <- 1
	fmt.Println("send data 2") // never reach
}

func (Core) RangeNilChannel() {
	var ch chan int
	for range ch { // block nil chan
		fmt.Println("range")
	}
	fmt.Println("end") // never reach
}

// note: 不能关闭未初始化的channel
func (Core) RangeNilChannelWithClose() {
	var ch chan int
	go func() {
		time.Sleep(time.Second * 5)
		close(ch) // panic: close of nil channel
	}()
	for range ch {
		fmt.Println("range")
	}
	fmt.Println("end")
}

// range 在channel关闭后会退出for循环， 底层实现是 v,ok = next() 当 next()迭代channel没有数据接受时，ok为false时退出循环
func (Core) RangeChannelWithClose() {
	var ch chan int = make(chan int)
	go func() {
		time.Sleep(time.Second * 5)
		close(ch)
	}()
	for range ch { // block nil chan for five seconds
		fmt.Println("range")
	}
	fmt.Println("end") // end
}

// 接收未初始化的chan则会阻塞goroutine
func (Core) ReceiveNilChan() {
	var ch chan int
	fmt.Println(<-ch)  // block goroutine
	fmt.Println("end") // never reach
}

// len(channel) 返回当前channel中buffer存储元素个数
func (Core) ChannelLen() {
	ch := make(chan int, 10)
	fmt.Println(len(ch)) // 0
	ch <- 1
	fmt.Println(len(ch)) // 1
	ch <- 2
	fmt.Println(len(ch)) // 2
	<-ch
	fmt.Println(len(ch)) // 1
}
