// Acá veremos channels: el mecanismo de Go para comunicar goroutines.
// La filosofía es "no compartas memoria, comunica mediante channels".

package main

import (
	"fmt"
	"sync"
)

func calcularPrima(polizaID string, ch chan<- string) {
	// ch chan<- string es un channel de solo escritura. Buena práctica en firmas de funciones.
	prima := 450.50
	ch <- fmt.Sprintf("%s: S/%.2f", polizaID, prima)
}

func generarIDs(ids []string, ch chan<- string) {
	for _, id := range ids {
		ch <- id // envía al channel, bloquea si está lleno
	}
	close(ch) // cerrar avisa a los receptores que no habrá más datos
}

func main() {
	// Channel sin buffer: el envío bloquea hasta que alguien lea.
	// Sincroniza automáticamente al emisor y receptor.
	fmt.Println("== Channel sin buffer ==")
	ch := make(chan string)
	go calcularPrima("POL-001", ch)
	resultado := <-ch // bloquea hasta recibir
	fmt.Println("recibido:", resultado)

	// Channel con buffer: el envío no bloquea hasta llenar el buffer.
	// Útil cuando el productor es más rápido que el consumidor.
	fmt.Println("\n== Channel con buffer ==")
	chBuf := make(chan int, 3)
	chBuf <- 100
	chBuf <- 200
	chBuf <- 300
	// Los tres enviados sin bloquear porque el buffer tiene capacidad 3.
	fmt.Println(<-chBuf, <-chBuf, <-chBuf)

	// Patrón fan-out: un productor, varios consumidores.
	fmt.Println("\n== Fan-out: múltiples consumidores ==")
	ids := []string{"POL-001", "POL-002", "POL-003", "POL-004"}
	chIds := make(chan string, len(ids))
	go generarIDs(ids, chIds)

	var wg sync.WaitGroup
	for workerID := range 2 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			// for-range sobre channel lee hasta que el channel se cierre.
			for id := range chIds {
				fmt.Printf("  worker %d procesó %s\n", workerID, id)
			}
		}()
	}
	wg.Wait()

	// Patrón pipeline: una goroutine produce, otra transforma, otra consume.
	fmt.Println("\n== Pipeline ==")
	numeros := make(chan int, 5)
	dobles := make(chan int, 5)

	go func() {
		for i := 1; i <= 5; i++ {
			numeros <- i
		}
		close(numeros)
	}()

	go func() {
		for n := range numeros {
			dobles <- n * 2
		}
		close(dobles)
	}()

	for d := range dobles {
		fmt.Printf("  %d\n", d)
	}
}
