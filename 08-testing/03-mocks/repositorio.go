// Interface del repositorio y el servicio que la usa. El servicio no sabe si
// trabaja con PostgreSQL real o un mock: solo conoce la interface.

package main

import "fmt"

type Poliza struct {
	ID      string
	Tipo    string
	Prima   float64
	Activa  bool
}

// PolizaRepository es el contrato. En producción habrá una implementación real con pgx.
// En tests usaremos un mock que implementa la misma interface.
type PolizaRepository interface {
	FindByID(id string) (*Poliza, error)
	FindAll() ([]*Poliza, error)
	Save(p *Poliza) error
}

type PolizaService struct {
	repo PolizaRepository
}

func NewPolizaService(repo PolizaRepository) *PolizaService {
	return &PolizaService{repo: repo}
}

func (s *PolizaService) ObtenerPoliza(id string) (*Poliza, error) {
	if id == "" {
		return nil, fmt.Errorf("id vacío")
	}
	return s.repo.FindByID(id)
}

func (s *PolizaService) ListarActivas() ([]*Poliza, error) {
	todas, err := s.repo.FindAll()
	if err != nil {
		return nil, fmt.Errorf("ListarActivas: %w", err)
	}
	var activas []*Poliza
	for _, p := range todas {
		if p.Activa {
			activas = append(activas, p)
		}
	}
	return activas, nil
}

func (s *PolizaService) CrearPoliza(p *Poliza) error {
	if p.ID == "" || p.Tipo == "" {
		return fmt.Errorf("id y tipo son obligatorios")
	}
	return s.repo.Save(p)
}
