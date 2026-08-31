package goroutines

import (
	"fmt"
	"log"
	"strings"
	"time"
)

func MiNombreLentooo(nombre string, canal chan bool) {
	letras := strings.SplitSeq(nombre, "")

	for letra := range letras {
		time.Sleep(1000 * time.Millisecond)
		fmt.Print(letra)
	}
	log.Println("Final de la go routine")
	canal <- true
}
