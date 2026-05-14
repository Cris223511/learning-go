// Entidad Poliza y las reglas de negocio que le pertenecen.
// Esta capa no sabe nada de HTTP, SQL ni Docker.

package main

import (
	"errors"
	"fmt"
	"time"
)

type Poliza struct {
	ID        int       `json:"id"`
	ClienteID string    `json:"cliente_id"`
	Tipo      string    `json:"tipo"`
	Prima     float64   `json:"prima"`
	Activa    bool      `json:"activa"`
	CreatedAt time.Time `json:"created_at"`
}

func (p *Poliza) Validar() error {
	if p.ClienteID == "" {
		return errors.New("cliente_id es obligatorio")
	}
	tipos := map[string]bool{"SOAT": true, "Vida": true, "Vehicular": true, "Hogar": true}
	if !tipos[p.Tipo] {
		return fmt.Errorf("tipo inválido: %q", p.Tipo)
	}
	if p.Prima <= 0 {
		return errors.New("prima debe ser mayor a cero")
	}
	return nil
}

func (p *Poliza) Desactivar() error {
	if !p.Activa {
		return errors.New("la póliza ya está inactiva")
	}
	p.Activa = false
	return nil
}

// PolizaRepo define qué necesita el servicio de la BD.
// La implementación concreta está en repo.go.
type PolizaRepo interface {
	Guardar(p *Poliza) (*Poliza, error)
	BuscarPorID(id int) (*Poliza, error)
	BuscarPorCliente(clienteID string) ([]*Poliza, error)
	BuscarTodos() ([]*Poliza, error)
	Actualizar(p *Poliza) error
}
