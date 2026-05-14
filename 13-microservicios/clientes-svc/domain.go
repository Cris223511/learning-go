// Entidad Cliente y sus reglas de negocio.
// clientes-svc es el único dueño de estos datos.

package main

import (
	"errors"
	"regexp"
	"time"
)

type Cliente struct {
	ID        string    `json:"id"`
	Nombre    string    `json:"nombre"`
	Email     string    `json:"email"`
	DNI       string    `json:"dni"`
	CreatedAt time.Time `json:"created_at"`
}

var reDNI = regexp.MustCompile(`^\d{8}$`)

func (c *Cliente) Validar() error {
	if c.ID == "" || c.Nombre == "" || c.Email == "" {
		return errors.New("id, nombre y email son obligatorios")
	}
	if !reDNI.MatchString(c.DNI) {
		return errors.New("dni debe tener exactamente 8 dígitos")
	}
	return nil
}

type ClienteRepo interface {
	Guardar(c *Cliente) error
	BuscarPorID(id string) (*Cliente, error)
	BuscarTodos() ([]*Cliente, error)
}

// PolizaResumen es el dato que clientes-svc recibe de polizas-svc via HTTP.
// clientes-svc no tiene acceso a la BD de pólizas, solo a esta representación.
type PolizaResumen struct {
	ID    int     `json:"id"`
	Tipo  string  `json:"tipo"`
	Prima float64 `json:"prima"`
	Activa bool   `json:"activa"`
}
