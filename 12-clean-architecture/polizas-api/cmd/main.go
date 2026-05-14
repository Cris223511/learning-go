// Punto de entrada de la API. Acá se conectan todas las capas:
// primero la infraestructura (BD), luego el usecase, luego el handler.
// Este archivo es el único que conoce todos los paquetes del proyecto.

package main

import (
	"context"
	"fmt"
	"log"

	"github.com/cpillihuaman/learning-go/12-clean-architecture/polizas-api/internal/handler"
	"github.com/cpillihuaman/learning-go/12-clean-architecture/polizas-api/internal/usecase"
	"github.com/cpillihuaman/learning-go/12-clean-architecture/polizas-api/infrastructure/migrations"
	"github.com/cpillihuaman/learning-go/12-clean-architecture/polizas-api/infrastructure/postgres"
	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
)

const dsn = "postgres://postgres:postgres@localhost:5432/aprendizago?sslmode=disable"

func main() {
	ctx := context.Background()

	// 1. Conectar a PostgreSQL.
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		log.Fatal("error conectando a PostgreSQL:", err)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		log.Fatal("no se pudo alcanzar PostgreSQL. ¿Está corriendo el contenedor?", err)
	}
	fmt.Println("conectado a PostgreSQL")

	// 2. Ejecutar migraciones automáticamente al arrancar.
	ejecutarMigraciones()

	// 3. Dependency Injection: construimos de adentro hacia afuera.
	// La infraestructura se le inyecta al usecase, y el usecase al handler.
	// Ninguna capa busca sus dependencias sola.
	repo := postgres.NuevoPolizaRepo(pool)
	uc := usecase.Nuevo(repo)
	ph := handler.NuevoPolizaHandler(uc)

	// 4. Levantar el servidor Gin con las rutas registradas.
	r := gin.Default()
	handler.RegistrarRutas(r, ph)

	fmt.Println("API corriendo en http://localhost:8080")
	fmt.Println("Endpoints disponibles:")
	fmt.Println("  GET    /api/v1/polizas")
	fmt.Println("  GET    /api/v1/polizas/:id")
	fmt.Println("  POST   /api/v1/polizas")
	fmt.Println("  PUT    /api/v1/polizas/:id/desactivar")
	fmt.Println("  PUT    /api/v1/polizas/:id/descuento")

	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

func ejecutarMigraciones() {
	// migrations.FS viene del paquete infrastructure/migrations que embebe los archivos SQL.
	fuente, err := iofs.New(migrations.FS, ".")
	if err != nil {
		log.Fatal("error leyendo migraciones:", err)
	}
	m, err := migrate.NewWithSourceInstance("iofs", fuente, dsn)
	if err != nil {
		log.Fatal("error configurando migraciones:", err)
	}
	defer m.Close()

	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		log.Fatal("error aplicando migraciones:", err)
	}
	fmt.Println("migraciones aplicadas")
}
