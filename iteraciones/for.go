package iteraciones

import "fmt"

func Iterar() {
	// i := 0
	// for i < 10 {
	// 	fmt.Println(i)
	// 	i++
	// }
	//
	for i := 1; i < 100; i *= 3 {
		if i > 9 {
			continue
		}
		fmt.Println(i)
		if i > 81 {
			break
		}
	}
}
