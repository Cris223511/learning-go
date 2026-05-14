// Tipos de error propios del dominio. En Go, error es una interface con un solo
// método: Error() string. Cualquier tipo que lo implemente es un error válido.

package main

import "fmt"

// ErrorValidacion es un tipo de error con campos propios. Así puedes transportar
// información estructurada junto con el mensaje de error.
type ErrorValidacion struct {
	Campo   string
	Mensaje string
}

func (e ErrorValidacion) Error() string {
	return fmt.Sprintf("validación fallida en %q: %s", e.Campo, e.Mensaje)
}

// ErrorNegocio representa una regla de negocio violada, distinta de un error técnico.
type ErrorNegocio struct {
	Codigo  string
	Detalle string
}

func (e ErrorNegocio) Error() string {
	return fmt.Sprintf("[%s] %s", e.Codigo, e.Detalle)
}
