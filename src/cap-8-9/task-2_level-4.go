package main

import (
	"fmt"
)

func main() {
	slice := []int{1, 2, 3, 4, 5}
	for i, v := range slice {
		fmt.Println(i, v)
	}
}

func lineVoid() {
	fmt.Println()
}
