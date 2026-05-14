// Acá veremos los closures: funciones que recuerdan el entorno donde fueron creadas.
// Son la base de patrones como generadores, middleware y funciones configurables.

package main

import "fmt"

// Esta función retorna otra función. La función retornada "cierra" sobre la variable
// contador: la recuerda y la modifica cada vez que la llamas. Eso es un closure.
func nuevoContador() func() int {
	contador := 0
	return func() int {
		contador++
		return contador
	}
}

// Genera una función que aplica siempre el mismo porcentaje de descuento.
// Útil para crear variantes preconfiguradas sin repetir lógica.
func generarDescuento(porcentaje float64) func(float64) float64 {
	return func(precio float64) float64 {
		return precio - (precio * porcentaje / 100)
	}
}

// Closure con estado mutable: acumula primas y lleva un historial interno.
func nuevaCalculadora() func(float64) (float64, int) {
	total := 0.0
	cantidad := 0
	return func(prima float64) (float64, int) {
		total += prima
		cantidad++
		return total, cantidad
	}
}

func main() {
	fmt.Println("== Contador con closure ==")
	contar := nuevoContador()
	fmt.Println(contar()) // 1
	fmt.Println(contar()) // 2
	fmt.Println(contar()) // 3

	// Cada llamada a nuevoContador() crea su propio contador independiente.
	otroContador := nuevoContador()
	fmt.Println(otroContador()) // 1, no continúa desde 3

	fmt.Println("\n== Descuentos preconfigurados ==")
	descuento10 := generarDescuento(10)
	descuento25 := generarDescuento(25)
	fmt.Printf("SOAT S/120 con 10%%: S/%.2f\n", descuento10(120))
	fmt.Printf("Vida S/450 con 25%%: S/%.2f\n", descuento25(450))

	fmt.Println("\n== Calculadora acumuladora ==")
	calcular := nuevaCalculadora()
	total, n := calcular(120.00)
	fmt.Printf("primas: %d | acumulado: S/%.2f\n", n, total)
	total, n = calcular(450.50)
	fmt.Printf("primas: %d | acumulado: S/%.2f\n", n, total)
	total, n = calcular(890.75)
	fmt.Printf("primas: %d | acumulado: S/%.2f\n", n, total)

	// Función anónima ejecutada inmediatamente (IIFE).
	// Útil para aislar un bloque de lógica sin contaminar el scope.
	fmt.Println("\n== Función anónima inmediata ==")
	resultado := func(a, b float64) float64 {
		return a * b
	}(450.50, 1.18)
	fmt.Printf("precio con IGV: S/%.2f\n", resultado)
}
