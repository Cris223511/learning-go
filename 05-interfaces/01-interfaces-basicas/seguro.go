// Definición de la interface Seguro y sus implementaciones concretas.

package main

import "fmt"

// Una interface en Go es un contrato: define qué métodos debe tener un tipo,
// no cómo los implementa. No hay "implements": si el tipo tiene los métodos, cumple la interface.
type Seguro interface {
	Descripcion() string
	PrimaAnual() float64
}

// SOAT y Vida son tipos distintos que ambos cumplen Seguro
// sin declararlo explícitamente. Go lo detecta automáticamente.

type SOAT struct {
	Placa    string
	Propietario string
}

func (s SOAT) Descripcion() string {
	return fmt.Sprintf("SOAT | placa: %s | propietario: %s", s.Placa, s.Propietario)
}

func (s SOAT) PrimaAnual() float64 {
	return 120.00
}

type Vida struct {
	Asegurado string
	Capital   float64
}

func (v Vida) Descripcion() string {
	return fmt.Sprintf("Vida | asegurado: %s | capital: S/%.2f", v.Asegurado, v.Capital)
}

func (v Vida) PrimaAnual() float64 {
	return v.Capital * 0.015
}

// Vehicular implementa Seguro y tiene un campo extra que no está en la interface.
type Vehicular struct {
	Placa    string
	Marca    string
	ValorVehiculo float64
}

func (v Vehicular) Descripcion() string {
	return fmt.Sprintf("Vehicular | %s %s | valor: S/%.2f", v.Marca, v.Placa, v.ValorVehiculo)
}

func (v Vehicular) PrimaAnual() float64 {
	return v.ValorVehiculo * 0.03
}
