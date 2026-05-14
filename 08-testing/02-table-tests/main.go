// Acá veremos table-driven tests: el patrón estándar de Go para tests.
// Correr los tests: go test ./08-testing/02-table-tests -v
// Correr un caso específico: go test -run TestValidarEmail/válido_básico

package main

import "fmt"

func main() {
	emails := []string{"user@mail.com", "invalid", "christopher@correo.com"}
	for _, e := range emails {
		err := ValidarEmail(e)
		if err != nil {
			fmt.Printf("inválido: %v\n", err)
		} else {
			fmt.Printf("válido: %s\n", e)
		}
	}
}
