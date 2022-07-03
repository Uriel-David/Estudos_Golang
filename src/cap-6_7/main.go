package main

import (
	"fmt"
)

func main() {
	/* for horas := 0; horas <= 12; horas++ {
		fmt.Println("Hora:", horas)
		for x := 0; x < 60; x++ {
			fmt.Println(" ", x)
		}
		lineVoid()
	}
	*/

	/* for mes := 1; mes <= 12; mes++ {
		fmt.Println("Mês:", mes)
		for x := 1; x <= 30; x++ {
			fmt.Println("Dia:", x, " ")
		}
		lineVoid()
	} */

	//x := 0

	/* for x < 10 {
		fmt.Println(x)
		x++
	} */

	/* for {
		if x < 10 {
			x++
			fmt.Println("X menor que deiz...")
		} else {
			fmt.Println("X maior que deiz tchau...")
			break
		}
	}
	fmt.Println("Loop foi taxado na base do break") */

	/* for i := 0; i < 20; i++ {
		if i%2 != 0 {
			continue
		}
		fmt.Println(i)
	} */

	/* for i := 33; i <= 122; i++ {
		fmt.Printf("Decimal: %d - String: %v", i, string(i))
		lineVoid()
	} */

	//y := 10

	/* if !(y < 100) {
		fmt.Println("Hello IF condition")
	} */

	/* if y := 10; y < 100 {
		fmt.Println("Hello IF condition")
	} */

	/* if z := 10; z > 100 {
		fmt.Println("Z maior que 100")
	} else if z > 10 {
		fmt.Println("Z maior que 10")
	} else {
		fmt.Println("Z não é maior que 10 e nem 100")
	} */

	//a := 10

	/* switch {
	case a < 5:
		fmt.Println("X menor que 5")
	case a == 5:
		fmt.Println("X igual a 5")
	case a > 5:
		fmt.Println("X maior que 5")
	} */

	/* quemEstaNoEscritorio := "algoégedes"

	switch quemEstaNoEscritorio {
	case "joãozinho", "maria":
		fmt.Println("Quem está no escritório é o Time 1")
		// fallthrough não vem por padrão
		fallthrough
	case "marquinhos", "algoégedes":
		fmt.Println("Quem está no escritório é o Time 2")
	case "joana", "pedrinho":
		fmt.Println("Quem está no escritório é a Time 3")
	default:
		fmt.Println("Tá vazio ou a balada tava ótima ou é feriado")
	} */

	/* b := 4

	switch {
	case (b == 4), (b == 8):
		fmt.Println("é 4 ou 8")
	case (b < 10):
		fmt.Println("3 ou 4")
	default:
		fmt.Println("Default porque sim")
	} */

	/* var c interface{}
	c = -10

	switch c.(type) {
	case int:
		fmt.Println("É um int.")
	case bool:
		fmt.Println("É um bool.")
	case string:
		fmt.Println("É um string.")
	case float64:
		fmt.Println("É um float64.")
	} */

	/* switch d := 2; {
	case d == 1:
		fmt.Println("É 1.")
	case d == 2:
		fmt.Println("É 2.")
	case d == 3:
		fmt.Println("É 3.")
	case d == 4:
		fmt.Println("É 4.")
	} */

	e := 4

	/* if (e == 1) || (e == 2) {
		fmt.Println("'e' é dois ou um.")
	} */

	if (e%2 == 0) && !(e%3 == 0) {
		fmt.Println("'e' é um multiplo de dois e não de três.")
	}

	if (e%2 == 0) || (e%3 == 0) {
		fmt.Println("'e' é um multiplo de dois ou de três.")
	}
}

func lineVoid() {
	fmt.Println()
}
