// Acá veremos los bucles en Go. Lo primero que sorprende: solo existe for.
// No hay while ni do-while, pero el for tiene varias formas que cubren todos esos casos.

package main

import "fmt"

func main() {
	// Forma clásica con tres partes: inicialización, condición y paso.
	fmt.Println("== for clásico ==")
	for i := 1; i <= 5; i++ {
		fmt.Printf("iteración %d\n", i)
	}

	// Si solo pones la condición, se comporta exactamente como un while.
	fmt.Println("== for estilo while ==")
	contador := 0
	for contador < 3 {
		fmt.Println("contador:", contador)
		contador++
	}

	// Sin condición el bucle es infinito. Se sale con break.
	// Útil para reintentos o servidores que escuchan indefinidamente.
	fmt.Println("== for infinito con break ==")
	intento := 0
	for {
		intento++
		if intento == 3 {
			fmt.Println("éxito en intento", intento)
			break
		}
		fmt.Println("reintentando...")
	}

	// for-range sobre un slice: entrega el índice y el valor en cada iteración.
	// Si no necesitas el índice, usa _ para descartarlo.
	fmt.Println("== for-range sobre slice ==")
	productos := []string{"SOAT", "Vida", "Salud", "Vehicular"}
	for i, p := range productos {
		fmt.Printf("[%d] %s\n", i, p)
	}

	// for-range sobre un map: entrega clave y valor. Importante: el orden no está
	// garantizado. Cada ejecución puede mostrar las claves en distinto orden.
	fmt.Println("== for-range sobre map ==")
	precios := map[string]float64{
		"SOAT":      120.00,
		"Vida":      450.50,
		"Vehicular": 890.75,
	}
	for nombre, precio := range precios {
		fmt.Printf("%s: S/ %.2f\n", nombre, precio)
	}

	// for-range sobre un string entrega la posición en bytes y cada carácter como rune.
	// Así puedes recorrer texto con tildes o caracteres especiales sin romper nada.
	fmt.Println("== for-range sobre string ==")
	for i, r := range "Perú" {
		fmt.Printf("posición %d: %c\n", i, r)
	}

	// continue salta al siguiente ciclo sin ejecutar lo que queda abajo.
	fmt.Println("== continue ==")
	for i := 1; i <= 6; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Println("impar:", i)
	}

	// break con label: cuando tienes un for dentro de otro, un break normal solo sale
	// del bucle interno. Con un label puedes romper el externo directamente desde adentro.
	fmt.Println("== break con label ==")
outer:
	for i := 1; i <= 3; i++ {
		for j := 1; j <= 3; j++ {
			if i*j > 4 {
				fmt.Printf("rompo en i=%d, j=%d\n", i, j)
				break outer
			}
			fmt.Printf("  %d x %d = %d\n", i, j, i*j)
		}
	}
}
