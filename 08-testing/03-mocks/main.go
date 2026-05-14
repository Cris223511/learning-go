// Acá veremos mocks con testify/mock: cómo testear servicios que dependen
// de interfaces (repositorios, clientes HTTP, etc.) sin tocar la BD real.
// Correr los tests: go test ./08-testing/03-mocks -v

package main

import "fmt"

func main() {
	fmt.Println("Módulo de mocks: correr con go test ./08-testing/03-mocks -v")
}
