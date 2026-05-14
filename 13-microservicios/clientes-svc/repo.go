// Repositorio de clientes con PostgreSQL.

package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type postgresClienteRepo struct {
	pool *pgxpool.Pool
}

func nuevoClienteRepo(pool *pgxpool.Pool) ClienteRepo {
	return &postgresClienteRepo{pool: pool}
}

func (r *postgresClienteRepo) Guardar(c *Cliente) error {
	_, err := r.pool.Exec(context.Background(),
		`INSERT INTO clientes (id, nombre, email, dni) VALUES ($1, $2, $3, $4)`,
		c.ID, c.Nombre, c.Email, c.DNI,
	)
	if err != nil {
		return fmt.Errorf("guardar cliente: %w", err)
	}
	return nil
}

func (r *postgresClienteRepo) BuscarPorID(id string) (*Cliente, error) {
	c := &Cliente{}
	err := r.pool.QueryRow(context.Background(),
		`SELECT id, nombre, email, dni, created_at FROM clientes WHERE id = $1`, id,
	).Scan(&c.ID, &c.Nombre, &c.Email, &c.DNI, &c.CreatedAt)
	if err != nil {
		return nil, fmt.Errorf("cliente %q no encontrado: %w", id, err)
	}
	return c, nil
}

func (r *postgresClienteRepo) BuscarTodos() ([]*Cliente, error) {
	rows, err := r.pool.Query(context.Background(),
		`SELECT id, nombre, email, dni, created_at FROM clientes ORDER BY id`,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var lista []*Cliente
	for rows.Next() {
		c := &Cliente{}
		if err := rows.Scan(&c.ID, &c.Nombre, &c.Email, &c.DNI, &c.CreatedAt); err != nil {
			return nil, err
		}
		lista = append(lista, c)
	}
	return lista, rows.Err()
}
