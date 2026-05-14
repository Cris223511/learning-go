// Acá veremos los tipos de datos básicos de Go. Hay dos detalles que sorprenden
// a developers de otros lenguajes: la conversión siempre es explícita y existe el tipo rune.

package main

import "fmt"

func main() {
	// Tipos numéricos más usados en el día a día.
	// int toma 32 o 64 bits según la arquitectura del sistema. float64 es el decimal estándar.
	var edad int = 24
	var salario float64 = 4500.75
	var codigo int32 = 10482

	// string en Go es una secuencia de bytes UTF-8. Se declara con comillas dobles.
	var nombre string = "Christopher"

	// bool solo acepta true o false, sin variantes como 0/1 o "yes"/"no".
	var activo bool = true

	// byte es un alias de uint8 (número del 0 al 255). Se usa para datos binarios o ASCII.
	// rune es un alias de int32 y representa un carácter Unicode completo, incluyendo tildes y emojis.
	// Fíjate que Go distingue entre el byte (posición en memoria) y el rune (carácter real).
	var inicial byte = 'C'
	var letra rune = 'ñ'

	// any es el alias moderno de interface{}. Acepta cualquier tipo de valor.
	// Úsalo con cuidado: pierdes la verificación de tipos en tiempo de compilación.
	var dato any = "puede ser cualquier cosa"

	fmt.Println("== Tipos numéricos ==")
	fmt.Printf("edad: %d | salario: %.2f | codigo: %d\n", edad, salario, codigo)

	fmt.Println("== String y bool ==")
	fmt.Printf("nombre: %s | activo: %t\n", nombre, activo)

	fmt.Println("== byte y rune ==")
	// %c imprime el carácter, %d el valor numérico detrás.
	fmt.Printf("inicial: %c (%d) | letra: %c (%d)\n", inicial, inicial, letra, letra)

	fmt.Println("== any ==")
	fmt.Println(dato)

	// Conversión explícita: Go no convierte tipos automáticamente, ni siquiera entre int y float64.
	// Tienes que envolverlo con el tipo destino como si fuera una función.
	fmt.Println("== Conversión explícita ==")
	var entero int = 7
	var decimal float64 = float64(entero) * 1.5
	fmt.Printf("entero: %d | convertido y multiplicado: %.2f\n", entero, decimal)
}
