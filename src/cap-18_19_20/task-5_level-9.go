package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

var wg sync.WaitGroup
var count int32

const totalOfGoroutines = 100

func main() {

	criarGoroutines(totalOfGoroutines)
	wg.Wait()
	fmt.Println("Goroutines total:\t", totalOfGoroutines, "\nCount total:\t", count)

}

func criarGoroutines(i int) {
	wg.Add(i)
	for j := 0; j < i; j++ {
		go func() {
			// func AddInt32(addr *int32, delta int32) (new int32)
			// func LoadInt32(addr *int32) (val int32)
			atomic.AddInt32(&count, 1)
			v := atomic.LoadInt32(&count)
			runtime.Gosched()
			fmt.Println(v)
			wg.Done()
		}()
	}
}

func lineVoid() {
	fmt.Println()
}
