package ejercicios

import (
	"fmt"

	"github.com/nicos92/gouno/helpers"
)

var (
	numero int
	err    error
	text   string
)

func VerTablaMultiplicarFile() string {
	for {

		numero, err = helpers.LeerEntero("Ingrese un número entero:")

		if err != nil {
			helpers.CallClear()
			fmt.Println(err.Error())
			continue
		}
		break
	}

	multiplicarNumeroEnteroFile()
	return text

}

func multiplicarNumeroEnteroFile() {
	for i := 1; i <= 10; i++ {
		text += fmt.Sprintf("%d * %d = %d\n", numero, i, numero*i)

	}
}
