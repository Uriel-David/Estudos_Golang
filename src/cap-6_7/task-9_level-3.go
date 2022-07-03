package main

import (
	"fmt"
)

func main() {
	switch esporteFavorito := "Futebol"; {
	case esporteFavorito == "Basquete":
		fmt.Println(esporteFavorito)
	case esporteFavorito == "Futebol":
		fmt.Println(esporteFavorito)
	case esporteFavorito == "Volei":
		fmt.Println(esporteFavorito)
	}
}

func lineVoid() {
	fmt.Println()
}
