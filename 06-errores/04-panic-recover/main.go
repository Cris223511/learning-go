// Acá veremos panic y recover. La regla en Go es clara: los errores esperados
// se retornan, panic es solo para situaciones realmente irrecuperables.

package main

import "fmt"

// panic detiene la ejecución normal de la goroutine y empieza a deshacer el stack,
// ejecutando los defers que encuentre en el camino.
// Úsalo solo cuando el programa no puede continuar con seguridad: índice fuera de rango,
// nil pointer, configuración inválida al arrancar, divisiones por cero no detectadas.
func dividir(a, b int) int {
	if b == 0 {
		// Acá podríamos retornar error, pero lo usamos para demostrar panic.
		panic("división por cero: operación no permitida")
	}
	return a / b
}

// recover solo funciona dentro de un defer. Captura el valor del panic
// y permite que la goroutine continue de forma controlada.
// Este patrón es exactamente lo que hace el middleware de recuperación en Gin.
func ejecutarSeguro(fn func()) (err error) {
	defer func() {
		if r := recover(); r != nil {
			// r es el valor que se le pasó a panic (puede ser string, error, cualquier cosa).
			err = fmt.Errorf("recuperado de panic: %v", r)
		}
	}()
	fn()
	return nil
}

func cargarConfiguracion(path string) {
	if path == "" {
		// panic en inicialización es aceptable: si no hay config, el servicio no puede arrancar.
		panic("path de configuración vacío: el servicio no puede iniciar")
	}
	fmt.Println("configuración cargada desde:", path)
}

func main() {
	fmt.Println("== recover captura un panic ==")
	err := ejecutarSeguro(func() {
		fmt.Println("antes del panic")
		dividir(10, 0)
		fmt.Println("esto nunca se ejecuta")
	})
	if err != nil {
		fmt.Println("error capturado:", err)
	}
	fmt.Println("el programa sigue corriendo después del recover")

	fmt.Println("\n== Ejecución normal sin panic ==")
	err = ejecutarSeguro(func() {
		resultado := dividir(10, 2)
		fmt.Println("resultado:", resultado)
	})
	if err == nil {
		fmt.Println("sin errores")
	}

	fmt.Println("\n== Panic en inicialización ==")
	err = ejecutarSeguro(func() {
		cargarConfiguracion("config/app.yaml")
		cargarConfiguracion("") // este paniquea
	})
	if err != nil {
		fmt.Println("error:", err)
	}

	// Los defers se ejecutan incluso cuando hay un panic sin recover.
	// Fíjate el orden: los defers corren antes de que el panic mate la goroutine.
	fmt.Println("\n== Defers corren aunque haya panic ==")
	err = ejecutarSeguro(func() {
		defer fmt.Println("defer 1: siempre corre")
		defer fmt.Println("defer 2: también corre")
		panic("algo grave")
	})
	fmt.Println("recuperado:", err)
}
