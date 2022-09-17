package main

import (
	"fmt"
)

func main() {
	slice1 := []int{1, 2, 3, 4, 5, 6, 7, 8, 9}
	slice1 = append(slice1, 10)
	slice1 = append(slice1, 11, 12, 13)
	slice2 := []int{14, 15, 16, 17, 18, 19, 20}
	slice1 = append(slice1, slice2...)
	fmt.Println(slice1)
}

func lineVoid() {
	fmt.Println()
}
