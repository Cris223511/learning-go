// Acá veremos el paquete sync: herramientas para coordinar goroutines cuando
// necesitan compartir memoria en lugar de comunicarse por channels.

package main

import (
	"fmt"
	"sync"
)

// Sin Mutex, múltiples goroutines modificando contador al mismo tiempo
// producen una race condition: el resultado final es impredecible.
// Con Mutex, solo una goroutine a la vez puede modificar el valor.
type ContadorSeguro struct {
	mu    sync.Mutex
	valor int
}

func (c *ContadorSeguro) Incrementar() {
	c.mu.Lock()   // bloquea el acceso
	c.valor++
	c.mu.Unlock() // libera para que otra goroutine pueda entrar
}

func (c *ContadorSeguro) Valor() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.valor
}

// RWMutex: permite múltiples lecturas simultáneas pero solo una escritura a la vez.
// Ideal para cachés o registros que se leen mucho y se escriben poco.
type CachePolizas struct {
	mu   sync.RWMutex
	datos map[string]string
}

func (c *CachePolizas) Guardar(id, estado string) {
	c.mu.Lock() // escritura: exclusivo
	defer c.mu.Unlock()
	c.datos[id] = estado
}

func (c *CachePolizas) Obtener(id string) (string, bool) {
	c.mu.RLock() // lectura: múltiples goroutines pueden leer al mismo tiempo
	defer c.mu.RUnlock()
	v, ok := c.datos[id]
	return v, ok
}

// Once garantiza que una función se ejecute exactamente una vez,
// sin importar cuántas goroutines la llamen. Patrón singleton.
var (
	once       sync.Once
	instanciaDB string
)

func obtenerConexionDB() string {
	once.Do(func() {
		fmt.Println("  inicializando conexión a DB (solo ocurre una vez)")
		instanciaDB = "postgres://localhost:5432/aprendizago"
	})
	return instanciaDB
}

func main() {
	fmt.Println("== Mutex: acceso exclusivo ==")
	contador := &ContadorSeguro{}
	var wg sync.WaitGroup

	for range 100 {
		wg.Add(1)
		go func() {
			defer wg.Done()
			contador.Incrementar()
		}()
	}
	wg.Wait()
	fmt.Println("resultado esperado: 100 | obtenido:", contador.Valor())

	fmt.Println("\n== RWMutex: lecturas concurrentes ==")
	cache := &CachePolizas{datos: make(map[string]string)}
	cache.Guardar("POL-001", "activa")
	cache.Guardar("POL-002", "vencida")

	var wg2 sync.WaitGroup
	for range 5 {
		wg2.Add(1)
		go func() {
			defer wg2.Done()
			// Múltiples goroutines leyendo al mismo tiempo sin bloquearse entre sí.
			v, _ := cache.Obtener("POL-001")
			fmt.Printf("  leído: %s\n", v)
		}()
	}
	wg2.Wait()

	fmt.Println("\n== Once: inicialización única ==")
	var wg3 sync.WaitGroup
	for range 5 {
		wg3.Add(1)
		go func() {
			defer wg3.Done()
			conn := obtenerConexionDB()
			fmt.Println("  usando:", conn)
		}()
	}
	wg3.Wait()
}
