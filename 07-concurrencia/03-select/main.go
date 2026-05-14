// Acá veremos select: permite esperar en múltiples channels al mismo tiempo
// y actuar con el primero que tenga datos disponibles.

package main

import (
	"fmt"
	"time"
)

func consultarBD(resultado chan<- string) {
	time.Sleep(80 * time.Millisecond)
	resultado <- "datos desde BD"
}

func consultarCache(resultado chan<- string) {
	time.Sleep(20 * time.Millisecond)
	resultado <- "datos desde caché"
}

func main() {
	// select elige el case cuyo channel esté listo primero.
	// Si varios están listos al mismo tiempo, elige uno al azar.
	fmt.Println("== select: el más rápido gana ==")
	bd := make(chan string, 1)
	cache := make(chan string, 1)
	go consultarBD(bd)
	go consultarCache(cache)

	select {
	case r := <-bd:
		fmt.Println("ganó BD:", r)
	case r := <-cache:
		fmt.Println("ganó caché:", r)
	}

	// Patrón de timeout: si ningún channel responde en el tiempo dado, tomamos otro camino.
	// time.After retorna un channel que recibe un valor después del tiempo indicado.
	fmt.Println("\n== Timeout ==")
	lento := make(chan string, 1)
	go func() {
		time.Sleep(200 * time.Millisecond)
		lento <- "respuesta tardía"
	}()

	select {
	case r := <-lento:
		fmt.Println("respuesta recibida:", r)
	case <-time.After(100 * time.Millisecond):
		fmt.Println("timeout: el servicio tardó demasiado")
	}

	// Default: hace que el select no bloquee. Si ningún channel está listo,
	// ejecuta el default inmediatamente. Útil para polling sin bloquear.
	fmt.Println("\n== select no bloqueante con default ==")
	chDatos := make(chan string, 1)

	select {
	case d := <-chDatos:
		fmt.Println("dato recibido:", d)
	default:
		fmt.Println("no hay datos disponibles ahora, continuando")
	}

	chDatos <- "póliza lista"
	select {
	case d := <-chDatos:
		fmt.Println("dato recibido:", d)
	default:
		fmt.Println("no hay datos")
	}

	// select en un bucle: procesa múltiples channels hasta recibir señal de stop.
	fmt.Println("\n== select en bucle con canal de stop ==")
	eventos := make(chan string, 5)
	stop := make(chan struct{})

	go func() {
		for _, e := range []string{"siniestro-001", "reclamo-002", "pago-003"} {
			eventos <- e
		}
		close(stop)
	}()

	for {
		select {
		case e, ok := <-eventos:
			if ok {
				fmt.Println("  evento:", e)
			}
		case <-stop:
			fmt.Println("  señal de stop recibida")
			goto fin
		}
	}
fin:
	fmt.Println("procesamiento terminado")
}
