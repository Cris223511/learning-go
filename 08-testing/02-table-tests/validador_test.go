// Table-driven tests: el patrón más idiomático de Go para tests.
// Defines una tabla de casos y los recorres con un bucle. Agregar un caso nuevo
// es solo agregar una línea a la tabla, sin duplicar código de test.

package main

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestValidarEmail(t *testing.T) {
	// Cada caso tiene un nombre, la entrada y lo que se espera.
	casos := []struct {
		nombre    string
		email     string
		esperaErr bool
	}{
		{"válido básico", "user@mail.com", false},
		{"válido corporativo", "christopher@correo.com", false},
		{"vacío", "", true},
		{"sin arroba", "usermail.com", true},
		{"sin dominio", "user@", true},
		{"con espacios", "user @mail.com", true},
	}

	for _, tc := range casos {
		// t.Run crea un sub-test con nombre. En la salida aparece como TestValidarEmail/válido_básico.
		// Si falla solo ese caso, puedes correrlo solo con: go test -run TestValidarEmail/válido_básico
		t.Run(tc.nombre, func(t *testing.T) {
			err := ValidarEmail(tc.email)
			if tc.esperaErr {
				assert.Error(t, err, "debería fallar con: %q", tc.email)
			} else {
				assert.NoError(t, err, "no debería fallar con: %q", tc.email)
			}
		})
	}
}

func TestValidarDNI(t *testing.T) {
	casos := []struct {
		nombre    string
		dni       string
		esperaErr bool
	}{
		{"8 dígitos válidos", "12345678", false},
		{"vacío", "", true},
		{"7 dígitos", "1234567", true},
		{"9 dígitos", "123456789", true},
		{"con letras", "1234567A", true},
		{"con guion", "1234-567", true},
	}

	for _, tc := range casos {
		t.Run(tc.nombre, func(t *testing.T) {
			err := ValidarDNI(tc.dni)
			if tc.esperaErr {
				assert.Error(t, err)
			} else {
				assert.NoError(t, err)
			}
		})
	}
}

func TestNormalizarNombre(t *testing.T) {
	casos := []struct {
		entrada  string
		esperado string
	}{
		{"christopher", "CHRISTOPHER"},
		{"  ana garcía  ", "ANA GARCÍA"},
		{"Luis Torres", "LUIS TORRES"},
	}

	for _, tc := range casos {
		t.Run(tc.entrada, func(t *testing.T) {
			assert.Equal(t, tc.esperado, NormalizarNombre(tc.entrada))
		})
	}
}
