// Definición del tipo Cliente con sus métodos. Los métodos van en el mismo
// paquete que el tipo, pueden estar en archivos separados sin problema.

package main

import "fmt"

type Cliente struct {
	Nombre  string
	Email   string
	Saldo   float64
	Activo  bool
}

// Receiver por valor: trabaja con una copia del struct. Úsalo cuando el método
// solo lee datos y no necesita modificar el original.
func (c Cliente) Descripcion() string {
	estado := "inactivo"
	if c.Activo {
		estado = "activo"
	}
	return fmt.Sprintf("%s <%s> | saldo: S/%.2f | %s", c.Nombre, c.Email, c.Saldo, estado)
}

// Receiver por puntero: trabaja con el original. Úsalo cuando el método
// modifica el struct. Por convención, si un método necesita puntero,
// todos los métodos del tipo deben usar puntero para ser consistentes.
func (c *Cliente) Depositar(monto float64) error {
	if monto <= 0 {
		return fmt.Errorf("el monto debe ser mayor a cero")
	}
	c.Saldo += monto
	return nil
}

func (c *Cliente) Retirar(monto float64) error {
	if monto <= 0 {
		return fmt.Errorf("el monto debe ser mayor a cero")
	}
	if monto > c.Saldo {
		return fmt.Errorf("saldo insuficiente: disponible S/%.2f", c.Saldo)
	}
	c.Saldo -= monto
	return nil
}

func (c *Cliente) Desactivar() {
	c.Activo = false
}
