// Acá veremos los sentinel errors: errores predefinidos como variables globales
// que se usan para identificar condiciones específicas con errors.Is().

package main

import (
	"errors"
	"fmt"
)

// Sentinel errors: variables de error que representan condiciones conocidas.
// Por convención se llaman Err<Nombre> y son exportadas si otros paquetes las necesitan.
// Esto es lo mismo que io.EOF, sql.ErrNoRows, etc. en la stdlib.
var (
	ErrPolizaNoEncontrada = errors.New("póliza no encontrada")
	ErrPolizaVencida      = errors.New("póliza vencida")
	ErrSaldoInsuficiente  = errors.New("saldo insuficiente")
)

type Poliza struct {
	ID     string
	Activa bool
}

var polizas = map[string]Poliza{
	"POL-001": {ID: "POL-001", Activa: true},
	"POL-002": {ID: "POL-002", Activa: false},
}

func buscarPoliza(id string) (Poliza, error) {
	p, ok := polizas[id]
	if !ok {
		return Poliza{}, ErrPolizaNoEncontrada
	}
	if !p.Activa {
		return Poliza{}, ErrPolizaVencida
	}
	return p, nil
}

func procesarPago(polizaID string, monto, saldo float64) error {
	if _, err := buscarPoliza(polizaID); err != nil {
		return err
	}
	if monto > saldo {
		return ErrSaldoInsuficiente
	}
	return nil
}

func main() {
	fmt.Println("== Sentinel errors con errors.Is ==")

	ids := []string{"POL-001", "POL-002", "POL-999"}
	for _, id := range ids {
		_, err := buscarPoliza(id)
		if err == nil {
			fmt.Printf("%s: encontrada y activa\n", id)
			continue
		}
		// errors.Is compara el error (y todos sus envueltos) contra el sentinel.
		// Es la forma correcta, no comparar con == directamente.
		switch {
		case errors.Is(err, ErrPolizaNoEncontrada):
			fmt.Printf("%s: no existe en el sistema\n", id)
		case errors.Is(err, ErrPolizaVencida):
			fmt.Printf("%s: existe pero está vencida\n", id)
		}
	}

	fmt.Println("\n== Encadenando decisiones con sentinel errors ==")
	casos := []struct {
		poliza string
		monto  float64
		saldo  float64
	}{
		{"POL-001", 100, 500},
		{"POL-001", 600, 500},
		{"POL-002", 100, 500},
		{"POL-999", 100, 500},
	}
	for _, c := range casos {
		err := procesarPago(c.poliza, c.monto, c.saldo)
		if err == nil {
			fmt.Printf("  %s: pago de S/%.0f aprobado\n", c.poliza, c.monto)
			continue
		}
		switch {
		case errors.Is(err, ErrPolizaNoEncontrada):
			fmt.Printf("  %s: póliza no existe\n", c.poliza)
		case errors.Is(err, ErrPolizaVencida):
			fmt.Printf("  %s: póliza vencida, no se puede cobrar\n", c.poliza)
		case errors.Is(err, ErrSaldoInsuficiente):
			fmt.Printf("  %s: saldo insuficiente para S/%.0f\n", c.poliza, c.monto)
		}
	}
}
