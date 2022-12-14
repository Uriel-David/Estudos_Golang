package main

import (
	"encoding/json"
	"fmt"
	"os"
	"sort"

	"golang.org/x/crypto/bcrypt"
)

type person struct {
	Name    string  `json:"Name"`
	Surname string  `json:"Surname"`
	Age     int     `json:"Age"`
	Carrer  string  `json:"Carrer"`
	Wallet  float64 `json:"Wallet"`
}

type car struct {
	name    string
	force   int
	consume int
}

type sortForForce []car

func (a sortForForce) Len() int           { return len(a) }
func (a sortForForce) Less(i, j int) bool { return a[i].force < a[j].force }
func (a sortForForce) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type sortForConsume []car

func (a sortForConsume) Len() int           { return len(a) }
func (a sortForConsume) Less(i, j int) bool { return a[i].consume > a[j].consume }
func (a sortForConsume) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

type sortForProfitStationOwner []car

func (a sortForProfitStationOwner) Len() int           { return len(a) }
func (a sortForProfitStationOwner) Less(i, j int) bool { return a[i].consume < a[j].consume }
func (a sortForProfitStationOwner) Swap(i, j int)      { a[i], a[j] = a[j], a[i] }

func main() {
	jamesBond := person{
		Name:    "James",
		Surname: "Bond",
		Age:     40,
		Carrer:  "Secret Agent",
		Wallet:  1000000.50,
	}

	darthVader := person{
		Name:    "Anakin",
		Surname: "Skywalker",
		Age:     50,
		Carrer:  "Intergalact Vilain",
		Wallet:  50000000.83,
	}

	james, err := json.Marshal(jamesBond)
	if err != nil {
		fmt.Println(err)
	}

	darth, err := json.Marshal(darthVader)
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(string(james))
	fmt.Println(string(darth))

	lineVoid()

	sliceBytes := []byte(`{"Name":"James","Surname":"Bond","Age":40,"Carrer":"Secret Agent","Wallet":1000000.5}`)

	var persons person
	errUnmarshal := json.Unmarshal(sliceBytes, &persons)
	if errUnmarshal != nil {
		fmt.Println("error: ", errUnmarshal)
	}

	fmt.Println(persons)
	fmt.Println(persons.Wallet)

	lineVoid()

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(jamesBond)

	lineVoid()

	sliceString := []string{"banana", "coconuts", "apple", "orange", "grapple"}
	fmt.Println(sliceString)

	sort.Strings(sliceString)
	fmt.Println(sliceString)

	lineVoid()

	sliceInt := []int{12, 32, 1, 3, 9}
	fmt.Println(sliceInt)

	sort.Ints(sliceInt)
	fmt.Println(sliceInt)

	lineVoid()

	cars := []car{
		{"Porsche", 300, 12},
		{"Chevete", 100, 6},
		{"Fusca", 20, 20},
	}

	fmt.Println(cars)
	sort.Sort(sortForForce(cars))
	fmt.Println(cars)

	lineVoid()

	fmt.Println(cars)
	sort.Sort(sortForConsume(cars))
	fmt.Println(cars)

	lineVoid()

	fmt.Println(cars)
	sort.Sort(sortForProfitStationOwner(cars))
	fmt.Println(cars)

	lineVoid()

	password := "20julho1980"
	wrongPassword := "20julho1990"

	sliceBytes2, err := bcrypt.GenerateFromPassword([]byte(password), 10)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println(string(sliceBytes2))

	lineVoid()

	fmt.Println(bcrypt.CompareHashAndPassword(sliceBytes2, []byte(password)))
	fmt.Println(bcrypt.CompareHashAndPassword(sliceBytes2, []byte(wrongPassword)))
}

func lineVoid() {
	fmt.Println()
}
