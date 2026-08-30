package teclado

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var (
	numero1 int
	numero2 int
	leyenda string
	err     error
)

func IngresoNumeros() {
	miScanner := bufio.NewScanner(os.Stdin)

	fmt.Println("Ingrese número 1: ")
	if miScanner.Scan() {
		numero1, err = strconv.Atoi(miScanner.Text())
		if err != nil {
			// si queremos abortar completamenta la aplicacion lanzamos un panic
			panic("El Dato ingresado es incorrecto. " + err.Error())
		}
	}

	fmt.Println("Ingrese número 2: ")
	if miScanner.Scan() {
		numero2, err = strconv.Atoi(miScanner.Text())
		if err != nil {
			// si queremos abortar completamenta la aplicacion lanzamos un panic
			panic("El Dato ingresado es incorrecto. " + err.Error())
		}
	}

	fmt.Println("Ingrese leyenda: ")
	if miScanner.Scan() {
		leyenda = miScanner.Text()
	}

	fmt.Println(leyenda, numero1*numero2)
}
