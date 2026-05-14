// Acá veremos defer: una palabra clave que programa la ejecución de una función
// para cuando la función actual termine, sin importar cómo salga (normal o con error).

package main

import "fmt"

// defer se usa mucho para limpiar recursos: cerrar archivos, liberar conexiones,
// desbloquear mutexes. Así no se te olvida el cierre aunque haya un return anticipado.
func procesarPoliza(id string) {
	fmt.Println("abriendo recurso para:", id)
	// Este defer se ejecuta cuando procesarPoliza termina, sin importar qué pase después.
	defer fmt.Println("cerrando recurso para:", id)

	if id == "" {
		fmt.Println("id inválido, saliendo")
		return // el defer igual se ejecuta aquí
	}
	fmt.Println("procesando póliza:", id)
}

// Cuando hay varios defers, se ejecutan en orden inverso: el último en registrarse
// es el primero en ejecutarse. Es una pila (LIFO).
func ordenDefers() {
	defer fmt.Println("defer 1 - primero en declararse, último en ejecutarse")
	defer fmt.Println("defer 2")
	defer fmt.Println("defer 3 - último en declararse, primero en ejecutarse")
	fmt.Println("función ejecutándose...")
}

// Detalle importante: los argumentos del defer se evalúan en el momento de la declaración,
// no cuando se ejecuta. El valor de i queda capturado en ese instante.
func evaluacionInmediata() {
	i := 0
	defer fmt.Println("valor de i al momento del defer:", i) // imprime 0, no 10
	i = 10
	fmt.Println("i al final de la función:", i)
}

func main() {
	fmt.Println("== defer básico ==")
	procesarPoliza("POL-001")
	fmt.Println()
	procesarPoliza("")

	fmt.Println("\n== orden LIFO ==")
	ordenDefers()

	fmt.Println("\n== argumentos evaluados al momento ==")
	evaluacionInmediata()

	// Patrón real: defer para medir el tiempo de ejecución de una función.
	fmt.Println("\n== defer para logging/timing ==")
	defer fmt.Println("[FIN] función main terminada")
	fmt.Println("haciendo trabajo...")
	fmt.Println("más trabajo...")
}
