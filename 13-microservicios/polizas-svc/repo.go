// Implementación del repositorio con PostgreSQL.
// Es la única capa de polizas-svc que habla con la base de datos.

package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresRepo struct {
	pool *pgxpool.Pool
}

func nuevoRepo(pool *pgxpool.Pool) PolizaRepo {
	return &postgresRepo{pool: pool}
}

func (r *postgresRepo) Guardar(p *Poliza) (*Poliza, error) {
	err := r.pool.QueryRow(context.Background(),
		`INSERT INTO polizas (cliente_id, tipo, prima, activa)
		 VALUES ($1, $2, $3, $4) RETURNING id, created_at`,
		p.ClienteID, p.Tipo, p.Prima, p.Activa,
	).Scan(&p.ID, &p.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("guardar póliza: %w", err)
	}
	return p, nil
}

func (r *postgresRepo) BuscarPorID(id int) (*Poliza, error) {
	p := &Poliza{}
	err := r.pool.QueryRow(context.Background(),
		`SELECT id, cliente_id, tipo, prima, activa, created_at
		 FROM polizas WHERE id = $1`, id,
	).Scan(&p.ID, &p.ClienteID, &p.Tipo, &p.Prima, &p.Activa, &p.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("póliza %d no encontrada: %w", id, err)
	}
	return p, nil
}

func (r *postgresRepo) BuscarPorCliente(clienteID string) ([]*Poliza, error) {
	rows, err := r.pool.Query(context.Background(),
		`SELECT id, cliente_id, tipo, prima, activa, created_at
		 FROM polizas WHERE cliente_id = $1 ORDER BY id`, clienteID,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return escanearPolizas(rows)
}

func (r *postgresRepo) BuscarTodos() ([]*Poliza, error) {
	rows, err := r.pool.Query(context.Background(),
		`SELECT id, cliente_id, tipo, prima, activa, created_at FROM polizas ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	return escanearPolizas(rows)
}

func (r *postgresRepo) Actualizar(p *Poliza) error {
	res, err := r.pool.Exec(context.Background(),
		`UPDATE polizas SET prima=$1, activa=$2 WHERE id=$3`,
		p.Prima, p.Activa, p.ID,
	)
	if err != nil {
		return fmt.Errorf("actualizar póliza: %w", err)
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("póliza %d no encontrada", p.ID)
	}
	return nil
}

func escanearPolizas(rows interface{ Next() bool; Scan(...any) error; Err() error }) ([]*Poliza, error) {
	var lista []*Poliza
	for rows.Next() {
		p := &Poliza{}
		if err := rows.Scan(&p.ID, &p.ClienteID, &p.Tipo, &p.Prima, &p.Activa, &p.CreatedAt); err != nil {
			return nil, err
		}
		lista = append(lista, p)
	}
	return lista, rows.Err()
}
