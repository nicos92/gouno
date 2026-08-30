package ejercicios

import (
	"fmt"
	"strconv"
)

func LaFuncion(param string) (int, string, error) {

	numero, err := strconv.Atoi(param)
	if err != nil {
		// Devolvemos 0, texto vacío y el error real
		return 0, "", fmt.Errorf("error de conversión en param '%s': %w", param, err)
	}

	if numero > 100 {
		return numero, "El número " + strconv.Itoa(numero) + " es mayor a 100.", err
	}
	if numero == 100 {
		return numero, "El número " + strconv.Itoa(numero) + " es igual a 100.", err
	}

	return numero, "El número " + strconv.Itoa(numero) + " es menor a 100", err
}
