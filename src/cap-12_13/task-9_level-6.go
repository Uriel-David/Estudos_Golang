package main

import (
	"fmt"
)

func main() {
	thisFuncGivesItBackFunc(givesItBackfunc)
}

func givesItBackfunc() {
	fmt.Println("Hello I'm here!!")
}

func thisFuncGivesItBackFunc(x func()) {
	fmt.Println("Presentation: ")
	x()
}

func lineVoid() {
	fmt.Println()
}
