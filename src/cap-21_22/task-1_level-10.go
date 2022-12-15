package main

import (
	"fmt"
)

func main() {
	// solution 1
	c := make(chan int, 1)
	c <- 42

	fmt.Println(<-c)

	lineVoid()

	// solution 2
	d := make(chan int)

	go func() {
		d <- 42
	}()

	fmt.Println(<-d)
}
