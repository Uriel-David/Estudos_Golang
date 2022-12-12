package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("Vem depois")
	fmt.Println("Vem antes")
}

func lineVoid() {
	fmt.Println()
}
