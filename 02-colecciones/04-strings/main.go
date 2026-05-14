// Acá veremos cómo trabaja Go con strings. El punto más importante:
// len() cuenta bytes, no caracteres, y los strings son inmutables.

package main

import (
	"fmt"
	"strings"
	"unicode/utf8"
)

func main() {
	nombre := "Ángel García"

	fmt.Println("== Longitud ==")
	// len cuenta bytes, no letras. "Á" ocupa 2 bytes en UTF-8, así que el resultado
	// puede ser mayor al número de caracteres que ves.
	fmt.Println("bytes:", len(nombre))
	// RuneCountInString cuenta los caracteres reales (runes), incluyendo tildes.
	fmt.Println("caracteres:", utf8.RuneCountInString(nombre))

	// Recorrer con range entrega runes, no bytes. Es la forma correcta de iterar letras.
	fmt.Println("\n== Recorrido con range ==")
	for i, r := range nombre {
		fmt.Printf("posición %d: %c\n", i, r)
	}

	// Convertir a []rune te da acceso por índice a cada carácter real.
	// Convertir a []byte te da acceso a los bytes crudos, útil para operaciones binarias.
	fmt.Println("\n== Conversiones ==")
	runes := []rune(nombre)
	fmt.Println("primer carácter:", string(runes[0]))
	fmt.Println("últimos 6 caracteres:", string(runes[len(runes)-6:]))

	// El paquete strings tiene todo lo que necesitas para manipular texto.
	fmt.Println("\n== Paquete strings ==")
	s := "  Go Learning  "
	fmt.Println("TrimSpace:", strings.TrimSpace(s))
	fmt.Println("ToUpper:", strings.ToUpper("go"))
	fmt.Println("ToLower:", strings.ToLower("POLIZA"))
	fmt.Println("Contains:", strings.Contains(s, "Learning"))
	fmt.Println("HasPrefix:", strings.HasPrefix(strings.TrimSpace(s), "Go"))
	fmt.Println("Replace:", strings.Replace("póliza-001-001", "001", "002", 1)) // reemplaza solo la primera
	fmt.Println("ReplaceAll:", strings.ReplaceAll("a-a-a", "a", "b"))
	fmt.Println("Split:", strings.Split("SOAT,Vida,Salud", ","))
	fmt.Println("Join:", strings.Join([]string{"Go", "es", "rápido"}, " "))
	fmt.Println("Count:", strings.Count("banana", "a"))
	fmt.Println("Index:", strings.Index("Go", "seguro")) // -1 si no encuentra

	// Los strings son inmutables. Para construir texto en un bucle
	// es más eficiente usar strings.Builder que concatenar con +.
	fmt.Println("\n== strings.Builder ==")
	var sb strings.Builder
	for i := 1; i <= 4; i++ {
		fmt.Fprintf(&sb, "ítem %d\n", i)
	}
	fmt.Print(sb.String())
}
