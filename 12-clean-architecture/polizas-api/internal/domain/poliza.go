// Acá está el corazón del proyecto: la entidad Poliza y las reglas de negocio.
// Esta capa no importa Gin, pgx, ni ningún framework externo.
// Si cambias de PostgreSQL a MongoDB, o de Gin a otro router, esta capa no se toca.

package domain

import (
	"errors"
	"fmt"
	"time"
)

type Poliza struct {
	ID        int
	ClienteID string
	Tipo      string
	Prima     float64
	Activa    bool
	CreadaEn  time.Time
}

// Las reglas de negocio viven acá, no en el handler ni en el usecase.
// Así se pueden testear sin levantar un servidor ni conectarse a la BD.
func (p *Poliza) Desactivar() error {
	if !p.Activa {
		return errors.New("la póliza ya está inactiva")
	}
	p.Activa = false
	return nil
}

func (p *Poliza) AplicarDescuento(porcentaje float64) error {
	if porcentaje <= 0 || porcentaje > 50 {
		return fmt.Errorf("el descuento debe estar entre 1 y 50, recibido: %.2f", porcentaje)
	}
	p.Prima = p.Prima * (1 - porcentaje/100)
	return nil
}

func (p *Poliza) Validar() error {
	if p.ClienteID == "" {
		return errors.New("cliente_id es obligatorio")
	}
	tipos := map[string]bool{"SOAT": true, "Vida": true, "Vehicular": true, "Hogar": true}
	if !tipos[p.Tipo] {
		return fmt.Errorf("tipo inválido: %q. Valores permitidos: SOAT, Vida, Vehicular, Hogar", p.Tipo)
	}
	if p.Prima <= 0 {
		return errors.New("prima debe ser mayor a cero")
	}
	return nil
}

// PolizaRepository define QUÉ operaciones necesita el dominio sobre la persistencia.
// El dominio no sabe si los datos están en PostgreSQL, memoria o cualquier otra cosa.
// La implementación concreta vive en infrastructure/postgres, no acá.
type PolizaRepository interface {
	Guardar(p *Poliza) (*Poliza, error)
	BuscarPorID(id int) (*Poliza, error)
	BuscarTodos() ([]*Poliza, error)
	Actualizar(p *Poliza) error
}
