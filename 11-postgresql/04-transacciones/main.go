// Acá veremos transacciones en PostgreSQL con pgx: cómo agrupar varias operaciones
// en una unidad atómica que o se completa entera o no se aplica ninguna.
//
// Requisito: PostgreSQL corriendo con las tablas creadas (correr 03-migraciones primero).
// Correr: go run ./11-postgresql/04-transacciones

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
)

const dsn = "postgres://postgres:postgres@localhost:5432/aprendizago?sslmode=disable"

type Cliente struct {
	ID     int
	Nombre string
	Email  string
	DNI    string
}

type Poliza struct {
	ID        int
	ClienteID int
	Tipo      string
	Prima     float64
}

func main() {
	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("error conectando:", err)
	}
	defer pool.Close()
	fmt.Println("conectado a PostgreSQL")

	// Caso 1: transacción exitosa. Cliente + póliza en un solo commit.
	fmt.Println("\n== Transacción exitosa ==")
	if err := registrarClienteConPoliza(ctx, pool, Cliente{
		Nombre: "Ana García",
		Email:  "ana@mail.com",
		DNI:    "87654321",
	}, Poliza{
		Tipo:  "Vida",
		Prima: 1500.00,
	}); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("cliente y póliza registrados correctamente")
	}

	// Caso 2: transacción fallida. Simula un error a mitad para ver el rollback.
	fmt.Println("\n== Transacción con rollback ==")
	if err := registrarConError(ctx, pool); err != nil {
		fmt.Println("rollback ejecutado, error:", err)
	}
}

// registrarClienteConPoliza inserta un cliente y su póliza inicial en una sola transacción.
// Si cualquiera de los dos falla, ninguno queda en la BD.
func registrarClienteConPoliza(ctx context.Context, pool *pgxpool.Pool, c Cliente, p Poliza) error {
	// pgx.BeginTx inicia la transacción con opciones configurables (nivel de aislamiento, etc.).
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return fmt.Errorf("iniciar tx: %w", err)
	}
	// El defer con rollback es el patrón estándar: si algo falla y no llegamos al Commit,
	// el rollback se ejecuta automáticamente y deshace todo lo hecho en la tx.
	defer tx.Rollback(ctx)

	var clienteID int
	err = tx.QueryRow(ctx,
		`INSERT INTO clientes (nombre, email, dni) VALUES ($1, $2, $3) RETURNING id`,
		c.Nombre, c.Email, c.DNI,
	).Scan(&clienteID)
	if err != nil {
		return fmt.Errorf("insertar cliente: %w", err)
	}
	fmt.Printf("  cliente insertado con ID=%d\n", clienteID)

	var polizaID int
	err = tx.QueryRow(ctx,
		`INSERT INTO polizas (cliente_id, tipo, prima) VALUES ($1, $2, $3) RETURNING id`,
		clienteID, p.Tipo, p.Prima,
	).Scan(&polizaID)
	if err != nil {
		return fmt.Errorf("insertar póliza: %w", err)
	}
	fmt.Printf("  póliza insertada con ID=%d\n", polizaID)

	// Commit aplica todas las operaciones de la tx. Solo si llegamos acá sin error.
	if err := tx.Commit(ctx); err != nil {
		return fmt.Errorf("commit: %w", err)
	}
	return nil
}

// registrarConError simula un fallo a mitad de la transacción para mostrar el rollback.
func registrarConError(ctx context.Context, pool *pgxpool.Pool) error {
	tx, err := pool.BeginTx(ctx, pgx.TxOptions{})
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	// Primera operación: ok.
	_, err = tx.Exec(ctx,
		`INSERT INTO clientes (nombre, email, dni) VALUES ($1, $2, $3)`,
		"Luis Torres", "luis@mail.com", "11223344",
	)
	if err != nil {
		return fmt.Errorf("primer insert: %w", err)
	}
	fmt.Println("  primer insert ejecutado (pendiente de commit)")

	// Segunda operación: falla intencionalmente (email duplicado).
	_, err = tx.Exec(ctx,
		`INSERT INTO clientes (nombre, email, dni) VALUES ($1, $2, $3)`,
		"Luis Clon", "luis@mail.com", "99887766", // mismo email → viola UNIQUE
	)
	if err != nil {
		// El defer Rollback() revertirá también el primer insert.
		return fmt.Errorf("segundo insert falló, se revertirá todo: %w", err)
	}

	return tx.Commit(ctx)
}
