package middleware

import "fmt"

func sumar(a int, b int) int {
	return a + b

}

func restar(a int, b int) int {
	return a - b

}
func multiplicar(a int, b int) int {
	return a * b

}

func MiMiddleware() {
	fmt.Println("Inicio")
	resultSumar := operacionesMidd(sumar)(2, 4)

	fmt.Println(resultSumar)
	resultRestar := operacionesMidd(restar)(2, 4)

	fmt.Println(resultRestar)
	resultMultiplicar := operacionesMidd(multiplicar)(2, 4)

	fmt.Println(resultMultiplicar)
}

func operacionesMidd(fun func(a int, b int) int) func(int, int) int {
	return func(a int, b int) int {
		fmt.Println("Middleware: Inicio de operacion ")
		return fun(a, b)
	}
}
