// Acá veremos cómo declarar constantes en Go y una herramienta muy particular
// del lenguaje llamada iota, que no existe en Java ni en TypeScript.

package main

import "fmt"

// Las constantes se declaran con const. Su valor se fija en tiempo de compilación
// y no puede cambiar en ningún momento de la ejecución.
const pi = 3.14159
const empresa = "Go"

// iota es un contador automático que solo existe dentro de bloques const.
// Empieza en 0 y sube de uno en uno con cada constante del bloque.
// Se usa mucho para representar estados, roles o categorías sin escribir los números a mano.
const (
	Pendiente = iota // 0
	EnProceso        // 1
	Aprobado         // 2
	Rechazado        // 3
)

// Puedes usar iota dentro de una expresión para que los valores no sean 0,1,2,3.
// Acá multiplicamos para representar porcentajes de descuento.
const (
	DescuentoBasico    = (iota + 1) * 10 // 10
	DescuentoIntermedio                  // 20
	DescuentoPremium                     // 30
)

func main() {
	fmt.Println("== Constantes simples ==")
	fmt.Println("Pi:", pi)
	fmt.Println("Empresa:", empresa)

	fmt.Println("== Estados con iota ==")
	fmt.Println("Pendiente:", Pendiente)
	fmt.Println("EnProceso:", EnProceso)
	fmt.Println("Aprobado:", Aprobado)
	fmt.Println("Rechazado:", Rechazado)

	fmt.Println("== Descuentos con iota expresión ==")
	fmt.Println("Básico:", DescuentoBasico, "%")
	fmt.Println("Intermedio:", DescuentoIntermedio, "%")
	fmt.Println("Premium:", DescuentoPremium, "%")
}
