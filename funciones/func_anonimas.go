package funciones

import "fmt"

func Calculos() {

	var numero3, numero4 int
	// función anonima
	calculo := func(numero1 int, numero2 int) int {
		return numero1 + numero2
	}

	fmt.Println(calculo(10, 20))

	calculo = func(numero1 int, numero2 int) int {
		return numero1 - numero2 - numero3 - numero4
	}

	fmt.Println(calculo(10, 20))
}
