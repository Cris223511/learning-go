// Acá veremos las distintas formas de declarar variables en Go y un concepto
// clave que no existe igual en otros lenguajes: los zero values.

package main

import "fmt"

func main() {
	// Forma larga: declaras el tipo explícitamente. Útil cuando quieres dejar claro el tipo
	// aunque el valor ya lo diga.
	var nombre string = "Christopher"
	var edad int = 24
	var activo bool = true

	// Forma corta con :=. Go adivina el tipo mirando el valor asignado.
	// Solo funciona dentro de funciones, nunca a nivel de paquete.
	pais := "Perú"
	salario := 4500.50

	// Cuando tienes varias variables relacionadas, puedes agruparlas en un bloque var()
	// para no repetir la palabra var en cada línea.
	var (
		empresa   = "Go"
		cargo     = "Desarrollador Go"
		ingreso   = "2026-06-01"
		modalidad = "Híbrida"
	)

	// Zero values: si declaras una variable sin asignarle nada, Go no la deja indefinida.
	// Le pone un valor seguro por defecto. Para string es "", para int es 0, para bool es false.
	var pendiente string
	var contador int
	var listo bool

	fmt.Println("== Declaración larga ==")
	fmt.Println(nombre, edad, activo)

	fmt.Println("== Inferencia con := ==")
	fmt.Println(pais, salario)

	fmt.Println("== Bloque var() ==")
	fmt.Println(empresa, cargo, ingreso, modalidad)

	// %q imprime el string entre comillas, útil para ver si realmente está vacío.
	// %d es para enteros, %t para booleanos.
	fmt.Println("== Zero values ==")
	fmt.Printf("string: %q | int: %d | bool: %t\n", pendiente, contador, listo)
}
