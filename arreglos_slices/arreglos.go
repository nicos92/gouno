package arreglosslices

import "fmt"

var tabla [10]int = [10]int{10, 0, 59}
var matriz [20][30]int

func MuestroArreglos() {
	tabla[7] = 33
	tabla[2] = 22
	fmt.Println(tabla)

	tabla2 := [20]string{"Nicolás", "Sandoval"}
	fmt.Println(tabla2)

	for i := range len(tabla) {
		fmt.Print(tabla[i])
	}

	matriz[7][25] = 15
	fmt.Println(matriz)

	for i := 0; i < len(matriz); i++ {
		for j := 0; j < len(matriz[0]); j++ {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println("")
	}

	for i := range len(matriz) {
		for j := range len(matriz[0]) {
			fmt.Print(matriz[i][j], " ")
		}
		fmt.Println("")
	}
}
