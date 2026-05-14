// Acá veremos las interfaces en Go: contratos implícitos que permiten trabajar
// con distintos tipos de forma uniforme sin herencia.

package main

import "fmt"

// Esta función acepta cualquier tipo que cumpla la interface Seguro.
// No sabe ni le importa si es SOAT, Vida o Vehicular.
func imprimirSeguro(s Seguro) {
	fmt.Printf("  %s | prima: S/%.2f\n", s.Descripcion(), s.PrimaAnual())
}

func main() {
	soat := SOAT{Placa: "ABC-123", Propietario: "Christopher"}
	vida := Vida{Asegurado: "Ana García", Capital: 100_000}
	veh := Vehicular{Placa: "XYZ-789", Marca: "Toyota", ValorVehiculo: 45_000}

	fmt.Println("== Usando la interface ==")
	imprimirSeguro(soat)
	imprimirSeguro(vida)
	imprimirSeguro(veh)

	// Puedes guardar distintos tipos concretos en un slice de interface.
	// Todos caben porque todos cumplen Seguro.
	fmt.Println("\n== Slice de interfaces ==")
	cartera := []Seguro{soat, vida, veh}
	total := 0.0
	for _, s := range cartera {
		total += s.PrimaAnual()
		fmt.Printf("  %s\n", s.Descripcion())
	}
	fmt.Printf("total de primas: S/%.2f\n", total)

	// La interface nil: una variable de tipo interface sin valor asignado es nil.
	// Llamar un método sobre nil paniquea, hay que verificar antes.
	fmt.Println("\n== Interface nil ==")
	var s Seguro
	fmt.Println("es nil:", s == nil)

	// fmt.Stringer es la interface más común de la stdlib.
	// Si tu tipo implementa String() string, fmt la usa automáticamente.
	fmt.Println("\n== fmt.Stringer ==")
	fmt.Println(soat.Descripcion()) // manual
	// En el módulo de polimorfismo veremos cómo integrar Stringer directamente con fmt.
}
