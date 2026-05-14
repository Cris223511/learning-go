// Tipo base Persona. En Go no hay herencia: se usa composición (embedding)
// para reusar comportamiento entre tipos.

package main

import "fmt"

type Persona struct {
	Nombre   string
	Apellido string
	DNI      string
}

func (p Persona) NombreCompleto() string {
	return fmt.Sprintf("%s %s", p.Nombre, p.Apellido)
}

func (p Persona) Presentarse() string {
	return fmt.Sprintf("Hola, soy %s, DNI %s", p.NombreCompleto(), p.DNI)
}
