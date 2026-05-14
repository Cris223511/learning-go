// Acá veremos benchmarks: cómo medir y comparar el rendimiento de distintas implementaciones.
// Correr benchmarks: go test ./08-testing/04-benchmarks -bench=. -benchmem
// Comparar dos funciones: go test -bench=BenchmarkConcatenar -benchmem

package main

import "fmt"

func main() {
	fmt.Println("con +:", len(ConcatenarConPlus(10)), "bytes")
	fmt.Println("con Builder:", len(ConcatenarConBuilder(10)), "bytes")
}
