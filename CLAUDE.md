# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

Archivo de contexto para Claude. Cuando trabajes en este repositorio, lee primero este archivo.

## Sobre el usuario

- **Nombre:** Christopher
- **Experiencia previa:** 3+ años como Full Stack Developer. Stack fuerte en Java/Spring Boot, Angular, TypeScript, Node.js, PHP, PostgreSQL, MySQL, Firebase, Docker.
- **Nivel real:** No es un programador nuevo. Ya domina los conceptos universales de programación. Lo que está aprendiendo es **Go específicamente**, no programación.

## Propósito del repositorio

Ruta de aprendizaje progresiva de Go, desde lo más básico hasta el nivel productivo. El objetivo final es dominar APIs REST en Go sobre PostgreSQL con clean architecture, microservicios, seguridad básica y testing.

## Roadmap de 14 módulos

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
| 10 | Seguridad básica en APIs: JWT, CORS, validación, rate limiting | Listo |
| 11 | PostgreSQL: database/sql, pgx, migraciones, transacciones | Listo |
| 12 | Clean Architecture: capas, inversión de dependencias, DTOs | Listo |
| 13 | Microservicios: comunicación HTTP, Docker, observabilidad | Listo |
| 14 | Git profesional: conventional commits, branching, PRs, CI | Listo |

## Convenciones del repositorio

- Cada módulo es una carpeta `XX-tema/`.
- Módulos 01-11: cada ejemplo es una subcarpeta con un `main.go` ejecutable independiente.
- Módulo 12: proyecto completo `polizas-api/` con clean architecture real.
- Módulo 13: dos microservicios (`polizas-svc/`, `clientes-svc/`) con `docker-compose.yml`.
- Se ejecutan con: `go run ./<modulo>/<ejemplo>`.
- El repo es un solo módulo Go declarado en `go.mod` con path `github.com/cpillihuaman/learning-go`.

## Reglas de estilo de código (CRÍTICAS)

### Comentarios

Estas reglas aplican SOLO dentro de este repositorio (`learning-go`). En cualquier otro proyecto del usuario, código senior limpio sin comentarios.

1. **Intro obligatoria al inicio de cada archivo `.go`**: 1 a 3 líneas con el formato `// Acá veremos...` antes del `package`. Resume qué se aprende en ese archivo.
2. **Comentarios internos breves**: máximo 1 a 3 líneas por bloque. Tono natural de profesor explicando, no de manual técnico.
3. **PROHIBIDO usar ecuaciones en comentarios**: nada de `// const = valor fijo`. Escribir frases completas.
4. **Comentar SOLO lo específico de Go**: `iota`, `:=`, `rune` vs `byte`, `any`, `type switch`, `for-range`, zero values, conversión explícita obligatoria, `range` sobre map sin orden garantizado, `break` con label, etc.
5. **NO comentar obviedades**: qué es una variable, qué es un if/else, qué es un bucle. El usuario ya lo sabe.
6. **Tono profesor pero conciso**: frases naturales tipo "Acá lo nuevo es...", "Esto sirve para...", "Fíjate que...". Sin párrafos didácticos largos.

### Otras reglas

- **NUNCA usar el guion largo "—" (em dash)**. Reemplazar por coma, dos puntos o reescribir.
- Código formateado con `gofmt`.
- Nombres en `camelCase` para privados, `PascalCase` para exportados.
- Errores se retornan, no se lanzan.

## Cómo trabajar con el usuario

- El usuario prefiere acción directa, no demasiada conversación previa.
- Cuando hay decisiones importantes, presentar opciones concretas con AskUserQuestion.
- El usuario es directo, si algo no le gusta lo dice fuerte. Tomarlo como feedback útil, no defenderse.

## Comandos esenciales

```bash
go version              # versión instalada
go run ./ruta/al/main   # ejecuta sin compilar binario
go build ./...          # compila todo el módulo
go vet ./...            # análisis estático
go fmt ./...            # formatea código
go mod tidy             # limpia dependencias
go test ./...           # corre todos los tests
go test -cover ./...    # tests con cobertura
```

## Base de datos (módulos 11-13)

PostgreSQL vía Docker. Nombre de la base: `aprendizago`.

```bash
docker run -d -p 5432:5432 -e POSTGRES_PASSWORD=postgres --name pg-aprendizago postgres:16
docker exec -it pg-aprendizago psql -U postgres -c "CREATE DATABASE aprendizago;"
docker start pg-aprendizago   # si ya existe el contenedor y apagaste la PC
```

DSN estándar: `postgres://postgres:postgres@localhost:5432/aprendizago?sslmode=disable`
