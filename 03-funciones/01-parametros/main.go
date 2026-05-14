// Acá veremos cómo se declaran funciones en Go: parámetros por valor, por puntero,
// variádicos y funciones como valores de primera clase.

package main

import "fmt"

// Go pasa los parámetros por valor: la función recibe una copia, no el original.
// Modificar el parámetro adentro no afecta a quien llamó la función.
func aplicarDescuento(precio float64, porcentaje float64) float64 {
	return precio - (precio * porcentaje / 100)
}

// Si quieres modificar el valor original, pasas un puntero con *.
// El & en la llamada toma la dirección de memoria de la variable.
func activarPoliza(activa *bool) {
	*activa = true
}

// Cuando varios parámetros seguidos son del mismo tipo, puedes agrupar el tipo al final.
func registrar(nombre, cargo string, salario float64) {
	fmt.Printf("%s (%s) - S/ %.2f\n", nombre, cargo, salario)
}

// Los ... declaran un parámetro variádico: acepta cero o más valores del mismo tipo.
// Dentro de la función se comporta como un slice.
func sumar(numeros ...float64) float64 {
	total := 0.0
	for _, n := range numeros {
		total += n
	}
	return total
}

// Las funciones son valores de primera clase en Go. Puedes guardarlas en variables,
// pasarlas como parámetros o retornarlas. Acá recibimos una función como argumento.
func aplicar(precio float64, fn func(float64) float64) float64 {
	return fn(precio)
}

func main() {
	fmt.Println("== Por valor ==")
	precio := 450.50
	descuento := aplicarDescuento(precio, 10)
	fmt.Printf("original: %.2f | con descuento: %.2f\n", precio, descuento)

	fmt.Println("\n== Por puntero ==")
	activa := false
	fmt.Println("antes:", activa)
	activarPoliza(&activa)
	fmt.Println("después:", activa)

	fmt.Println("\n== Tipos agrupados ==")
	registrar("Christopher", "Desarrollador Go", 4800)

	fmt.Println("\n== Variádico ==")
	fmt.Println("suma:", sumar(120, 450.50, 890.75))
	fmt.Println("suma vacía:", sumar())

	// Si ya tienes un slice, lo expandes con ... para pasarlo a una función variádica.
	primas := []float64{100, 200, 300}
	fmt.Println("desde slice:", sumar(primas...))

	fmt.Println("\n== Función como argumento ==")
	igv := func(p float64) float64 { return p * 1.18 }
	fmt.Printf("precio con IGV: %.2f\n", aplicar(450.50, igv))
}
