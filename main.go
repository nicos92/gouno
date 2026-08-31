package main

import (
	"fmt"
	"log"

	"github.com/nicos92/gouno/goroutines"
	"github.com/nicos92/gouno/helpers"
)

func main() {
	// CallClear()
	inicioMain()

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
	// arreglosslices.MuestroArreglos()
	//
	// arreglosslices.MuestroSlice()
	//
	// arreglosslices.Capacidad()
	//
	// mapas.MostrarMapas()
	//
	// users.AltaUsuario()
	// users.AltaUsuario()
	// users.AltaUsuario()
	// users.AltaUsuario()
	//
	// pedro := new(modelos.Hombre)
	// fmt.Println(pedro.EstaVivo())
	// ejercicios.HumanosRespirando(pedro)

	// mabel := new(modelos.Mujer)
	// fmt.Println(mabel.EstaVivo())
	// ejercicios.HumanosRespirando(mabel)
	//
	// deferspanic.VemosDefer()
	// deferspanic.EjemploPanic()
	//
	canal := make(chan bool)
	go goroutines.MiNombreLentooo("Nicolás", canal)
	defer func() {
		<-canal
	}()
	fmt.Println("Estoy acá")
	var x string
	fmt.Scanln(&x)

	finMain()
}

func inicioMain() {
	helpers.CallClear()
	log.Printf("Iniciando main. 🚀\n\n")
}
func finMain() {
	fmt.Printf("\n\n")
	log.Println("Fin de main. 🔚")

}
