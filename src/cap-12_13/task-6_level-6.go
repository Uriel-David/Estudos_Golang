package main

import (
	"fmt"
)

func main() {
	func() {
		fmt.Println("Hello Func!!")
	}()
}

func lineVoid() {
	fmt.Println()
}
