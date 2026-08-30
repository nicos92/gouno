package main

import (
	"fmt"

	"github.com/nicos92/gouno/variables"
)

func main() {

	// variables.MostrarEnteros()
	// variables.RestoVariables()
	estado, texto := variables.ConviertoATexto(1234)

	fmt.Println(estado)
	fmt.Println(texto)

}
