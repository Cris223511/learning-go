// Tipo Empleado que embebe Persona. El embedding promueve todos los campos
// y métodos de Persona directamente en Empleado, sin herencia.

package main

import "fmt"

type Empleado struct {
	Persona          // embedding: no tiene nombre de campo, solo el tipo
	Cargo   string
	Salario float64
}

// Empleado puede sobreescribir métodos de Persona simplemente definiéndolos.
// Go llama al método más específico (el de Empleado), no al de Persona.
func (e Empleado) Presentarse() string {
	return fmt.Sprintf("%s | cargo: %s | S/%.2f",
		e.Persona.Presentarse(), // así accedes explícitamente al método original
		e.Cargo,
		e.Salario,
	)
}

type Asegurado struct {
	Persona
	NumPolizas int
	Vigente    bool
}

func (a Asegurado) Resumen() string {
	activo := "sin cobertura activa"
	if a.Vigente {
		activo = "con cobertura vigente"
	}
	return fmt.Sprintf("%s | %d póliza(s) | %s", a.NombreCompleto(), a.NumPolizas, activo)
}

// Un struct puede embeber múltiples tipos a la vez.
type AgenteBroker struct {
	Persona
	Empleado
	CodigoAgente string
}

func (ab AgenteBroker) Ficha() string {
	return fmt.Sprintf("agente %s | código: %s | %s",
		ab.Persona.NombreCompleto(),
		ab.CodigoAgente,
		ab.Empleado.Cargo,
	)
}

