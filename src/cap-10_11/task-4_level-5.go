package main

import (
	"fmt"
)

func main() {
	x := struct {
		name         string
		taste        string
		whereIs      []string
		goesWellWith map[string]string
	}{
		name:    "Stroopwafel",
		taste:   "Sweet",
		whereIs: []string{"Netherlands", "Patinhas Uncle"},
		goesWellWith: map[string]string{
			"breafast": "tea",
			"lunch":    "cooffe",
			"dinner":   "nothing or cooffe",
		},
	}

	fmt.Println(x)
}

func lineVoid() {
	fmt.Println()
}
