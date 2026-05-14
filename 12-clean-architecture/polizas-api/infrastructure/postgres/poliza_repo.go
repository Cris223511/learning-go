// Implementación concreta del repositorio con PostgreSQL usando pgx.
// Esta capa sí conoce SQL, pgx y los detalles de la base de datos.
// El dominio y el usecase nunca importan este paquete directamente.

package postgres

import (
	"context"
	"fmt"

	"github.com/cpillihuaman/learning-go/12-clean-architecture/polizas-api/internal/domain"
	"github.com/jackc/pgx/v5/pgxpool"
)

type PolizaRepo struct {
	pool *pgxpool.Pool
}

func NuevoPolizaRepo(pool *pgxpool.Pool) *PolizaRepo {
	return &PolizaRepo{pool: pool}
}

func (r *PolizaRepo) Guardar(p *domain.Poliza) (*domain.Poliza, error) {
	var id int
	err := r.pool.QueryRow(context.Background(),
		`INSERT INTO polizas (cliente_id, tipo, prima, activa)
		 VALUES ($1, $2, $3, $4)
		 RETURNING id, created_at`,
		p.ClienteID, p.Tipo, p.Prima, p.Activa,
	).Scan(&id, &p.CreadaEn)
	if err != nil {
		return nil, fmt.Errorf("insertar póliza: %w", err)
	}
	p.ID = id
	return p, nil
}

func (r *PolizaRepo) BuscarPorID(id int) (*domain.Poliza, error) {
	p := &domain.Poliza{}
	err := r.pool.QueryRow(context.Background(),
		`SELECT id, cliente_id, tipo, prima, activa, created_at
		 FROM polizas WHERE id = $1`, id,
	).Scan(&p.ID, &p.ClienteID, &p.Tipo, &p.Prima, &p.Activa, &p.CreadaEn)
	if err != nil {
		return nil, fmt.Errorf("póliza %d no encontrada: %w", id, err)
	}
	return p, nil
}

func (r *PolizaRepo) BuscarTodos() ([]*domain.Poliza, error) {
	rows, err := r.pool.Query(context.Background(),
		`SELECT id, cliente_id, tipo, prima, activa, created_at
		 FROM polizas ORDER BY id`,
	)
	if err != nil {
		return nil, fmt.Errorf("listar pólizas: %w", err)
	}
	defer rows.Close()

	var polizas []*domain.Poliza
	for rows.Next() {
		p := &domain.Poliza{}
		if err := rows.Scan(&p.ID, &p.ClienteID, &p.Tipo, &p.Prima, &p.Activa, &p.CreadaEn); err != nil {
			return nil, err
		}
		polizas = append(polizas, p)
	}
	return polizas, rows.Err()
}

func (r *PolizaRepo) Actualizar(p *domain.Poliza) error {
	resultado, err := r.pool.Exec(context.Background(),
		`UPDATE polizas SET prima = $1, activa = $2 WHERE id = $3`,
		p.Prima, p.Activa, p.ID,
	)
	if err != nil {
		return fmt.Errorf("actualizar póliza: %w", err)
	}
	if resultado.RowsAffected() == 0 {
		return fmt.Errorf("póliza %d no encontrada", p.ID)
	}
	return nil
}
