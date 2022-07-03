package main

import (
	"fmt"
)

func main() {
	for x := 1; x <= 10000; x++ {
		fmt.Printf("%v", x)
		lineVoid()
	}
}

func lineVoid() {
	fmt.Println()
}
