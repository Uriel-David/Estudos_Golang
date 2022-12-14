package main

import "fmt"

type person struct {
	nome  string
	idade int
}

func (p *person) falar() {
	fmt.Println(p.nome, "say hello!")
}

type humans interface {
	falar()
}

func saySomething(h humans) {
	h.falar()
}

func main() {

	p1 := person{"Genghis Khan", 1000}

	p1.falar() // ← É um shortcut pra (&p1).falar()

	(&p1).falar() // ← É a maneira "mais correta."

	saySomething(&p1) // ← Funciona!

	// saySomething(p1) // ← Não funciona!
}

func lineVoid() {
	fmt.Println()
}
