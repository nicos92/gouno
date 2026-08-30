package helpers

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func LeerEntero(mensaje string) (int, error) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(mensaje)

	if !scanner.Scan() {
		if err := scanner.Err(); err != nil {
			return 0, err
		}
		return 0, fmt.Errorf("lectura de entrada finalizada")
	}

	entrada := strings.TrimSpace(scanner.Text())

	numero, err := strconv.Atoi(entrada)
	if err != nil {
		// return 0, fmt.Errorf("ingreso no válido: %w", err)
		return 0, fmt.Errorf("Ingreso no válido")
	}
	return numero, nil
}
