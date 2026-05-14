// Acá veremos el type switch: la forma idiomática de Go para manejar
// una variable que puede ser de varios tipos distintos.

package main

import "fmt"

type Poliza struct {
	ID    string
	Prima float64
}

type Siniestro struct {
	ID     string
	Monto  float64
	Motivo string
}

type Reclamo struct {
	ID      string
	Estado  string
}

// El type switch es como un switch normal pero en lugar de comparar valores,
// compara el tipo concreto de una variable de interface.
// La variable v dentro de cada case ya tiene el tipo correcto, sin necesidad de assertion.
func describir(evento any) string {
	switch v := evento.(type) {
	case Poliza:
		return fmt.Sprintf("Póliza %s | prima S/%.2f", v.ID, v.Prima)
	case Siniestro:
		return fmt.Sprintf("Siniestro %s | monto S/%.2f | motivo: %s", v.ID, v.Monto, v.Motivo)
	case Reclamo:
		return fmt.Sprintf("Reclamo %s | estado: %s", v.ID, v.Estado)
	case string:
		return fmt.Sprintf("texto: %q", v)
	case int, float64:
		return fmt.Sprintf("número: %v", v)
	case nil:
		return "valor nil recibido"
	default:
		// %T imprime el nombre del tipo de la variable.
		return fmt.Sprintf("tipo desconocido: %T", v)
	}
}

func main() {
	fmt.Println("== Type switch ==")
	eventos := []any{
		Poliza{ID: "POL-001", Prima: 450.50},
		Siniestro{ID: "SIN-007", Monto: 3200, Motivo: "choque frontal"},
		Reclamo{ID: "REC-003", Estado: "en revisión"},
		"mensaje suelto",
		42,
		nil,
		true, // tipo bool no está en ningún case, va al default
	}

	for _, e := range eventos {
		fmt.Println(" ", describir(e))
	}

	// Type switch con múltiples tipos en un case: cuando quieres el mismo
	// comportamiento para varios tipos pero no necesitas acceder a los campos.
	fmt.Println("\n== Case con múltiples tipos ==")
	for _, e := range eventos {
		switch e.(type) {
		case Poliza, Siniestro, Reclamo:
			fmt.Printf("  evento de negocio: %T\n", e)
		case nil:
			fmt.Println("  nada")
		default:
			fmt.Printf("  dato primitivo: %v\n", e)
		}
	}
}
