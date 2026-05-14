# Learning Go

Ruta de aprendizaje progresiva de Go, desde lo más básico hasta construir APIs REST en producción con PostgreSQL, clean architecture y microservicios.

## Stack

| Capa | Tecnología |
|------|-----------|
| Lenguaje | Go 1.22+ |
| Base de datos | PostgreSQL |
| HTTP framework | Gin |
| Testing | testify |
| Contenedores | Docker |
| Control de versiones | Git + GitHub Actions |

## Roadmap

| Módulo | Tema | Estado |
|--------|------|--------|
| 01 | Fundamentos: sintaxis, variables, tipos, control de flujo | Listo |
| 02 | Colecciones: arrays, slices, maps, strings | Listo |
| 03 | Funciones: parámetros, retornos múltiples, closures, defer | Listo |
| 04 | Structs y métodos: composición, punteros, receivers | Listo |
| 05 | Interfaces y polimorfismo: design by contract, type assertion | Listo |
| 06 | Manejo de errores: error wrapping, panic, recover, sentinel errors | Listo |
| 07 | Concurrencia: goroutines, channels, select, sync, context | Listo |
| 08 | Testing: unit tests, table tests, mocks, benchmarks, coverage | Listo |
| 09 | HTTP server: net/http, Gin, routing, middleware, JSON | Listo |
| 10 | Seguridad en APIs: JWT, CORS, validación, rate limiting | Listo |
| 11 | PostgreSQL: database/sql, pgx, migraciones, transacciones | Listo |
| 12 | Clean Architecture: capas, inversión de dependencias, DTOs | Listo |
| 13 | Microservicios: comunicación HTTP, Docker, observabilidad | Listo |
| 14 | Git profesional: conventional commits, branching, PRs, CI | Listo |

## Estructura

```
learning-go/
├── go.mod
├── 01-fundamentos/
│   └── 01-hello/main.go
├── 02-colecciones/
│   └── 01-arrays/main.go
├── ...
├── 12-clean-architecture/
│   └── polizas-api/        ← proyecto completo con capas
└── 13-microservicios/
    ├── polizas-svc/
    ├── clientes-svc/
    └── docker-compose.yml
```

Los módulos 01 al 11 tienen subcarpetas independientes, cada una con un `main.go` ejecutable:

```bash
go run ./01-fundamentos/01-hello
go run ./09-http-server/02-routing
```

Los módulos 12 y 13 son proyectos reales con estructura de producción.

## Comandos esenciales

```bash
go run ./ruta/al/ejemplo   # ejecutar un ejemplo
go test ./08-testing/... -v  # correr los tests del módulo 08
go test -bench=. ./08-testing/04-benchmarks  # correr benchmarks
go vet ./...               # análisis estático
go fmt ./...               # formatear código
go mod tidy                # limpiar dependencias
```
