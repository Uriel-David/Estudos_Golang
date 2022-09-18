package main

import (
	"fmt"
)

type client struct {
	name    string
	surname string
	smoker  bool
}

type person struct {
	name string
	age  int
}

type professional struct {
	person
	title  string
	salary int
}

func main() {
	client1 := client{
		name:    "joãozinho",
		surname: "da silva",
		smoker:  false,
	}

	client2 := client{"Joana", "Pereira", true}

	fmt.Println(client1)
	fmt.Println(client2)

	lineVoid()

	person1 := person{
		name: "Alfredo",
		age:  30,
	}

	person2 := professional{
		person: person{
			name: "Maricota",
			age:  31,
		},
		title:  "Pizzaiola",
		salary: 10000,
	}

	fmt.Println(person1)
	fmt.Println(person2)

	lineVoid()

	fmt.Println(person1.name)
	fmt.Println(person2.title)
	fmt.Println(person2.person.name)
	fmt.Println(person2.age)

	lineVoid()

	person3 := person{"Mauricio", 40}
	person4 := professional{person{"Vanderlei", 50}, "Politico", 10000000}

	fmt.Println(person3)
	fmt.Println(person4)

	lineVoid()

	x := struct {
		name string
		age  int
	}{
		name: "Mario",
		age:  50,
	}

	fmt.Println(x)
}

func lineVoid() {
	fmt.Println()
}
