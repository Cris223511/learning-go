// Tests unitarios en Go: cada función de test empieza con Test, recibe *testing.T
// y usa t.Errorf/t.Fatalf para reportar fallos. testify/assert simplifica las comparaciones.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestCalcularPrima_SOAT(t *testing.T) {
	prima, err := CalcularPrima(SOAT, 1000)
	// require detiene el test inmediatamente si falla (útil cuando lo siguiente depende de esto).
	require.NoError(t, err)
	assert.Equal(t, 120.00, prima)
}

func TestCalcularPrima_Vida(t *testing.T) {
	prima, err := CalcularPrima(Vida, 100_000)
	require.NoError(t, err)
	assert.Equal(t, 1500.00, prima)
}

func TestCalcularPrima_BaseInvalida(t *testing.T) {
	_, err := CalcularPrima(SOAT, -100)
	// assert.Error verifica que haya un error. assert.EqualError compara el mensaje exacto.
	assert.Error(t, err)
	assert.EqualError(t, err, "la base debe ser mayor a cero")
}

func TestCalcularPrima_TipoDesconocido(t *testing.T) {
	_, err := CalcularPrima("Dental", 1000)
	assert.Error(t, err)
	assert.Contains(t, err.Error(), "tipo desconocido")
}

func TestAplicarDescuento(t *testing.T) {
	resultado, err := AplicarDescuento(100, 10)
	require.NoError(t, err)
	assert.Equal(t, 90.00, resultado)
}

func TestAplicarDescuento_PorcentajeInvalido(t *testing.T) {
	_, err := AplicarDescuento(100, 110)
	assert.Error(t, err)

	_, err = AplicarDescuento(100, -5)
	assert.Error(t, err)
}

func TestValidarID(t *testing.T) {
	assert.NoError(t, ValidarID("POL-001"))
	assert.Error(t, ValidarID(""))
	assert.Error(t, ValidarID("AB"))
}
