// Acá veremos goroutines: la unidad de concurrencia de Go. Son funciones que
// corren de forma independiente, mucho más livianas que un thread del sistema.

package main

import (
	"fmt"
	"sync"
	"time"
)

func procesarPoliza(id string, wg *sync.WaitGroup) {
	// defer wg.Done() le avisa al WaitGroup que esta goroutine terminó.
	// El defer garantiza que se llame aunque la función salga con error.
	defer wg.Done()
	time.Sleep(50 * time.Millisecond) // simula trabajo real (consulta a BD, llamada externa)
	fmt.Printf("  póliza %s procesada\n", id)
}

func main() {
	fmt.Println("== Goroutine básica ==")
	// La palabra go lanza la función en una goroutine separada.
	// main no espera: si termina antes, las goroutines mueren con ella.
	go fmt.Println("  goroutine lanzada")
	time.Sleep(10 * time.Millisecond) // solo para este ejemplo; en producción se usa WaitGroup
	fmt.Println("  main continúa")

	// WaitGroup es la forma correcta de esperar a que un grupo de goroutines termine.
	// Add(n) registra cuántas hay. Done() descuenta una. Wait() bloquea hasta llegar a cero.
	fmt.Println("\n== WaitGroup: esperar a todas ==")
	var wg sync.WaitGroup
	polizas := []string{"POL-001", "POL-002", "POL-003", "POL-004", "POL-005"}

	for _, id := range polizas {
		wg.Add(1)
		go procesarPoliza(id, &wg) // &wg porque WaitGroup no se copia
	}
	wg.Wait() // bloquea hasta que todas las goroutines llamen Done()
	fmt.Println("todas las pólizas procesadas")

	// Las goroutines corren concurrentemente: el orden de salida no está garantizado.
	// En cada ejecución pueden terminar en distinto orden.
	fmt.Println("\n== Orden no garantizado ==")
	var wg2 sync.WaitGroup
	for i := 1; i <= 5; i++ {
		wg2.Add(1)
		i := i // captura el valor actual de i para cada goroutine (importante)
		go func() {
			defer wg2.Done()
			fmt.Printf("  goroutine %d\n", i)
		}()
	}
	wg2.Wait()
}
