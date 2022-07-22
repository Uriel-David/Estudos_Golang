package main

import (
	"fmt"
)

var x [5]int
var y [6]int

func main() {
	x[0] = 1
	x[1] = 10
	fmt.Println(x[0], x[1])
	fmt.Println(x)
	fmt.Printf("%T\n", x)
	fmt.Printf("%T\n", y)
	fmt.Println(len(x))

	lineVoid()

	array := [5]int{1, 2, 3, 4, 5}
	fmt.Println(array)
	slice := []int{1, 2, 3, 4, 5}
	fmt.Println(slice)

	lineVoid()

	slice2 := append(slice, 6)
	fmt.Println(slice2)
	fmt.Println(slice[3])
	slice[3] = 348756
	fmt.Println(slice[3])
	// panic/runtime_error - slice/array não possui o indice 20
	/* slice[20] = 1
	fmt.Println(slice) */

	lineVoid()

	sliceRange := []string{"banana", "maçã", "uva"}
	for indice, valor := range sliceRange {
		fmt.Println("No índice", indice, "temos o valor:", valor)
	}

	sliceRange = append(sliceRange, "melancia")

	for _, valor := range sliceRange {
		fmt.Printf("Um dos valores desse slice é %v.\n", valor)
	}

	lineVoid()

	sliceRange2 := []int{1, 2, 3, 4, 5}
	total := 0

	for _, valor := range sliceRange2 {
		total += valor
	}

	fmt.Println("O valor desse slice é:", total)

	lineVoid()

	sabores := []string{"calabresa", "mozzarela", "4 queijos", "marguerita"}

	fatia := sabores[0:2]
	fmt.Println(fatia)

	lineVoid()

	for i := 0; i < len(sabores); i++ {
		fmt.Println(sabores[i])
	}

	lineVoid()

	sabores = append(sabores[:2], sabores[2:3]...)
	fmt.Println(sabores)

	lineVoid()

	oneSlice := []int{1, 2, 3, 4}
	otherSlice := []int{9, 10, 11, 12}

	fmt.Println(otherSlice)

	oneSlice = append(oneSlice, 5, 6, 7, 8)
	fmt.Println(oneSlice)

	oneSlice = append(oneSlice, otherSlice...)
	fmt.Println(oneSlice)

	lineVoid()

	makeSlice := make([]int, 5, 10)

	makeSlice[0], makeSlice[1], makeSlice[2], makeSlice[3], makeSlice[4] = 1, 2, 3, 4, 5

	makeSlice = append(makeSlice, 6)
	makeSlice = append(makeSlice, 7)
	makeSlice = append(makeSlice, 8)
	makeSlice = append(makeSlice, 9)
	makeSlice = append(makeSlice, 10)

	fmt.Println(makeSlice, len(makeSlice), cap(makeSlice))

	makeSlice = append(makeSlice, 11)

	fmt.Println(makeSlice, len(makeSlice), cap(makeSlice))

	lineVoid()

	SS := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}

	fmt.Println(SS[2][2])

	lineVoid()

	firstSlice := []int{1, 2, 3, 4, 5}
	fmt.Println(firstSlice)

	lineVoid()

	secondSlice := append(firstSlice[:2], firstSlice[4:]...)
	fmt.Println(secondSlice)
	fmt.Println(firstSlice)

	lineVoid()

	friend := map[string]int{
		"alfredo": 123456789,
		"joana":   987654321,
	}
	fmt.Println(friend)
	fmt.Println(friend["joana"])

	lineVoid()

	friend["gopher"] = 111000
	fmt.Println(friend)
	fmt.Println(friend["gopher"])

	lineVoid()

	if maybe, ok := friend["joazinho"]; !ok {
		fmt.Println("!Yo no conosco señor!")
	} else {
		fmt.Println(maybe)
	}

	lineVoid()

	whatever := map[int]string{
		123: "very cool",
		98:  "minus cool a little bit",
		983: "this is funny",
		18:  "Age go to party",
	}

	fmt.Println(whatever)

	lineVoid()

	for key, value := range whatever {
		fmt.Println(key, value)
	}

	lineVoid()

	delete(whatever, 123)
	fmt.Println(whatever)
}

func lineVoid() {
	fmt.Println()
}
