package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg sync.WaitGroup
var count int

const totalOfGoroutines = 100

func main() {
	createGoroutines(totalOfGoroutines)
	wg.Wait()
	fmt.Println("Goroutines total:\t", totalOfGoroutines, "\nCount total:\t", count)
}

func createGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			v := count
			runtime.Gosched()
			v++
			count = v
			wg.Done()
		}()
	}
}

func lineVoid() {
	fmt.Println()
}
