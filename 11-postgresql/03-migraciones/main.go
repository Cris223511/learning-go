// Acá veremos migraciones con golang-migrate: versionar cambios de esquema
// igual que Git versiona código. Cada migración tiene un UP y un DOWN.
//
// Requisito: PostgreSQL corriendo (ver instrucciones en 01-database-sql).
// Correr: go run ./11-postgresql/03-migraciones
//
// Los archivos SQL están en la carpeta migrations/ y se embeben en el binario
// con //go:embed para no depender de rutas en tiempo de ejecución.

package main

import (
	"embed"
	"fmt"
	"log"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
)

const dsn = "postgres://postgres:postgres@localhost:5432/aprendizago?sslmode=disable"

// //go:embed embebe todos los archivos .sql de la carpeta migrations en el binario.
// Así el ejecutable funciona en cualquier máquina sin necesitar los archivos sueltos.
//
//go:embed migrations/*.sql
var archivosSQL embed.FS

func main() {
	// iofs.New lee las migraciones desde el FS embebido.
	fuente, err := iofs.New(archivosSQL, "migrations")
	if err != nil {
		log.Fatal("error leyendo migraciones:", err)
	}

	m, err := migrate.NewWithSourceInstance("iofs", fuente, dsn)
	if err != nil {
		log.Fatal("error configurando migrate:", err)
	}
	defer m.Close()

	// Versión actual antes de migrar.
	version, dirty, _ := m.Version()
	fmt.Printf("versión actual: %d | dirty: %v\n", version, dirty)

	// Up() aplica todas las migraciones pendientes.
	// Si ya están todas aplicadas, retorna migrate.ErrNoChange.
	fmt.Println("aplicando migraciones...")
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("error aplicando migraciones:", err)
	} else if err == migrate.ErrNoChange {
		fmt.Println("ya estaba al día, sin cambios")
	} else {
		fmt.Println("migraciones aplicadas con éxito")
	}

	version, _, _ = m.Version()
	fmt.Printf("versión después de Up: %d\n", version)

	// Steps(-1) revierte la última migración (equivale a Down de la última).
	// En producción nunca hagas Down automático, solo en desarrollo.
	fmt.Println("\nreviertiendo la última migración...")
	if err := m.Steps(-1); err != nil {
		log.Println("error en rollback:", err)
	} else {
		version, _, _ = m.Version()
		fmt.Printf("versión después de Steps(-1): %d\n", version)
	}

	// Volver a aplicar para dejar el esquema completo.
	_ = m.Up()
	version, _, _ = m.Version()
	fmt.Printf("versión final: %d\n", version)
}
