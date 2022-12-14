package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
	"time"
)

var wg sync.WaitGroup

func main() {
	/* fmt.Println(runtime.NumCPU())
	fmt.Println(runtime.NumGoroutine()) */

	totalGoRoutines := 10
	wg.Add(totalGoRoutines) // add (functions totals)

	/* go func1()
	go func2() */

	// fmt.Println(runtime.NumGoroutine())

	// var mu sync.Mutex

	var count int64
	// count := 0
	for i := 0; i < totalGoRoutines; i++ {
		go func() {
			// mu.Lock()
			// v := count
			runtime.Gosched()
			atomic.AddInt64(&count, 1)
			/* v++
			count = v */
			fmt.Println("Count:\t", atomic.LoadInt64(&count))
			// mu.Unlock()
			wg.Done()
		}()
	}

	wg.Wait() // wait!
	fmt.Println(count)
}

func func1() {
	for i := 0; i < 10; i++ {
		fmt.Println("func 1 exec:", i)
		time.Sleep(20)
	}
	wg.Done()
}

func func2() {
	for i := 0; i < 10; i++ {
		fmt.Println("func 2 exec:", i)
		time.Sleep(20)
	}
	wg.Done()
}

func lineVoid() {
	fmt.Println()
}
