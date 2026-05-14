// Acá veremos coverage: qué porcentaje del código está cubierto por los tests.
//
// Comandos clave:
//   go test ./08-testing/05-coverage -cover                      → porcentaje en terminal
//   go test ./08-testing/05-coverage -coverprofile=coverage.out  → genera archivo
//   go tool cover -html=coverage.out                             → abre reporte en navegador
//   go test ./... -cover                                         → coverage de todo el módulo

package main

import "fmt"

func main() {
	nivel, factor, _ := EvaluarRiesgo(30, 0)
	fmt.Printf("nivel: %s | factor: %.2f\n", nivel, factor)
}
