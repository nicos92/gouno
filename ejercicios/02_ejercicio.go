package ejercicios

import (
	"fmt"

	"github.com/nicos92/gouno/helpers"
)

func VerTablaMultiplicar() {
	var (
		numero int
		err    error
	)
	for {

		numero, err = helpers.LeerEntero("Ingrese un número entero:")

		if err != nil {
			helpers.CallClear()
			fmt.Println(err.Error())
			continue
		}
		break
	}

	multiplicarNumeroEntero(numero)
}

func multiplicarNumeroEntero(numero int) {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%d * %d = %d\n", numero, i, numero*i)

	}
}
