package mapas

import "fmt"

func MostrarMapas() {
	paisProvincia := make(map[string]string)
	fmt.Println(paisProvincia)

	paisProvincia["Mexico"] = "D.F."
	paisProvincia["Argentina"] = "Buenos Aires"
	fmt.Println(paisProvincia)
	fmt.Println(paisProvincia["Argentina"])

	campeonato := map[string]int{
		"Barcelona":    39,
		"Real Madrid":  38,
		"Chivas":       29,
		"Boca Juniors": 44,
	}

	fmt.Println(campeonato)

	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d.\n", equipo, puntaje)
	}

	delete(campeonato, "Chivas")
	fmt.Println("Equipo chivas eliminado")
	for equipo, puntaje := range campeonato {
		fmt.Printf("Equipo %s, tiene un puntaje de %d.\n", equipo, puntaje)
	}

	puntaje, existe := campeonato["Boca Juniors"]

	fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntaje, existe)
	puntajeDos, existeDos := campeonato["River Plate"]

	if !existe {

		fmt.Printf("El puntaje capturado es %d, y el equipo existe = %t \n", puntajeDos, existeDos)
	} else {
		fmt.Println("El equipo no existe")
	}
}
