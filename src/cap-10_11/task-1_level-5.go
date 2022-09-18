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
	person1 := person{
		name:    "Renata",
		surname: "Glass",
		tastes:  []string{"pistache", "morango", "baunilha"},
	}

	person2 := person{
		"Frederico",
		"da Prússia",
		[]string{"sabão em pó", "pé de coelho", "feijão"},
	}

	fmt.Println(person1)
	fmt.Println(person2)

	lineVoid()

	fmt.Println(person1.name, person1.surname)
	for _, v := range person1.tastes {
		fmt.Println("\t", v)
	}

	fmt.Println(person2.name, person2.surname)
	for _, v := range person2.tastes {
		fmt.Println("\t", v)
	}
}

func lineVoid() {
	fmt.Println()
}
