// Acá veremos los structs en Go: cómo definirlos, inicializarlos y organizarlos
// en múltiples archivos dentro del mismo paquete.

package main

import "fmt"

func main() {
	// Inicialización con nombre de campos. Es la forma recomendada: si alguien
	// agrega un campo al struct después, tu código no se rompe.
	p1 := Poliza{
		ID:       "POL-001",
		Producto: "SOAT",
		Prima:    120.00,
		Activa:   true,
	}

	// Sin nombres de campo: tienes que respetar el orden exacto de la definición.
	// Frágil, se evita en código real.
	p2 := Poliza{"POL-002", "Vida", 450.50, true}

	// Zero value de un struct: todos sus campos quedan en su zero value respectivo.
	var p3 Poliza
	p3.ID = "POL-003"
	p3.Producto = "Vehicular"
	p3.Prima = 890.75
	p3.Activa = false

	fmt.Println("== Structs ==")
	fmt.Printf("%+v\n", p1) // %+v muestra los nombres de los campos además de los valores
	fmt.Printf("%+v\n", p2)
	fmt.Printf("%+v\n", p3)

	// Usando el constructor definido en poliza.go
	fmt.Println("\n== Constructor ==")
	p4, err := NewPoliza("POL-004", "Salud", 320.00)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Printf("creada: %+v\n", p4)
	}

	_, err = NewPoliza("", "Hogar", 200)
	if err != nil {
		fmt.Println("error esperado:", err)
	}

	// Los structs se copian por valor igual que los arrays.
	// Modificar la copia no afecta al original.
	fmt.Println("\n== Copia por valor ==")
	copia := p1
	copia.Prima = 999
	fmt.Printf("original: %.2f | copia: %.2f\n", p1.Prima, copia.Prima)

	// Struct anónimo: útil para datos temporales o respuestas de API
	// donde no vale la pena definir un tipo con nombre.
	fmt.Println("\n== Struct anónimo ==")
	resumen := struct {
		Total    int
		Activas  int
	}{Total: 4, Activas: 3}
	fmt.Printf("pólizas: %+v\n", resumen)
}
