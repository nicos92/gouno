package files

import (
	"bufio"
	"fmt"
	"os"

	"github.com/nicos92/gouno/ejercicios"
)

const FILE_NAME string = "./files/txts/tabla.txt"

func GrabarTabla() {
	var texto string = ejercicios.VerTablaMultiplicarFile()
	archivo, err := os.Create(FILE_NAME)
	if err != nil {
		fmt.Println("Error al crear el archivo", err.Error())
		return
	}
	fmt.Fprintln(archivo, texto)
	archivo.Close()

}

func SumaTabla() {
	var texto string = ejercicios.VerTablaMultiplicarFile()

	if !Append(texto) {
		fmt.Println("Error al concatenar contenido")
	}
}

func Append(texto string) bool {
	arch, err := os.OpenFile(FILE_NAME, os.O_WRONLY|os.O_APPEND, 0644)
	if err != nil {
		fmt.Println("Error durante el append. ", err.Error())
		return false
	}

	_, err = arch.WriteString(texto)
	if err != nil {
		fmt.Println("Error durante el writestring", err.Error())
		return false
	}
	arch.Close()
	return true
}

func LeoArchivo() {
	archivo, err := os.Open(FILE_NAME)

	if err != nil {
		fmt.Println("Error al leer el archivo.", err.Error())
		return
	}

	defer archivo.Close()

	scanner := bufio.NewScanner(archivo)
	for scanner.Scan() {
		fmt.Println(scanner.Text())
	}
	// Check for errors during scanning
	if err := scanner.Err(); err != nil {
		fmt.Println("Error durante el escaneo del archivo:", err.Error())
	}

}
