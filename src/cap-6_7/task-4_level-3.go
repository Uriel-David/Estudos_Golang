package main

import (
	"fmt"
)

func main() {
	yearCurrent := 2022
	birthday := 1998

	for {
		if birthday > yearCurrent {
			break
		}
		fmt.Println(birthday)
		birthday++
	}
}

func lineVoid() {
	fmt.Println()
}
