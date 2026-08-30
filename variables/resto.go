package variables

import (
	"fmt"
	"strconv"
	"time"
)

var Nombre string
var Estado bool
var Sueldo float32
var Fecha time.Time
var Intde32 int32

func RestoVariables() {
	// MostrarEnteros()

	Nombre = "Pedro"
	Estado = true
	Sueldo = 1444.21
	Fecha = time.Now()
	Intde32 = 32
	fmt.Println(Nombre)
	fmt.Println(Estado)
	fmt.Println(Sueldo)
	fmt.Println(Fecha)
	fmt.Println(Intde32)
}

func ConviertoATexto(numero int) (bool, string) {
	texto := strconv.Itoa(numero)
	return true, texto
}
