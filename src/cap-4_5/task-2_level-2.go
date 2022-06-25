package main

import "fmt"

func main() {
	a := 10 == 100
	b := 10 != 100
	c := 10 > 100
	d := 10 >= 100
	e := 10 < 100
	f := 10 <= 100

	fmt.Printf("%v %v %v %v %v %v \n", a, b, c, d, e, f)
}
