package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

type person2 struct {
	name    string
	surname string
	age     int
}

type dentist struct {
	person2
	toothOut int
	salary   float64
}

type arquitect struct {
	person2
	typeBuild string
	heightMad string
}

type people interface {
	helloGoodMorning()
}

func main() {
	value := soma(10, 10)
	fmt.Println(value)

	lineVoid()

	args("Manhã")
	args("Tarde")
	args("Outra")

	lineVoid()

	total, many, hi := mutiplesReturns(10, 10, 1, 2, 3, 5)
	fmt.Println(total, many, hi)

	lineVoid()

	total2, many2, hi2 := variadic("Tarde", 10, 10, 1, 2, 3, 5)
	fmt.Println(total2, many2, hi2)

	lineVoid()

	si := []int{10, 10, 1, 2, 3, 5}
	total3 := soma2(si...)
	fmt.Println(total3)

	lineVoid()

	//defer fmt.Println("com defer, veio primeiro")
	fmt.Println("sem defer, veio depois")

	lineVoid()

	//defer fmt.Println("1")
	//defer fmt.Println("2")
	fmt.Println("3")
	fmt.Println("4")

	lineVoid()

	mauro := person{"Mauro", 30}
	mauro.helloGoodMorning()

	lineVoid()

	mrTooth := dentist{
		person2: person2{
			name:    "Mister Tooth",
			surname: "Da Silva",
			age:     50,
		},
		toothOut: 8973,
		salary:   3000,
	}
	mrBuild := arquitect{
		person2: person2{
			name:    "Mister Build",
			surname: "Da Silva",
			age:     51,
		},
		typeBuild: "Mad Builds",
		heightMad: "so much",
	}

	human(mrTooth)
	lineVoid()
	human(mrBuild)

	lineVoid()

	x := 387
	func(x int) {
		fmt.Println(x, "mutiple 873648 is:")
		fmt.Println(x * 873648)
	}(x)

	lineVoid()

	y := func(x int) int {
		//fmt.Println(x, "mutiple 873648 is:")
		//fmt.Println(x * 873648)
		return x * 873648
	}

	fmt.Println(x, "mutiple 873648 is:", y(x))
	//y(x)

	lineVoid()

	someValue := returnAnFunc()
	otherVariable := someValue(3)

	fmt.Println(otherVariable)
	fmt.Println(returnAnFunc()(3))

	lineVoid()

	t := onlyPairs(sum2, []int{50, 51, 52, 53, 54, 55, 56, 57, 58, 59, 60}...)
	fmt.Println(t)

	lineVoid()

	a := i()
	fmt.Println(a())
	fmt.Println(a())
	fmt.Println(a())

	b := i()
	fmt.Println(b())
	fmt.Println(b())
	fmt.Println(b())

	lineVoid()

	fmt.Println(fatorial(4))
	fmt.Println(loopsFatorial(4))
}

func fatorial(x int) int {
	if x == 1 {
		return x
	}
	return x * fatorial(x-1)
}

func loopsFatorial(x int) int {
	total := 1
	for x > 1 {
		total *= x
		x--
	}
	return total
}

func i() func() int {
	x := 0
	return func() int {
		x++
		return x
	}
}

func sum2(x ...int) int {
	n := 0
	for _, v := range x {
		n += v
	}
	return n
}

func onlyPairs(f func(x ...int) int, y ...int) int {
	var slice []int
	for _, v := range y {
		if v%2 == 0 {
			slice = append(slice, v)
		}
	}
	total := f(slice...)
	return total
}

func returnAnFunc() func(int) int {
	return func(i int) int {
		return i * 10
	}
}

func human(p people) {
	p.helloGoodMorning()

	switch p.(type) {
	case dentist:
		fmt.Println("I gained", p.(dentist).salary)
	case arquitect:
		fmt.Println("I building", p.(arquitect).typeBuild)
	}
}

func (x dentist) helloGoodMorning() {
	fmt.Println("My name is", x.name, ".I'm pulled out", x.toothOut, "tooth.\nSay good moorning!")
}

func (x arquitect) helloGoodMorning() {
	fmt.Println("My name is", x.name, "\nSay good moorning!")
}

func (p person) helloGoodMorning() {
	fmt.Println(p.name, "say good moorning!")
}

func soma2(x ...int) int {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum
}

func variadic(s string, x ...int) (int, int, string) {
	hi2 := ""
	if s == "Manhã" {
		hi2 = "Manhã"
	} else if s == "Tarde" {
		hi2 = "Tarde"
	} else {
		hi2 = "Noite"
	}

	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum, len(x), hi2
}

func mutiplesReturns(x ...int) (int, int, string) {
	sum := 0
	for _, v := range x {
		sum += v
	}
	return sum, len(x), "Good Moorning!"
}

func args(s string) {
	if s == "Manhã" {
		fmt.Println("Manhã")
	} else if s == "Tarde" {
		fmt.Println("Tarde")
	} else {
		fmt.Println("Noite")
	}
}

func soma(x, y int) int {
	return x + y
}

func lineVoid() {
	fmt.Println()
}
