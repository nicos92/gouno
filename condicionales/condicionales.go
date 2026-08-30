package condicionales

import (
	"fmt"
	"runtime"
)

func VerCondicionales() {

	// os := runtime.GOOS

	if os := runtime.GOOS; os == "Linux." || os == "OS X." {
		fmt.Println("Esto no es Windows, es:", os)
	} else if os == "windows" {
		fmt.Println("Esto es", os)
	} else {
		fmt.Println("No se lo que es.", os)
	}

	switch os := runtime.GOOS; os {
	case "linux":
		fmt.Println("Esto es Linux")
	case "windows":
		fmt.Println("Esto es Windows")
	case "darwin":
		fmt.Println("Esto es Darwin")
	default:
		fmt.Printf("%s", os)
	}
}
