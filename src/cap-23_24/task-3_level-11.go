package main

import "fmt"

type specialError struct {
	anything string
}

func (e specialError) Error() string {
	return fmt.Sprintf("deu zica! e olha isso: %v", e.anything)
}

func paramWithError(e error) {
	fmt.Println("Passaram como argumento pra mim, o seguinte: ", e.(specialError).anything, "\nno método Error, eu tenho:", e)
}

func main() {
	arg := specialError{"EMOJIS!!!!!!!!"}
	paramWithError(arg)
}
