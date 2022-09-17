package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 = append(slice1[:3], slice1[len(slice1)-4:]...)
	fmt.Println(slice1)
}

func lineVoid() {
	fmt.Println()
}
