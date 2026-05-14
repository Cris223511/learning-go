// Acá veremos los retornos múltiples de Go, que son la base del manejo de errores
// en el lenguaje. También veremos los retornos nombrados.

package main

import (
	"errors"
	"fmt"
)

// Go puede retornar varios valores en una sola función. El patrón más común es
// retornar el resultado y un error juntos. Si todo salió bien, el error va como nil.
func dividir(a, b float64) (float64, error) {
	if b == 0 {
		return 0, errors.New("no se puede dividir entre cero")
	}
	return a / b, nil
}

// Puedes buscar una póliza y retornar si existe o no, junto con los datos.
func buscarPoliza(id string) (string, bool) {
	polizas := map[string]string{
		"POL-001": "activa",
		"POL-002": "vencida",
	}
	estado, ok := polizas[id]
	return estado, ok
}

// Retornos nombrados: les pones nombre a los valores de retorno en la firma.
// La función puede usar return sin argumentos y Go devuelve los valores nombrados.
// Úsalos solo cuando mejoran la legibilidad, no por costumbre.
func calcularPrima(base, riesgo float64) (prima float64, igv float64, total float64) {
	prima = base * (1 + riesgo/100)
	igv = prima * 0.18
	total = prima + igv
	return
}

func main() {
	fmt.Println("== Resultado y error ==")
	// El patrón estándar: verificas el error antes de usar el resultado.
	resultado, err := dividir(100, 3)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("resultado: %.4f\n", resultado)
	}

	// Caso con error.
	_, err = dividir(10, 0)
	if err != nil {
		fmt.Println("error capturado:", err)
	}

	fmt.Println("\n== Dos valores (dato + bool) ==")
	if estado, ok := buscarPoliza("POL-001"); ok {
		fmt.Println("POL-001 estado:", estado)
	}
	if _, ok := buscarPoliza("POL-999"); !ok {
		fmt.Println("POL-999 no existe")
	}

	fmt.Println("\n== Retornos nombrados ==")
	prima, igv, total := calcularPrima(400, 15)
	fmt.Printf("prima: %.2f | IGV: %.2f | total: %.2f\n", prima, igv, total)

	// Puedes ignorar valores que no necesitas con _
	fmt.Println("\n== Ignorar con _ ==")
	_, igvSolo, _ := calcularPrima(500, 10)
	fmt.Printf("solo IGV: %.2f\n", igvSolo)
}
