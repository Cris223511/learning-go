// clientes-svc: microservicio responsable de gestionar clientes.
// Expone una API REST en el puerto 8082.
// Para obtener las pólizas de un cliente llama a polizas-svc via HTTP,
// nunca accede directamente a la base de datos de ese servicio.
//
// Correr solo: go run ./13-microservicios/clientes-svc
// Correr con Docker: docker-compose up (desde 13-microservicios/)

package main

import (
	"context"
	"embed"
	"log/slog"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	"github.com/golang-migrate/migrate/v4/source/iofs"
	"github.com/jackc/pgx/v5/pgxpool"
)

//go:embed migrations/*.sql
var migraciones embed.FS

func main() {
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	dsn := getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/aprendizago?sslmode=disable")
	port := getEnv("PORT", "8082")
	// La URL de polizas-svc se configura por variable de entorno.
	// En Docker el hostname es "polizas-svc" (nombre del servicio en docker-compose).
	// Localmente es localhost:8081.
	polizasURL := getEnv("POLIZAS_SVC_URL", "http://localhost:8081")

	ctx := context.Background()

	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		slog.Error("error conectando a PostgreSQL", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		slog.Error("ping fallido", "error", err)
		os.Exit(1)
	}
	slog.Info("conectado a PostgreSQL")

	if err := migrar(dsn); err != nil {
		slog.Error("error en migraciones", "error", err)
		os.Exit(1)
	}

	// Inyección de dependencias: repo + cliente HTTP de polizas-svc → handler.
	repo := nuevoClienteRepo(pool)
	polizasClient := nuevoPolizasClient(polizasURL)

	r := gin.New()
	r.Use(gin.Recovery())
	r.Use(func(c *gin.Context) {
		c.Next()
		slog.Info("request",
			"método", c.Request.Method,
			"ruta", c.Request.URL.Path,
			"status", c.Writer.Status(),
		)
	})

	registrarRutas(r, repo, polizasClient)

	slog.Info("clientes-svc iniciado", "puerto", port, "polizas_svc", polizasURL)
	if err := r.Run(":" + port); err != nil {
		slog.Error("error iniciando servidor", "error", err)
		os.Exit(1)
	}
}

func migrar(dsn string) error {
	fuente, err := iofs.New(migraciones, "migrations")
	if err != nil {
		return err
	}
	m, err := migrate.NewWithSourceInstance("iofs", fuente, dsn)
	if err != nil {
		return err
	}
	defer m.Close()
	if err := m.Up(); err != nil && err != migrate.ErrNoChange {
		return err
	}
	slog.Info("migraciones aplicadas")
	return nil
}

func getEnv(clave, defecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	slog.Warn("variable de entorno no definida, usando defecto", "clave", clave, "defecto", defecto)
	return defecto
}
