// Mocks con testify/mock: generan automáticamente un objeto que implementa
// la interface y te permite controlar qué retorna cada método en el test.

package main

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

// MockPolizaRepository implementa PolizaRepository. testify/mock usa reflection
// para registrar llamadas y retornos configurados por el test.
type MockPolizaRepository struct {
	mock.Mock
}

func (m *MockPolizaRepository) FindByID(id string) (*Poliza, error) {
	// Called registra que este método fue llamado con estos argumentos
	// y retorna lo que el test configuró con On(...).Return(...).
	args := m.Called(id)
	if args.Get(0) == nil {
		return nil, args.Error(1)
	}
	return args.Get(0).(*Poliza), args.Error(1)
}

func (m *MockPolizaRepository) FindAll() ([]*Poliza, error) {
	args := m.Called()
	return args.Get(0).([]*Poliza), args.Error(1)
}

func (m *MockPolizaRepository) Save(p *Poliza) error {
	args := m.Called(p)
	return args.Error(0)
}

func TestObtenerPoliza_Exitoso(t *testing.T) {
	repo := new(MockPolizaRepository)
	// On define qué retorna el mock cuando se llame FindByID con "POL-001".
	repo.On("FindByID", "POL-001").Return(&Poliza{ID: "POL-001", Tipo: "SOAT", Activa: true}, nil)

	svc := NewPolizaService(repo)
	poliza, err := svc.ObtenerPoliza("POL-001")

	require.NoError(t, err)
	assert.Equal(t, "POL-001", poliza.ID)
	// AssertExpectations verifica que todos los On() configurados fueron llamados.
	repo.AssertExpectations(t)
}

func TestObtenerPoliza_IDVacio(t *testing.T) {
	repo := new(MockPolizaRepository)
	svc := NewPolizaService(repo)

	_, err := svc.ObtenerPoliza("")
	assert.Error(t, err)
	// El repo nunca debería ser llamado si el id está vacío.
	repo.AssertNotCalled(t, "FindByID", mock.Anything)
}

func TestObtenerPoliza_NoEncontrada(t *testing.T) {
	repo := new(MockPolizaRepository)
	repo.On("FindByID", "POL-999").Return(nil, errors.New("no encontrado"))

	svc := NewPolizaService(repo)
	_, err := svc.ObtenerPoliza("POL-999")
	assert.Error(t, err)
	repo.AssertExpectations(t)
}

func TestListarActivas(t *testing.T) {
	repo := new(MockPolizaRepository)
	repo.On("FindAll").Return([]*Poliza{
		{ID: "POL-001", Activa: true},
		{ID: "POL-002", Activa: false},
		{ID: "POL-003", Activa: true},
	}, nil)

	svc := NewPolizaService(repo)
	activas, err := svc.ListarActivas()

	require.NoError(t, err)
	assert.Len(t, activas, 2)
	repo.AssertExpectations(t)
}

func TestCrearPoliza(t *testing.T) {
	repo := new(MockPolizaRepository)
	poliza := &Poliza{ID: "POL-004", Tipo: "Vehicular", Prima: 890}
	// mock.Anything acepta cualquier valor para ese argumento.
	repo.On("Save", poliza).Return(nil)

	svc := NewPolizaService(repo)
	err := svc.CrearPoliza(poliza)
	assert.NoError(t, err)
	repo.AssertExpectations(t)
}
