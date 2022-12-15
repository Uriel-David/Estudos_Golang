package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	channel := make(chan int)

	go func() {
		channel <- 42
	}()

	fmt.Println(<-channel)

	lineVoid()

	channelSendRecieve := make(chan int)

	go send(channelSendRecieve)
	receive(channelSendRecieve)

	lineVoid()

	c := make(chan int)
	cr := make(<-chan int) // recieve
	cs := make(chan<- int) // send

	fmt.Println("---------")
	fmt.Println("c\t%T\n", c)
	fmt.Println("cr\t%T\n", cr)
	fmt.Println("cs\t%T\n", cs)
	fmt.Println("---------")

	// general to specific converts
	fmt.Println("---------")
	fmt.Println("c\t%T\n", (<-chan int)(c))
	fmt.Println("c\t%T\n", (chan<- int)(c))
	fmt.Println("---------")

	lineVoid()

	channel2 := make(chan int)
	go myLoop(10, channel2)

	/* for v := range channel2 {
		fmt.Println("Recieve of channel:", v)
	} */
	prints(channel2)

	lineVoid()

	a := make(chan int)
	b := make(chan int)
	x := 500

	go func(x int) {
		for i := 0; i < x; i++ {
			a <- i
		}
	}(x / 2)

	go func(x int) {
		for i := 0; i < x; i++ {
			b <- i
		}
	}(x / 2)

	for i := 0; i < x; i++ {
		select {
		case v := <-a:
			fmt.Println("Channel A:", v)
		case v := <-b:
			fmt.Println("Channel B:", v)
		}
	}

	lineVoid()

	quit := make(chan int)

	go receiveQuit(channel, quit)
	sendForChannel(channel, quit)

	lineVoid()

	par := make(chan int)
	impar := make(chan int)
	quit2 := make(chan bool)

	go sendNumbers(par, impar, quit2)
	recieveNumbers(par, impar, quit2)

	lineVoid()

	go func() {
		channel <- 42
		close(channel)
	}()

	v, ok := <-channel
	fmt.Println(v, ok)

	lineVoid()

	par2 := make(chan int)
	impar2 := make(chan int)
	converge := make(chan int)

	go send2(par2, impar2)
	go receive2(par2, impar2, converge)

	for v := range converge {
		fmt.Println("Value receive:", v)
	}

	lineVoid()

	channel3 := converge2(work("First"), work("Second"))
	for x := 0; x < 15; x++ {
		fmt.Println(<-channel3)
	}

	lineVoid()

	channel4 := make(chan int)
	channel5 := make(chan int)

	go giveMe(20, channel4)
	go other(channel4, channel5)

	for v := range channel5 {
		fmt.Println(v)
	}

	lineVoid()

	channel6 := make(chan int)
	channel7 := make(chan int)
	functions := 5

	go giveMe(50, channel6)
	go other2(functions, channel6, channel7)

	for v := range channel7 {
		fmt.Println(v)
	}

	lineVoid()

	ctx := context.Background()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("----------")

	ctx, cancel := context.WithCancel(ctx)

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")

	cancel()

	fmt.Println("context:\t", ctx)
	fmt.Println("context err:\t", ctx.Err())
	fmt.Printf("context type:\t%T\n", ctx)
	fmt.Println("cancel:\t\t", cancel)
	fmt.Printf("cancel type:\t%T\n", cancel)
	fmt.Println("----------")
}

func other2(functions int, channel1, channel2 chan int) {
	var wg sync.WaitGroup

	for i := 0; i < functions; i++ {
		wg.Add(1)
		go func() {
			for v := range channel1 {
				channel2 <- work2(v)
				wg.Done()
			}
		}()
	}
	wg.Wait()
	close(channel2)
}

func giveMe(n int, channel chan int) {
	for i := 0; i < n; i++ {
		channel <- i
	}
	close(channel)
}

func other(channel1, channel2 chan int) {
	var wg sync.WaitGroup

	for v := range channel1 {
		wg.Add(1)
		go func(x int) {
			channel2 <- work2(x)
			wg.Done()
		}(v)
	}
	wg.Wait()
	close(channel2)
}

func work2(n int) int {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
	return n
}

func work(s string) chan string {
	channel := make(chan string)

	go func(s string, c chan string) {
		for i := 0; ; i++ {
			c <- fmt.Sprintf("Function %v say: %v", s, i)
			time.Sleep(time.Millisecond * time.Duration(rand.Intn(1e3)))
		}
	}(s, channel)

	return channel
}

func converge2(x, y chan string) chan string {
	newC := make(chan string)

	go func() {
		for {
			newC <- <-x
		}
	}()

	go func() {
		for {
			newC <- <-y
		}
	}()

	return newC
}

func send2(p, i chan int) {
	x := 100
	for n := 0; n < x; n++ {
		if n%2 == 0 {
			p <- n
		} else {
			i <- n
		}
	}
	close(p)
	close(i)
}

func receive2(p, i, c chan int) {
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		for v := range p {
			c <- v
		}
		wg.Done()
	}()

	wg.Add(1)
	go func() {
		for v := range i {
			c <- v
		}
		wg.Done()
	}()
	wg.Wait()
	close(c)
}

func myLoop(t int, s chan<- int) {
	for i := 0; i < t; i++ {
		s <- i
	}
	close(s)
}

func prints(r <-chan int) {
	for v := range r {
		fmt.Println("Receive of channel:", v)
	}
}

func send(s chan<- int) {
	s <- 42
}

func receive(r <-chan int) {
	fmt.Println("The value receive:", <-r)
}

func receiveQuit(cr chan int, quit chan int) {
	for i := 0; i < 500; i++ {
		fmt.Println("Recieve:", <-cr)
	}
	quit <- 0
}

func sendForChannel(cs chan int, quit chan int) {
	anything := 10
	for {
		select {
		case cs <- anything:
			anything++
		case <-quit:
			return
		}
	}
}

func sendNumbers(par, impar chan int, quit chan bool) {
	total := 500
	for i := 0; i < total; i++ {
		if i%2 == 0 {
			par <- i
		} else {
			impar <- i
		}
	}
	close(par)
	close(impar)
	quit <- true
}

func recieveNumbers(par, impar chan int, quit chan bool) {
	for {
		select {
		case v := <-par:
			fmt.Println("The number:", v, "is par")
		case v := <-impar:
			fmt.Println("The number:", v, "is impar")
		case v, ok := <-quit:
			if !ok {
				fmt.Println("Comma ok:", v)
			} else {
				fmt.Println("Comma ok:", v)
			}
			return
		}
	}
}

func lineVoid() {
	fmt.Println()
}
