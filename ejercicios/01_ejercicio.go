package ejercicios

import (
	"strconv"
)

func LaFuncion(param string) (int, string) {

	numero, error := strconv.Atoi(param)
	if error != nil {
		return -1, "No se puede convertir " + param + " a número entero.\n" + error.Error()
	}

	if numero > 100 {
		return numero, "El número " + strconv.Itoa(numero) + " es mayor a 100."
	}

	return numero, "El número " + strconv.Itoa(numero) + " es menor a 100"
}
