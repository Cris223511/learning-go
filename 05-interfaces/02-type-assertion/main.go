// Acá veremos type assertion: cómo recuperar el tipo concreto que hay
// detrás de una variable de interface.

package main

import "fmt"

type Pagable interface {
	Cobrar() float64
}

type PagoTarjeta struct {
	NumeroTarjeta string
	Monto         float64
}

func (p PagoTarjeta) Cobrar() float64 { return p.Monto }

type PagoEfectivo struct {
	Monto float64
}

func (p PagoEfectivo) Cobrar() float64 { return p.Monto }

func procesarPago(p Pagable) {
	fmt.Printf("cobrado: S/%.2f\n", p.Cobrar())

	// Type assertion: le dices a Go "esto es en realidad un PagoTarjeta".
	// Si el tipo concreto no coincide, sin el ok paniquea en runtime.
	// Con ok, simplemente devuelve false y puedes manejarlo con seguridad.
	tarjeta, ok := p.(PagoTarjeta)
	if ok {
		// Ahora tarjeta es PagoTarjeta y puedes acceder a sus campos específicos.
		fmt.Printf("  pago con tarjeta terminada en %s\n",
			tarjeta.NumeroTarjeta[len(tarjeta.NumeroTarjeta)-4:])
	} else {
		fmt.Println("  pago en efectivo, sin datos de tarjeta")
	}
}

func main() {
	fmt.Println("== Type assertion segura (con ok) ==")
	procesarPago(PagoTarjeta{NumeroTarjeta: "4111111111111234", Monto: 450.50})
	procesarPago(PagoEfectivo{Monto: 120.00})

	// Type assertion directa (sin ok): paniquea si el tipo no coincide.
	// Solo úsala cuando estés absolutamente seguro del tipo.
	fmt.Println("\n== Type assertion directa ==")
	var p Pagable = PagoEfectivo{Monto: 300}
	efectivo := p.(PagoEfectivo) // ok porque sabemos que es PagoEfectivo
	fmt.Printf("monto en efectivo: S/%.2f\n", efectivo.Monto)

	// any (interface vacía) acepta cualquier tipo. La assertion te permite
	// recuperar el valor original con su tipo real.
	fmt.Println("\n== Assertion sobre any ==")
	valores := []any{42, "Go", true, 4800.00}
	for _, v := range valores {
		if n, ok := v.(int); ok {
			fmt.Println("entero:", n)
		} else if s, ok := v.(string); ok {
			fmt.Println("string:", s)
		} else {
			fmt.Printf("otro tipo: %v\n", v)
		}
	}
}
