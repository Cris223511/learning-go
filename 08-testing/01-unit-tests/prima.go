// Funciones de cálculo de primas que se van a testear en prima_test.go.

package main

import (
	"errors"
	"fmt"
)

type TipoSeguro string

const (
	SOAT      TipoSeguro = "SOAT"
	Vida      TipoSeguro = "Vida"
	Vehicular TipoSeguro = "Vehicular"
)

func CalcularPrima(tipo TipoSeguro, base float64) (float64, error) {
	if base <= 0 {
		return 0, errors.New("la base debe ser mayor a cero")
	}
	switch tipo {
	case SOAT:
		return 120.00, nil
	case Vida:
		return base * 0.015, nil
	case Vehicular:
		return base * 0.03, nil
	default:
		return 0, fmt.Errorf("tipo desconocido: %s", tipo)
	}
}

func AplicarDescuento(prima, porcentaje float64) (float64, error) {
	if porcentaje < 0 || porcentaje > 100 {
		return 0, fmt.Errorf("porcentaje inválido: %.2f", porcentaje)
	}
	return prima * (1 - porcentaje/100), nil
}

func ValidarID(id string) error {
	if id == "" {
		return errors.New("el ID no puede estar vacío")
	}
	if len(id) < 5 {
		return errors.New("el ID debe tener al menos 5 caracteres")
	}
	return nil
}
