package main

import (
	"log"

	arreglosslices "github.com/nicos92/gouno/arreglos_slices"
	"github.com/nicos92/gouno/helpers"
)

func main() {
	// CallClear()
	helpers.CallClear()
	log.Printf("Iniciando main.\n\n")

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
	// files.SumaTabla()
	// files.LeoArchivo()
	//
	// funciones.Calculos()
	//
	// funciones.LlamarClosure()
	//
	// funciones.Exponencia(2)
	//
	arreglosslices.MuestroArreglos()
}
