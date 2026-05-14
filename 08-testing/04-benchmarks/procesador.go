// Funciones a comparar en benchmarks. Los benchmarks miden cuánto tarda
// cada implementación y permiten elegir la más eficiente.

package main

import (
	"fmt"
	"strings"
)

// ConcatenarConPlus construye un string con el operador +.
// Ineficiente: crea un string nuevo en cada iteración.
func ConcatenarConPlus(n int) string {
	resultado := ""
	for i := range n {
		resultado += fmt.Sprintf("póliza-%d,", i)
	}
	return resultado
}

// ConcatenarConBuilder usa strings.Builder, que escribe en un buffer interno
// sin crear strings intermedios. Mucho más eficiente para muchas concatenaciones.
func ConcatenarConBuilder(n int) string {
	var sb strings.Builder
	for i := range n {
		fmt.Fprintf(&sb, "póliza-%d,", i)
	}
	return sb.String()
}

// BuscarLineal recorre todo el slice hasta encontrar el elemento.
func BuscarLineal(ids []string, target string) bool {
	for _, id := range ids {
		if id == target {
			return true
		}
	}
	return false
}

// BuscarConMap usa un map para O(1) en lugar de O(n).
func BuscarConMap(index map[string]bool, target string) bool {
	return index[target]
}
