package ejercicios

import (
	"fmt"

	"github.com/nicos92/gouno/interfaces"
)

func HumanosRespirando(humano interfaces.Humano) {
	humano.Respirar()
	fmt.Printf("Soy un/a %s, y estoy respirando.\n", humano.Sexo())

}
