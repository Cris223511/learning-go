// Repositorio de pólizas usando pgx directamente. A diferencia de database/sql,
// pgx expone tipos nativos de PostgreSQL y errores más descriptivos.

package main

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

type Poliza struct {
	ID         int
	ClienteID  int
	Tipo       string
	Prima      float64
	Activa     bool
	CreatedAt  time.Time
}

type PolizaRepo struct {
	pool *pgxpool.Pool
}

func NuevoPolizaRepo(pool *pgxpool.Pool) *PolizaRepo {
	return &PolizaRepo{pool: pool}
}

func (r *PolizaRepo) CrearTabla(ctx context.Context) error {
	_, err := r.pool.Exec(ctx, `
		CREATE TABLE IF NOT EXISTS polizas (
			id          SERIAL PRIMARY KEY,
			cliente_id  INTEGER NOT NULL,
			tipo        VARCHAR(50) NOT NULL,
			prima       NUMERIC(10,2) NOT NULL,
			activa      BOOLEAN DEFAULT true,
			created_at  TIMESTAMP DEFAULT NOW()
		)
	`)
	return err
}

func (r *PolizaRepo) Insertar(ctx context.Context, p Poliza) (int, error) {
	var id int
	err := r.pool.QueryRow(ctx,
		`INSERT INTO polizas (cliente_id, tipo, prima, activa) VALUES ($1, $2, $3, $4) RETURNING id`,
		p.ClienteID, p.Tipo, p.Prima, p.Activa,
	).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("insertar póliza: %w", err)
	}
	return id, nil
}

func (r *PolizaRepo) Listar(ctx context.Context) ([]Poliza, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, cliente_id, tipo, prima, activa, created_at FROM polizas ORDER BY id`,
	)
	if err != nil {
		return nil, fmt.Errorf("listar pólizas: %w", err)
	}
	// pgx.CollectRows es la forma moderna de convertir rows a un slice sin el bucle manual.
	polizas, err := pgx.CollectRows(rows, pgx.RowToStructByName[Poliza])
	if err != nil {
		return nil, fmt.Errorf("scanear pólizas: %w", err)
	}
	return polizas, nil
}

func (r *PolizaRepo) BuscarPorID(ctx context.Context, id int) (Poliza, error) {
	rows, err := r.pool.Query(ctx,
		`SELECT id, cliente_id, tipo, prima, activa, created_at FROM polizas WHERE id = $1`, id,
	)
	if err != nil {
		return Poliza{}, err
	}
	// pgx.CollectOneRow devuelve pgx.ErrNoRows si no hay resultados.
	p, err := pgx.CollectOneRow(rows, pgx.RowToStructByName[Poliza])
	if err != nil {
		return Poliza{}, fmt.Errorf("póliza %d: %w", id, err)
	}
	return p, nil
}

func (r *PolizaRepo) Actualizar(ctx context.Context, id int, activa bool) error {
	resultado, err := r.pool.Exec(ctx,
		`UPDATE polizas SET activa = $1 WHERE id = $2`, activa, id,
	)
	if err != nil {
		return fmt.Errorf("actualizar póliza: %w", err)
	}
	// RowsAffected verifica si realmente se modificó algo.
	if resultado.RowsAffected() == 0 {
		return fmt.Errorf("póliza %d no encontrada", id)
	}
	return nil
}
