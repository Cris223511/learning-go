// Benchmarks en Go: funciones que empiezan con Benchmark y reciben *testing.B.
// Se corren con: go test ./08-testing/04-benchmarks -bench=. -benchmem
// -bench=. corre todos los benchmarks. -benchmem muestra allocations de memoria.

package main

import (
	"fmt"
	"testing"
)

// b.N es el número de iteraciones que el runner ajusta automáticamente
// hasta que el resultado sea estadísticamente estable.
func BenchmarkConcatenarConPlus(b *testing.B) {
	for range b.N {
		ConcatenarConPlus(100)
	}
}

func BenchmarkConcatenarConBuilder(b *testing.B) {
	for range b.N {
		ConcatenarConBuilder(100)
	}
}

// Para benchmarks que necesitan setup, usa b.ResetTimer() después del setup
// para que el tiempo de preparación no cuente en la medición.
func BenchmarkBuscarLineal(b *testing.B) {
	ids := make([]string, 1000)
	for i := range 1000 {
		ids[i] = fmt.Sprintf("POL-%04d", i)
	}
	target := "POL-0999"

	b.ResetTimer()
	for range b.N {
		BuscarLineal(ids, target)
	}
}

func BenchmarkBuscarConMap(b *testing.B) {
	index := make(map[string]bool, 1000)
	for i := range 1000 {
		index[fmt.Sprintf("POL-%04d", i)] = true
	}
	target := "POL-0999"

	b.ResetTimer()
	for range b.N {
		BuscarConMap(index, target)
	}
}
