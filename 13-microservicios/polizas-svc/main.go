// polizas-svc: microservicio responsable de gestionar pólizas de seguro.
// Expone una API REST en el puerto 8081.
// Solo este servicio puede leer y escribir en la tabla "polizas".
//
// Correr solo (sin Docker): go run ./13-microservicios/polizas-svc
// Correr con Docker:        docker-compose up   (desde 13-microservicios/)

package main

import (
	"context"
	"embed"
	"fmt"
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
	// slog es el paquete de logging estructurado de la stdlib (Go 1.21+).
	// Los logs estructurados son clave en microservicios para que herramientas como
	// Datadog, Grafana o CloudWatch puedan filtrar y buscar por campos específicos.
	logger := slog.New(slog.NewJSONHandler(os.Stdout, nil))
	slog.SetDefault(logger)

	// Las variables de entorno permiten configurar el servicio sin recompilar.
	// En Docker, docker-compose las inyecta. Localmente usamos el valor por defecto.
	dsn := getEnv("DATABASE_URL", "postgres://postgres:postgres@localhost:5432/aprendizago?sslmode=disable")
	port := getEnv("PORT", "8081")

	ctx := context.Background()

	// Conectar al pool de PostgreSQL.
	pool, err := pgxpool.New(ctx, dsn)
	if err != nil {
		slog.Error("error conectando a PostgreSQL", "error", err)
		os.Exit(1)
	}
	defer pool.Close()

	if err := pool.Ping(ctx); err != nil {
		slog.Error("ping fallido a PostgreSQL", "error", err)
		os.Exit(1)
	}
	slog.Info("conectado a PostgreSQL")

	// Ejecutar migraciones al arrancar. Si la tabla ya existe, no hace nada.
	if err := migrar(dsn); err != nil {
		slog.Error("error en migraciones", "error", err)
		os.Exit(1)
	}

	// Dependency injection: repo → handler.
	repo := nuevoRepo(pool)

	r := gin.New()
	// gin.Recovery() captura panics y devuelve 500 en lugar de matar el proceso.
	r.Use(gin.Recovery())
	// Middleware de logging: registra cada request con método, ruta y status.
	r.Use(loggerMiddleware())

	registrarRutas(r, repo)

	slog.Info("polizas-svc iniciado", "puerto", port)
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

// loggerMiddleware registra cada request con campos estructurados.
// En producción esto alimenta herramientas de observabilidad como Grafana o Datadog.
func loggerMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Next()
		slog.Info("request",
			"método", c.Request.Method,
			"ruta", c.Request.URL.Path,
			"status", c.Writer.Status(),
		)
	}
}

// getEnv lee una variable de entorno y devuelve un valor por defecto si no existe.
// Patrón estándar en microservicios Go para configuración.
func getEnv(clave, defecto string) string {
	if v := os.Getenv(clave); v != "" {
		return v
	}
	fmt.Printf("⚠ %s no definido, usando defecto: %s\n", clave, defecto)
	return defecto
}
