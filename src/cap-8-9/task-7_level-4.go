package main

import (
	"fmt"
)

func main() {
	ss := [][]string{
		[]string{
			"meow",
			"meru",
			"miar",
		},
		[]string{
			"mimi",
			"mary",
			"eat",
		},
		[]string{
			"me-name",
			"me-surname",
			"me-hobbie",
		},
	}

	for _, v := range ss {
		fmt.Println(v[0])
		for _, thing := range v {
			fmt.Println("\t", thing)
		}
	}
}

func lineVoid() {
	fmt.Println()
}
