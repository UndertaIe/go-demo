package core

import (
	"context"
	"fmt"
)

func (c Core) Context() {
	ctx := context.Background()
	select {
	case <-ctx.Done():
	default:
	}
	fmt.Println("done")
}

func (c Core) WithCancel() {
	gen := func(ctx context.Context) <-chan int {
		dst := make(chan int)
		n := 1
		go func() {
			for {
				select {
				case <-ctx.Done():
					return // returning not to leak the goroutine
				case dst <- n:
					n++
				}
			}
		}()
		return dst
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	for n := range gen(ctx) {
		fmt.Println(n)
		if n == 5 {
			break
		}
	}
	fmt.Println(ctx.Deadline())
	fmt.Println(ctx.Err())
	cancel()
	fmt.Println(ctx.Deadline())
	fmt.Println(ctx.Err())
}

func (c Core) WithCancelOfContext() {
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel() // cancel when we are finished consuming integers

	dst := make(chan int)
	n := 1
	go func() {
		for {
			select {
			case <-ctx.Done():
				return
			case dst <- n:
				n++
			}
		}
	}()
	const x = 5
	for v := range dst {
		fmt.Println(v)
		if v == x {
			cancel() // or break
		}
	}
}
