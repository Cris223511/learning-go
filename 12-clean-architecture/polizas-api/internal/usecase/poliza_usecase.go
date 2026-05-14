// Acá viven los casos de uso: la lógica de aplicación que orquesta el dominio.
// El usecase sabe QUÉ tiene que pasar (el flujo), pero delega las reglas al dominio
// y la persistencia al repositorio. No sabe nada de HTTP ni de SQL.

package usecase

import (
	"fmt"

	"github.com/cpillihuaman/learning-go/12-clean-architecture/polizas-api/internal/domain"
)

type PolizaUseCase struct {
	// Depende de la interface del dominio, nunca de la implementación concreta.
	// Esto hace que los tests puedan inyectar un mock sin tocar este archivo.
	repo domain.PolizaRepository
}

func Nuevo(repo domain.PolizaRepository) *PolizaUseCase {
	return &PolizaUseCase{repo: repo}
}

func (uc *PolizaUseCase) Crear(clienteID, tipo string, prima float64) (*domain.Poliza, error) {
	p := &domain.Poliza{
		ClienteID: clienteID,
		Tipo:      tipo,
		Prima:     prima,
		Activa:    true,
	}
	// El usecase le pide al dominio que valide: las reglas no le pertenecen al usecase.
	if err := p.Validar(); err != nil {
		return nil, fmt.Errorf("datos inválidos: %w", err)
	}
	creada, err := uc.repo.Guardar(p)
	if err != nil {
		return nil, fmt.Errorf("guardar póliza: %w", err)
	}
	return creada, nil
}

func (uc *PolizaUseCase) ObtenerPorID(id int) (*domain.Poliza, error) {
	p, err := uc.repo.BuscarPorID(id)
	if err != nil {
		return nil, fmt.Errorf("póliza %d: %w", id, err)
	}
	return p, nil
}

func (uc *PolizaUseCase) Listar() ([]*domain.Poliza, error) {
	return uc.repo.BuscarTodos()
}

func (uc *PolizaUseCase) Desactivar(id int) (*domain.Poliza, error) {
	p, err := uc.repo.BuscarPorID(id)
	if err != nil {
		return nil, err
	}
	// La regla "no puedes desactivar lo que ya está inactivo" vive en el dominio.
	if err := p.Desactivar(); err != nil {
		return nil, err
	}
	if err := uc.repo.Actualizar(p); err != nil {
		return nil, fmt.Errorf("actualizar póliza: %w", err)
	}
	return p, nil
}

func (uc *PolizaUseCase) AplicarDescuento(id int, porcentaje float64) (*domain.Poliza, error) {
	p, err := uc.repo.BuscarPorID(id)
	if err != nil {
		return nil, err
	}
	if err := p.AplicarDescuento(porcentaje); err != nil {
		return nil, err
	}
	if err := uc.repo.Actualizar(p); err != nil {
		return nil, fmt.Errorf("actualizar póliza: %w", err)
	}
	return p, nil
}
