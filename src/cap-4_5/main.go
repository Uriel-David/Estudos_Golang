package main

import (
	"fmt"
	"runtime"
)

const constante = 10
const (
	testeA = iota
	testeB = iota
	testeC = iota
)
const (
	_  = iota             // 0
	KB = 1 << (iota * 10) // 1 << (1 * 10)
	MB = 1 << (iota * 10) // 1 << (2 * 10)
	GB = 1 << (iota * 10) // 1 << (3 * 10)
	TB = 1 << (iota * 10) // 1 << (4 * 10)
)

var testeConstante float64
var x bool

func main() {
	fmt.Println(x)

	x = true
	fmt.Println(x)

	x = (10 < 100)
	y := (10 == 100)
	z := (10 > 100)

	fmt.Println(x)
	fmt.Println(y)
	fmt.Println(z)

	lineVoid()

	a := "e"
	b := "é"
	c := "&"

	fmt.Printf("%v %v %v \n", a, b, c)
	fmt.Printf("%T %T %T \n", a, b, c)

	lineVoid()

	d := []byte(a)
	f := []byte(b)
	e := []byte(c)

	fmt.Printf("%v %v %v \n", d, e, f)
	fmt.Printf("%T %T %T \n", d, e, f)

	lineVoid()

	fmt.Println(runtime.GOOS)
	fmt.Println(runtime.GOARCH)

	lineVoid()

	var i uint16
	i = 65535 // Limit 65535
	fmt.Println(i)
	i++
	fmt.Println(i)

	lineVoid()

	str := "Hello, playground"
	fmt.Printf("%v - %T\n", str, str)
	strBytes := []byte(str)
	fmt.Printf("%v - %T\n", strBytes, strBytes)

	lineVoid()

	nums := 100
	fmt.Printf("%d - %b - %#x\n", nums, nums, nums)

	lineVoid()

	fmt.Println(constante)

	lineVoid()

	fmt.Println(testeA, testeB, testeC)

	lineVoid()

	fmt.Printf("%b \n", KB)
	fmt.Printf("%d \n", KB)
	fmt.Printf("%b \n", MB)
	fmt.Printf("%d \n", MB)
	fmt.Printf("%b \n", GB)
	fmt.Printf("%d \n", GB)
	fmt.Printf("%b \n", TB)
	fmt.Printf("%d \n", TB)
}

func lineVoid() {
	fmt.Println()
}
