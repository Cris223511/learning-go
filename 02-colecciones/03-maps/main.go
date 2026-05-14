// Acá veremos los maps en Go: la estructura clave-valor del lenguaje.
// Hay dos detalles clave: cómo leer de forma segura y que el orden no está garantizado.

package main

import "fmt"

func main() {
	// Declaración literal. La clave y el valor pueden ser casi cualquier tipo.
	precios := map[string]float64{
		"SOAT":      120.00,
		"Vida":      450.50,
		"Vehicular": 890.75,
	}

	// make crea un map vacío listo para usar.
	// Un map declarado con var queda nil y paniquea al escribir en él.
	polizas := make(map[string]string)
	polizas["POL-001"] = "activa"
	polizas["POL-002"] = "vencida"
	polizas["POL-003"] = "activa"

	fmt.Println("== Maps ==")
	fmt.Println("precios:", precios)
	fmt.Println("polizas:", polizas)

	// Lectura simple: si la clave no existe, devuelve el zero value del tipo (0 para float64).
	// Para saber si realmente existe, usa la forma de dos valores: valor, ok.
	fmt.Println("\n== Lectura segura ==")
	precio, ok := precios["SOAT"]
	fmt.Printf("SOAT: %.2f | existe: %t\n", precio, ok)

	precio2, ok2 := precios["Dental"]
	fmt.Printf("Dental: %.2f | existe: %t\n", precio2, ok2)

	// Patrón común en Go: leer y verificar en el mismo if.
	if estado, existe := polizas["POL-002"]; existe {
		fmt.Println("estado de POL-002:", estado)
	} else {
		fmt.Println("póliza no encontrada")
	}

	// delete elimina una clave. Si la clave no existe, no hace nada ni da error.
	fmt.Println("\n== delete ==")
	delete(polizas, "POL-002")
	fmt.Println("después de delete:", polizas)

	// for-range sobre un map entrega clave y valor. El orden de iteración no está
	// garantizado. Cada vez que corras el programa pueden salir en distinto orden.
	fmt.Println("\n== Iteración (orden no garantizado) ==")
	for producto, precio := range precios {
		fmt.Printf("%s: S/ %.2f\n", producto, precio)
	}

	// len sobre un map devuelve la cantidad de claves que tiene.
	fmt.Println("\ntotal de productos:", len(precios))

	// Map con slice como valor: útil para agrupar elementos por categoría.
	fmt.Println("\n== Map de slices ==")
	categorias := map[string][]string{
		"vida":    {"Vida Entera", "Vida Temporal"},
		"general": {"SOAT", "Vehicular", "Hogar"},
	}
	for cat, items := range categorias {
		fmt.Printf("%s: %v\n", cat, items)
	}
}
