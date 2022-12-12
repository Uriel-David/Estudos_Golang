package main

import (
	"fmt"
)

type person struct {
	name    string
	surname string
	age     int
}

func (someone person) show() {
	fmt.Println("The fullname that person is:", someone.name, someone.surname, "\n This person have", someone.age, "years.")
}

func main() {
	someone := person{
		name:    "Doze",
		surname: "Esquisita",
		age:     13,
	}

	someone.show()
}

func lineVoid() {
	fmt.Println()
}
