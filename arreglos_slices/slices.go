package arreglosslices

import "fmt"

var tablaSlice []int = []int{2, 5, 4}
var arreglo [10]int = [10]int{6, 23, 5, 7, 45, 234, 6, 7, 8}

func MuestroSlice() {
	fmt.Println(tablaSlice)

	porcion := arreglo[3:] // Slice creado desde un vector, desde la posición 3

	porcion2 := arreglo[:5] // Slice creado desde un vector, desde la posición 0 hasta la 5

	porcion3 := arreglo[6:7] //Slice creado desde un vecto, desde la posición 6 hasta la 7

	fmt.Println(porcion)
	fmt.Println(porcion2)
	fmt.Println(porcion3)
}

func Capacidad() {
	elemento := make([]int, 5, 10)
	fmt.Printf("\nelemento: Largo %d, Capacidad %d", len(elemento), cap(elemento))

	nums := make([]int, 0)

	fmt.Printf("\nnums: Largo %d, Capacidad %d", len(nums), cap(nums))
	for i := range 100 {
		nums = append(nums, i)
	}

	fmt.Printf("\nnums: Largo %d, Capacidad %d", len(nums), cap(nums))
	for i := range 100 {
		nums = append(nums, i)
	}

	fmt.Printf("\nnums: Largo %d, Capacidad %d", len(nums), cap(nums))

}
