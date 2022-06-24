package main

import "fmt"

type game int

var newGame game

func main() {
	fmt.Printf("newGame: %v %T\n", newGame, newGame)

	newGame = 42

	fmt.Printf("newGame: %v %T\n", newGame, newGame)
}
