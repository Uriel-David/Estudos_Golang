package main

import (
	"fmt"
)

func main() {
	q := make(chan int)
	c := gen2(q)

	receiveTask4(c, q)

	fmt.Println("about to exit")
}

func gen2(q chan<- int) <-chan int {
	c := make(chan int)
	go func() {
		for i := 0; i < 100; i++ {
			c <- i
		}
		close(c)
		q <- 0
	}()
	return c
}

func receiveTask4(c <-chan int, q chan int) {
	for {
		select {
		case v := <-c:
			fmt.Println(v)
		case <-q:
			return
		}
	}
}
