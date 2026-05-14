// Acá veremos composición en Go: cómo un struct puede embeber otro para reusar
// sus campos y métodos, sin herencia clásica.

package main

import "fmt"

func main() {
	fmt.Println("== Embedding básico ==")
	emp := Empleado{
		Persona: Persona{
			Nombre:   "Christopher",
			Apellido: "Pillihuamán",
			DNI:      "12345678",
		},
		Cargo:   "Desarrollador Go",
		Salario: 4800,
	}

	// Los campos y métodos de Persona se acceden directamente en Empleado
	// gracias al embedding. Go los "promueve" al nivel superior.
	fmt.Println("nombre:", emp.Nombre)           // promovido desde Persona
	fmt.Println("DNI:", emp.DNI)                 // promovido desde Persona
	fmt.Println("completo:", emp.NombreCompleto()) // método promovido

	fmt.Println("\n== Método sobreescrito ==")
	// Empleado tiene su propio Presentarse() que reemplaza al de Persona.
	fmt.Println(emp.Presentarse())
	// Aun así puedes llamar al original con el tipo explícito.
	fmt.Println("original de Persona:", emp.Persona.Presentarse())

	fmt.Println("\n== Otro tipo que embebe Persona ==")
	aseg := Asegurado{
		Persona:    Persona{Nombre: "Ana", Apellido: "García", DNI: "87654321"},
		NumPolizas: 3,
		Vigente:    true,
	}
	fmt.Println(aseg.Resumen())
	fmt.Println(aseg.Presentarse()) // usa el Presentarse() de Persona, no hay uno propio

	fmt.Println("\n== Embedding múltiple ==")
	agente := AgenteBroker{
		Persona:      Persona{Nombre: "Luis", Apellido: "Torres", DNI: "11223344"},
		Empleado:     Empleado{Cargo: "Broker Senior", Salario: 6500},
		CodigoAgente: "AG-099",
	}
	fmt.Println(agente.Ficha())
	// Cuando hay ambigüedad (Persona y Empleado ambos tienen NombreCompleto
	// a través del embedding), debes ser explícito.
	fmt.Println("nombre vía Persona:", agente.Persona.NombreCompleto())
}
