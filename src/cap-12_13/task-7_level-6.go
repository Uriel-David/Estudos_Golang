package main

import (
	"fmt"
)

func main() {
	x := func() {
		fmt.Println("Hello Func!!")
	}

	x()
}

func lineVoid() {
	fmt.Println()
}
