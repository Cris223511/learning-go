// Acá veremos los slices, que son la colección más usada en Go.
// Un slice no guarda datos por sí solo: es una ventana sobre un array subyacente.

package main

import "fmt"

func main() {
	// Declaración literal. Go crea el array por debajo automáticamente.
	productos := []string{"SOAT", "Vida", "Salud"}

	fmt.Println("== Slice básico ==")
	fmt.Println("productos:", productos)
	// len es cuántos elementos tiene ahora. cap es cuántos puede tener antes de que
	// Go tenga que crear un array más grande por debajo.
	fmt.Printf("len: %d | cap: %d\n", len(productos), cap(productos))

	// make([]Tipo, len, cap) crea un slice con longitud y capacidad definidas.
	// Todos los elementos quedan en su zero value (cadena vacía para string).
	reservas := make([]string, 3, 5)
	reservas[0] = "póliza-001"
	reservas[1] = "póliza-002"
	reservas[2] = "póliza-003"
	fmt.Println("\n== make ==")
	fmt.Println("reservas:", reservas)
	fmt.Printf("len: %d | cap: %d\n", len(reservas), cap(reservas))

	// append agrega elementos al final y devuelve el slice resultante.
	// Si la cap se agota, Go crea un array nuevo más grande por debajo y mueve todo.
	// Por eso siempre hay que reasignar: productos = append(productos, ...)
	fmt.Println("\n== append ==")
	productos = append(productos, "Vehicular")
	productos = append(productos, "Accidentes", "Hogar")
	fmt.Println("productos:", productos)
	fmt.Printf("len: %d | cap: %d\n", len(productos), cap(productos))

	// Slicing: puedes tomar una porción de un slice con [inicio:fin].
	// El índice de inicio es inclusivo, el de fin es exclusivo.
	// Ojo: la porción comparte el array subyacente con el original.
	fmt.Println("\n== Slicing ==")
	primeros := productos[0:3]
	ultimos := productos[3:]
	fmt.Println("primeros:", primeros)
	fmt.Println("ultimos:", ultimos)

	// copy crea una copia independiente. Sin copy, modificar la porción
	// puede modificar el slice original también.
	fmt.Println("\n== copy ==")
	destino := make([]string, len(productos))
	n := copy(destino, productos)
	destino[0] = "MODIFICADO"
	fmt.Printf("copiados: %d\n", n)
	fmt.Println("original intacto:", productos)
	fmt.Println("copia modificada:", destino)

	// Slice de dos dimensiones: slice de slices. Útil para matrices o tablas.
	fmt.Println("\n== Slice 2D ==")
	tabla := [][]int{
		{1, 2, 3},
		{4, 5, 6},
		{7, 8, 9},
	}
	for _, fila := range tabla {
		fmt.Println(fila)
	}
}
