// Acá veremos database/sql: la interface estándar de Go para bases de datos.
// Usa pgx como driver por debajo pero trabaja con la abstracción de la stdlib.
//
// Requisito: PostgreSQL corriendo. Con Docker:
//   docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres --name pg-aprendizago postgres:16
//   docker exec -it pg-aprendizago psql -U postgres -c "CREATE DATABASE aprendizago;"
//
// Correr: go run ./11-postgresql/01-database-sql

package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"time"

	_ "github.com/jackc/pgx/v5/stdlib" // el _ importa el driver sin usarlo directamente
)

const dsn = "postgres://postgres:postgres@localhost:5432/aprendizago?sslmode=disable"

type Cliente struct {
	ID        int
	Nombre    string
	Email     string
	DNI       string
	CreatedAt time.Time
}

func main() {
	// sql.Open no abre la conexión todavía, solo valida el DSN.
	// La conexión real ocurre en el primer query o al llamar Ping.
	db, err := sql.Open("pgx", dsn)
	if err != nil {
		log.Fatal("error al configurar DB:", err)
	}
	defer db.Close()

	// Configuración del pool de conexiones.
	db.SetMaxOpenConns(10)
	db.SetMaxIdleConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	ctx := context.Background()

	// Ping verifica que la conexión real funciona.
	if err := db.PingContext(ctx); err != nil {
		log.Fatal("no se pudo conectar a PostgreSQL:", err)
	}
	fmt.Println("conexión exitosa a PostgreSQL")

	// Crear tabla si no existe.
	crearTabla(ctx, db)

	// Insertar un cliente.
	id := insertar(ctx, db, "Christopher", "c@correo.com", "12345678")
	fmt.Println("cliente insertado con ID:", id)

	// Consultar todos.
	clientes := listar(ctx, db)
	for _, c := range clientes {
		fmt.Printf("  [%d] %s | %s\n", c.ID, c.Nombre, c.Email)
	}

	// Consultar uno por ID.
	c, err := buscarPorID(ctx, db, id)
	if err != nil {
		fmt.Println("no encontrado:", err)
	} else {
		fmt.Printf("encontrado: %+v\n", c)
	}
}

func crearTabla(ctx context.Context, db *sql.DB) {
	_, err := db.ExecContext(ctx, `
		CREATE TABLE IF NOT EXISTS clientes (
			id         SERIAL PRIMARY KEY,
			nombre     VARCHAR(100) NOT NULL,
			email      VARCHAR(150) NOT NULL UNIQUE,
			dni        VARCHAR(8)   NOT NULL UNIQUE,
			created_at TIMESTAMP    DEFAULT NOW()
		)
	`)
	if err != nil {
		log.Fatal("error creando tabla:", err)
	}
}

func insertar(ctx context.Context, db *sql.DB, nombre, email, dni string) int {
	// QueryRowContext para queries que retornan una fila. RETURNING id recupera el ID generado.
	var id int
	err := db.QueryRowContext(ctx,
		`INSERT INTO clientes (nombre, email, dni) VALUES ($1, $2, $3) RETURNING id`,
		nombre, email, dni,
	).Scan(&id)
	if err != nil {
		log.Fatal("error insertando:", err)
	}
	return id
}

func listar(ctx context.Context, db *sql.DB) []Cliente {
	// QueryContext para queries que retornan múltiples filas.
	rows, err := db.QueryContext(ctx, `SELECT id, nombre, email, dni, created_at FROM clientes ORDER BY id`)
	if err != nil {
		log.Fatal("error consultando:", err)
	}
	defer rows.Close() // siempre cerrar rows para liberar la conexión al pool

	var clientes []Cliente
	for rows.Next() {
		var c Cliente
		// Scan mapea cada columna a un campo del struct, en el mismo orden del SELECT.
		if err := rows.Scan(&c.ID, &c.Nombre, &c.Email, &c.DNI, &c.CreatedAt); err != nil {
			log.Fatal("error escaneando:", err)
		}
		clientes = append(clientes, c)
	}
	// rows.Err() detecta errores que ocurrieron durante la iteración.
	if err := rows.Err(); err != nil {
		log.Fatal("error en rows:", err)
	}
	return clientes
}

func buscarPorID(ctx context.Context, db *sql.DB, id int) (Cliente, error) {
	var c Cliente
	err := db.QueryRowContext(ctx,
		`SELECT id, nombre, email, dni, created_at FROM clientes WHERE id = $1`, id,
	).Scan(&c.ID, &c.Nombre, &c.Email, &c.DNI, &c.CreatedAt)
	// sql.ErrNoRows es el sentinel error cuando la query no devuelve filas.
	if err == sql.ErrNoRows {
		return Cliente{}, fmt.Errorf("cliente %d no encontrado", id)
	}
	return c, err
}
