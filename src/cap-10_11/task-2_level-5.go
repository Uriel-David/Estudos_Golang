package main

import (
	"fmt"
)

type person struct {
	name    string
	surname string
	tastes  []string
}

func main() {
	myMepe1 := make(map[string]person)

	myMepe2 := map[string]person{
		"Pimentão": person{
			name:    "Renata",
			surname: "Pimentão",
			tastes:  []string{"pistache", "morango", "baunilha"},
		},
		"da Prússia": person{
			name:    "Frederico",
			surname: "da Prússia",
			tastes:  []string{"sabão em pó", "pé de coelho", "feijão"},
		},
	}

	myMepe1["Pimentão"] = person{
		name:    "Renata",
		surname: "Pimentão",
		tastes:  []string{"pistache", "morango", "baunilha"},
	}

	myMepe1["da Prússia"] = person{
		name:    "Frederico",
		surname: "da Prússia",
		tastes:  []string{"sabão em pó", "pé de coelho", "feijão"},
	}

	for _, value := range myMepe1 {
		fmt.Println("My name is ", value.name, "and my tastes favorites: ")
		for _, value := range value.tastes {
			fmt.Println(" - ", value)
		}
	}

	lineVoid()

	for _, value := range myMepe2 {
		fmt.Println("My name is ", value.name, "and my tastes favorites: ")
		for _, value := range value.tastes {
			fmt.Println(" - ", value)
		}
	}
}

func lineVoid() {
	fmt.Println()
}
