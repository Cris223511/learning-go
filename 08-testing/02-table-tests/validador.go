// Funciones de validación que se testearán con table-driven tests.

package main

import (
	"fmt"
	"regexp"
	"strings"
)

var reEmail = regexp.MustCompile(`^[a-zA-Z0-9._%+\-]+@[a-zA-Z0-9.\-]+\.[a-zA-Z]{2,}$`)
var reDNI = regexp.MustCompile(`^\d{8}$`)

func ValidarEmail(email string) error {
	if email == "" {
		return fmt.Errorf("email vacío")
	}
	if !reEmail.MatchString(email) {
		return fmt.Errorf("email inválido: %q", email)
	}
	return nil
}

func ValidarDNI(dni string) error {
	if !reDNI.MatchString(dni) {
		return fmt.Errorf("DNI inválido: debe tener 8 dígitos")
	}
	return nil
}

func NormalizarNombre(nombre string) string {
	return strings.TrimSpace(strings.ToUpper(nombre))
}
