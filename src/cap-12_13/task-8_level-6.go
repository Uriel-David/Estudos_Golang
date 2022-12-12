package main

import (
	"fmt"
)

func main() {

	x := givesItBackfunc()

	x()
}

func givesItBackfunc() func() {
	return func() {
		fmt.Println("Hello I'm here!!")
	}
}

func lineVoid() {
	fmt.Println()
}
