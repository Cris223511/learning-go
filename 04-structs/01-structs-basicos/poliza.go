// Definición del struct Poliza y su constructor. En Go es común separar
// los tipos en su propio archivo dentro del mismo paquete.

package main

import "fmt"

// Un struct agrupa campos relacionados bajo un mismo tipo. Es el equivalente
// a una clase sin herencia. Los nombres en PascalCase son exportados (públicos).
type Poliza struct {
	ID       string
	Producto string
	Prima    float64
	Activa   bool
}

// En Go no hay constructores del lenguaje. Por convención se usa una función
// llamada New<Tipo> que valida y retorna el struct listo para usar.
func NewPoliza(id, producto string, prima float64) (Poliza, error) {
	if id == "" || producto == "" {
		return Poliza{}, fmt.Errorf("id y producto son obligatorios")
	}
	if prima <= 0 {
		return Poliza{}, fmt.Errorf("la prima debe ser mayor a cero")
	}
	return Poliza{
		ID:       id,
		Producto: producto,
		Prima:    prima,
		Activa:   true,
	}, nil
}
