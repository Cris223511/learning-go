// Acá veremos los receivers: la forma en que Go le agrega comportamiento a los structs.
// Es lo más cercano a los métodos de una clase, sin herencia.

package main

import "fmt"

func main() {
	// Cliente está definido en cliente.go, mismo paquete, todo visible.
	c := &Cliente{
		Nombre: "Christopher",
		Email:  "c@correo.com",
		Saldo:  1000.00,
		Activo: true,
	}

	fmt.Println("== Receiver por valor (solo lectura) ==")
	fmt.Println(c.Descripcion())

	fmt.Println("\n== Receiver por puntero (modifica el struct) ==")
	if err := c.Depositar(500); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("después de depositar:", c.Descripcion())

	if err := c.Retirar(200); err != nil {
		fmt.Println("error:", err)
	}
	fmt.Println("después de retirar:", c.Descripcion())

	// Error controlado: intento de retirar más de lo disponible.
	if err := c.Retirar(9999); err != nil {
		fmt.Println("error esperado:", err)
	}

	fmt.Println("\n== Desactivar ==")
	c.Desactivar()
	fmt.Println(c.Descripcion())
}
