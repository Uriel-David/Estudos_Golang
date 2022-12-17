package main

import (
	"errors"
	"fmt"
	"log"
)

var ErrNorgatMath = errors.New("norgate math: square root of negative number")

type norgatMathError struct {
	lat  string
	long string
	err  error
}

func (n norgatMathError) Error() string {
	return fmt.Sprintf("a norgate math error ocurred: %v %v %v", n.lat, n.long, n.err)
}

func main() {
	/* var answer1, answer2 string

	fmt.Println("Name: ")
	_, err := fmt.Scan(&answer1)
	if err != nil {
		panic(err)
	}

	fmt.Println("Fav Food: ")
	_, err = fmt.Scan(&answer2)
	if err != nil {
		panic(err)
	}

	fmt.Println(answer1, answer2) */

	// lineVoid()

	// f, err := os.Create("names.txt")
	/* f, err := os.Open("names.txt")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	defer fmt.Println("Run a function with defer")

	r := strings.NewReader("James Bond")
	io.Copy(f, r)

	lineVoid()

	bs, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Println(string(bs)) */

	lineVoid()

	/* f2, err3 := os.Create("log.txt")
	if err3 != nil {
		fmt.Println(err3)
	}
	defer f2.Close()
	log.SetOutput(f2) */

	// defer foo()
	//_, err2 := os.Open("no-file.txt")
	//if err2 != nil {
	// fmt.Println("err happened", err2)
	// log.Println("err happened", err2)
	// log.Fatalln("err happened", err2)
	// log.Panicln(err2)
	// panic(err2)
	//}
	// defer f3.Close()

	// fmt.Println("check the log.txt file in the directory")

	/* f()
	fmt.Println("Returned normally from f.") */

	lineVoid()

	_, err := sqrt(-10)
	if err != nil {
		log.Fatalln(err)
	}
}

func sqrt(f float64) (float64, error) {
	if f < 0 {
		// ErrNorgatMath2 := fmt.Errorf("norgate math - square root of negative number: %v", f)
		nme := fmt.Errorf("norgate math - square root of negative number: %v", f)
		return 0, norgatMathError{"50.2289 N", "99.4656 W", nme}
	}

	return 42, nil
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}

	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}

func foo() {
	fmt.Println("When os.Exit() is called, deferred functions don't run")
}

func lineVoid() {
	fmt.Println()
}
