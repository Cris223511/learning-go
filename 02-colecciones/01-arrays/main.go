// Acá veremos los arrays en Go. El detalle más importante: el tamaño es parte del tipo,
// y los arrays se copian por valor, no por referencia como en otros lenguajes.

package main

import "fmt"

func main() {
	// El tamaño va entre corchetes y es fijo para siempre. [3]int y [4]int son tipos distintos,
	// no se pueden asignar uno al otro.
	var precios [3]float64
	precios[0] = 120.00
	precios[1] = 450.50
	precios[2] = 890.75

	// Forma corta: declaras e inicializas en una sola línea.
	productos := [4]string{"SOAT", "Vida", "Salud", "Vehicular"}

	// Si usas ... en lugar del número, Go cuenta los elementos y pone el tamaño solo.
	codigos := [...]int{101, 102, 103}

	fmt.Println("== Arrays ==")
	fmt.Println("precios:", precios)
	fmt.Println("productos:", productos)
	fmt.Println("codigos:", codigos)
	fmt.Println("largo de productos:", len(productos))

	// Los arrays se copian por valor. Modificar la copia no afecta al original.
	// Esto es distinto a Java o TypeScript donde los arrays son referencias.
	fmt.Println("== Copia por valor ==")
	original := [3]int{1, 2, 3}
	copia := original
	copia[0] = 99
	fmt.Println("original:", original) // [1 2 3], no cambia
	fmt.Println("copia:", copia)       // [99 2 3]

	// for-range sobre array entrega el índice y el valor.
	fmt.Println("== Recorrido ==")
	for i, p := range productos {
		fmt.Printf("[%d] %s -> S/ %.2f\n", i, p, precios[i%3])
	}

	// En la práctica, los arrays se usan poco directamente en Go.
	// Casi siempre trabajarás con slices, que son más flexibles.
}
