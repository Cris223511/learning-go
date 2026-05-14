// Acá veremos los unit tests en Go con testify. Los tests viven en archivos _test.go
// del mismo paquete y se corren con: go test ./08-testing/01-unit-tests -v

package main

import "fmt"

func main() {
	prima, _ := CalcularPrima(Vida, 100_000)
	fmt.Printf("prima Vida S/100k: S/%.2f\n", prima)

	conDescuento, _ := AplicarDescuento(prima, 10)
	fmt.Printf("con 10%% descuento: S/%.2f\n", conDescuento)
}
