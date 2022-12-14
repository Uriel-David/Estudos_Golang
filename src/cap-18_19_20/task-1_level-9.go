package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func main() {
	newGoRoutines(20)
	wg.Wait()
}

func newGoRoutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		x := j
		go func(i int) {
			fmt.Println("I'm goroutines number:", i+1)
			wg.Done()
		}(x)
	}
}

func lineVoid() {
	fmt.Println()
}
