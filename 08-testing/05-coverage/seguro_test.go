// Tests de cobertura: algunos casos cubiertos, otros no, para mostrar el reporte.
// Ver cobertura: go test ./08-testing/05-coverage -cover
// Reporte visual:
//   go test ./08-testing/05-coverage -coverprofile=coverage.out
//   go tool cover -html=coverage.out

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestEvaluarRiesgo_Bajo(t *testing.T) {
	nivel, factor, err := EvaluarRiesgo(35, 0)
	require.NoError(t, err)
	assert.Equal(t, RiesgoBajo, nivel)
	assert.Equal(t, 1.0, factor)
}

func TestEvaluarRiesgo_EdadInvalida(t *testing.T) {
	_, _, err := EvaluarRiesgo(-1, 0)
	assert.Error(t, err)
}

func TestEvaluarRiesgo_Medio(t *testing.T) {
	nivel, _, err := EvaluarRiesgo(22, 0) // menor de 25 → medio
	require.NoError(t, err)
	assert.Equal(t, RiesgoMedio, nivel)
}

// RiesgoAlto y historial negativo quedan sin cubrir intencionalmente.
// Corre go test -cover para ver el porcentaje y -coverprofile para el reporte HTML.

func TestCalcularPrimaConRiesgo(t *testing.T) {
	prima, err := CalcularPrimaConRiesgo(400, 35, 0)
	require.NoError(t, err)
	assert.Equal(t, 400.0, prima)
}
