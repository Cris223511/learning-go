// Acá veremos los operadores de Go. La mayoría son iguales a Java o TypeScript,
// pero hay dos que engañan: la división entre enteros y el operador módulo.

package main

import "fmt"

func main() {
	fmt.Println("== Aritméticos ==")
	a, b := 10, 3

	fmt.Println("suma:", a+b)
	fmt.Println("resta:", a-b)
	fmt.Println("multiplicación:", a*b)

	// Si los dos operandos son enteros, la división descarta el decimal sin redondear.
	// Para obtener decimal, al menos uno de los dos tiene que ser float64.
	fmt.Println("división entera:", a/b)                    // 3, no 3.333
	fmt.Println("división decimal:", float64(a)/float64(b)) // 3.333...

	// El módulo devuelve el resto de la división. En Go, el signo del resultado
	// sigue al dividendo, no al divisor.
	fmt.Println("módulo:", a%b) // 1

	fmt.Println("== Comparación ==")
	fmt.Println("10 > 3:", a > b)
	fmt.Println("10 == 3:", a == b)
	fmt.Println("10 != 3:", a != b)

	fmt.Println("== Lógicos ==")
	esMayor := a > b
	esPositivo := a > 0
	// && es "y", || es "o", ! es negación. Igual que en Java o TypeScript.
	fmt.Println("mayor Y positivo:", esMayor && esPositivo)
	fmt.Println("mayor O negativo:", esMayor || !esPositivo)

	fmt.Println("== Asignación compuesta ==")
	contador := 0
	contador += 5
	contador -= 1
	contador *= 2
	fmt.Println("contador:", contador) // (0+5-1)*2 = 8

	// Go tiene ++ y -- pero solo como sentencias, no como expresiones.
	// No puedes hacer x := contador++ como en Java o TypeScript.
	contador++
	fmt.Println("contador++:", contador)

	fmt.Println("== Operadores sobre strings ==")
	saludo := "Hola" + ", " + "Go"
	fmt.Println(saludo)
	fmt.Println("¿iguales?:", "Go" == "Go")
}
