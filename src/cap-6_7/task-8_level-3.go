package main

import (
	"fmt"
)

func main() {
	x := 1

	switch {
	case x == 0:
		fmt.Println(x)
	case x == 1:
		fmt.Println(x)
	case x == 2:
	}
}

func lineVoid() {
	fmt.Println()
}
