package main

import "fmt"

type game int

var newGame game

var otherGame int

func main() {
	fmt.Printf("newGame: %v %T\n", newGame, newGame)

	newGame = 42

	fmt.Printf("newGame: %v %T\n", newGame, newGame)

	otherGame = int(newGame)

	fmt.Printf("otherGame: %v %T\n", otherGame, otherGame)
}
