// Acá veremos punteros en Go: cuándo usarlos con structs y cuál es la diferencia
// real entre pasar un struct por valor versus por puntero.

package main

import "fmt"

type Cliente struct {
	Nombre string
	Email  string
	Saldo  float64
}

// Recibe una copia del struct. El original no cambia.
func aplicarBonificacionValor(c Cliente, monto float64) {
	c.Saldo += monto
	fmt.Printf("  dentro de la función: %.2f\n", c.Saldo)
}

// Recibe un puntero al struct original. Los cambios persisten afuera.
// En Go, el * en el tipo indica que es un puntero. El . para acceder a campos
// funciona igual, Go desreferencia automáticamente (no necesitas ->).
func aplicarBonificacionPuntero(c *Cliente, monto float64) {
	c.Saldo += monto
	fmt.Printf("  dentro de la función: %.2f\n", c.Saldo)
}

func main() {
	fmt.Println("== Por valor: el original no cambia ==")
	c1 := Cliente{Nombre: "Christopher", Email: "c@mail.com", Saldo: 1000}
	fmt.Printf("antes: %.2f\n", c1.Saldo)
	aplicarBonificacionValor(c1, 500)
	fmt.Printf("después: %.2f\n", c1.Saldo) // sigue en 1000

	fmt.Println("\n== Por puntero: el original cambia ==")
	c2 := Cliente{Nombre: "Ana", Email: "ana@mail.com", Saldo: 1000}
	fmt.Printf("antes: %.2f\n", c2.Saldo)
	aplicarBonificacionPuntero(&c2, 500)
	fmt.Printf("después: %.2f\n", c2.Saldo) // ahora es 1500

	// new() crea un puntero a un struct con todos sus campos en zero value.
	// Es lo mismo que &Cliente{}, solo una forma alternativa.
	fmt.Println("\n== Puntero con new() ==")
	c3 := new(Cliente)
	c3.Nombre = "Luis"
	c3.Saldo = 2000
	fmt.Printf("%+v\n", *c3) // * desreferencia para ver el valor

	// Cuando el struct es grande o se modifica frecuentemente,
	// pasar puntero es más eficiente porque no copia toda la memoria.
	fmt.Println("\n== Puntero directo con & ==")
	c4 := &Cliente{Nombre: "María", Email: "maria@mail.com", Saldo: 500}
	c4.Saldo += 200 // Go desreferencia automáticamente, no necesitas (*c4).Saldo
	fmt.Printf("%+v\n", *c4)
}
