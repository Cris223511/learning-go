// Función con múltiples ramas para demostrar coverage parcial.
// Algunas ramas quedarán sin cubrir intencionalmente.

package main

import "fmt"

type NivelRiesgo string

const (
	RiesgoBajo  NivelRiesgo = "bajo"
	RiesgoMedio NivelRiesgo = "medio"
	RiesgoAlto  NivelRiesgo = "alto"
)

func EvaluarRiesgo(edad int, historialSiniestros int) (NivelRiesgo, float64, error) {
	if edad < 0 || edad > 120 {
		return "", 0, fmt.Errorf("edad inválida: %d", edad)
	}
	if historialSiniestros < 0 {
		return "", 0, fmt.Errorf("historial negativo: %d", historialSiniestros)
	}

	switch {
	case historialSiniestros == 0 && edad >= 25 && edad <= 60:
		return RiesgoBajo, 1.0, nil
	case historialSiniestros <= 2 || (edad < 25 || edad > 60):
		return RiesgoMedio, 1.35, nil
	default:
		return RiesgoAlto, 1.75, nil
	}
}

func CalcularPrimaConRiesgo(primaBase float64, edad, siniestros int) (float64, error) {
	_, factor, err := EvaluarRiesgo(edad, siniestros)
	if err != nil {
		return 0, err
	}
	return primaBase * factor, nil
}
