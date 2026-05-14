// Acá veremos los condicionales de Go: if, else if, else y switch.
// Lo más distinto respecto a otros lenguajes es el if con inicialización y el switch sin break.

package main

import "fmt"

func main() {
	fmt.Println("== if / else if / else ==")
	edad := 24

	if edad >= 18 {
		fmt.Println("mayor de edad")
	} else if edad >= 13 {
		fmt.Println("adolescente")
	} else {
		fmt.Println("menor de edad")
	}

	// If con inicialización: puedes declarar una variable justo antes de la condición,
	// separada por punto y coma. Esa variable solo existe dentro del bloque if/else,
	// no contamina el scope de afuera. Muy común al evaluar resultados de funciones.
	fmt.Println("== if con inicialización ==")
	if puntaje := 87; puntaje >= 90 {
		fmt.Println("excelente")
	} else if puntaje >= 70 {
		fmt.Println("aprobado con puntaje:", puntaje)
	} else {
		fmt.Println("reprobado")
	}

	// En Go el switch no necesita break. Cada case termina solo.
	// Si quieres que continúe al siguiente case, usas fallthrough explícitamente.
	fmt.Println("== switch ==")
	dia := "lunes"
	switch dia {
	case "lunes", "martes", "miércoles", "jueves", "viernes":
		fmt.Println("día hábil")
	case "sábado", "domingo":
		fmt.Println("fin de semana")
	default:
		fmt.Println("día desconocido")
	}

	// Switch sin expresión: cada case lleva su propia condición booleana.
	// Es una forma más limpia de escribir cadenas largas de if/else if.
	fmt.Println("== switch sin expresión ==")
	salario := 4800.0
	switch {
	case salario > 10000:
		fmt.Println("rango alto")
	case salario > 4000:
		fmt.Println("rango medio, salario:", salario)
	default:
		fmt.Println("rango bajo")
	}
}
