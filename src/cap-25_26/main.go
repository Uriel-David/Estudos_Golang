// Package main referente a documentação do código de estudos
package main

import (
	"Estudos_Golang/src/cap-25_26/dog"
	"fmt"
)

func main() {
	// go doc <pkg>
	// go doc <sym>[.<method>]
	// go doc [<pkg>.<sym>.<method>]
	// go doc [<pkg>.][<sym>.]<method>
	// go doc <pkg> <sym>[.<method>]

	fmt.Println(dog.Age(10))
}

// Função para imprimir uma linha vazia na comand line do terminal
func lineVoid() {
	fmt.Println()
}
