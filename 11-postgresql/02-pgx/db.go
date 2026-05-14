// Configuración del pool de conexiones con pgxpool. pgx directamente (sin database/sql)
// ofrece mejor rendimiento y acceso a tipos nativos de PostgreSQL.

package main

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

const dsn = "postgres://postgres:postgres@localhost:5432/aprendizago?sslmode=disable"

// NuevoPool crea el pool de conexiones listo para usar.
// pgxpool maneja reconexiones automáticamente y distribuye las queries entre conexiones disponibles.
func NuevoPool(ctx context.Context) (*pgxpool.Pool, error) {
	config, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return nil, fmt.Errorf("parsear config: %w", err)
	}
	// Ajustes del pool.
	config.MaxConns = 10
	config.MinConns = 2

	pool, err := pgxpool.NewWithConfig(ctx, config)
	if err != nil {
		return nil, fmt.Errorf("crear pool: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		return nil, fmt.Errorf("ping fallido: %w", err)
	}
	return pool, nil
}
