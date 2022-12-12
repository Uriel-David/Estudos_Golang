package main

import (
	"fmt"
)

type person struct {
	name    string
	surname string
	age     int
}

func main() {
	jackjack := person{"JackJack", "Silva", 10}
	fmt.Println(jackjack)

	changeMe(&jackjack)
	fmt.Println(jackjack)
}

func changeMe(p *person) {
	(*p).name = "ZéZé"
	p.surname = "Santos"
}

func lineVoid() {
	fmt.Println()
}
