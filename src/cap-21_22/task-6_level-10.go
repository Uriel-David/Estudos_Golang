package main

import "fmt"

func main() {
	c := make(chan int)
	go putHere(c)
	for v := range c {
		fmt.Println(v)
	}
}

func putHere(c chan int) {
	for i := 0; i < 100; i++ {
		c <- i
	}
	close(c)
}
