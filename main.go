package main

import (
	"github.com/nicos92/gouno/files"
	"github.com/nicos92/gouno/helpers"
)

func main() {
	// CallClear()
	helpers.CallClear()

	// variables.MostrarEnteros()
	// variables.RestoVariables()
	// estado, texto := variables.ConviertoATexto(1234)

	// fmt.Println(estado)
	// fmt.Println(texto)
	//
	// condicionales.VerCondicionales()
	//
	// numero, mensaje, err := ejercicios.LaFuncion("1")
	// if err != nil {
	// 	fmt.Println("Ocurrio un error: ", err)
	// 	return
	// }
	// fmt.Println(numero)
	// fmt.Println(mensaje)
	//

	// teclado.IngresoNumeros()
	//
	// iteraciones.Iterar()
	//
	// text := ejercicios.VerTablaMultiplicarFile()
	// fmt.Println(text)
	//
	// files.GrabarTabla()
	//
	files.SumaTabla()
	files.LeoArchivo()
}
