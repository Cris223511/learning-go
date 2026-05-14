// Acá veremos la estructura mínima de un programa en Go: qué es un paquete,
// cómo se importa una librería estándar y las dos formas básicas de imprimir en consola.

// Todo programa ejecutable en Go necesita declarar el paquete main. Si tuviera
// cualquier otro nombre, Go lo trataría como una librería, no como un programa que se puede correr.
package main

// fmt viene incluido en la librería estándar de Go, no se instala. Sirve para imprimir
// y dar formato a texto en consola.
import "fmt"

func main() {
	// Println imprime el texto y baja de línea automáticamente al final.
	fmt.Println("Hola, Go.")

	// Printf no baja de línea solo, por eso le ponemos \n al final.
	// Los % son marcadores de posición: %s espera un string, %d un entero, %f un decimal.
	// Los valores que los llenan van después de la cadena, en el mismo orden.
	fmt.Printf("Bienvenido %s, aprendiendo Go.\n", "Christopher")
}
