// Acá veremos Conventional Commits: un estándar de mensajes de commit que hace
// el historial legible y permite generar changelogs automáticamente.
//
// Formato: tipo(scope): descripción
// Ejemplos:
//   feat(polizas): agregar endpoint de descuento
//   fix(auth): corregir expiración de token JWT
//   docs: actualizar README con instrucciones de Docker
//   test(usecase): agregar tests de tabla para CalcularPrima
//   chore: actualizar dependencias de gin y pgx
//
// Correr: go run ./14-git-profesional/01-conventional-commits "feat(api): agregar endpoint"
// Si no pasas argumento, valida una lista de ejemplos.

package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
)

// El regex valida la estructura exacta de un conventional commit.
// tipo(scope opcional): descripción obligatoria
var reCommit = regexp.MustCompile(
	`^(feat|fix|docs|style|refactor|test|chore|ci|perf|build|revert)(\(.+\))?: .{10,}$`,
)

// tipos describe qué significa cada prefijo. En un proyecto real esto guía al equipo.
var tipos = map[string]string{
	"feat":     "nueva funcionalidad para el usuario",
	"fix":      "corrección de bug",
	"docs":     "solo cambios de documentación",
	"style":    "formato, sin cambios de lógica (gofmt, espacios)",
	"refactor": "refactor que no agrega feature ni corrige bug",
	"test":     "agregar o corregir tests",
	"chore":    "tareas de mantenimiento (deps, config)",
	"ci":       "cambios en pipelines de CI/CD",
	"perf":     "mejora de rendimiento",
	"build":    "cambios en el sistema de build",
	"revert":   "revierte un commit anterior",
}

func validar(mensaje string) (bool, string) {
	mensaje = strings.TrimSpace(mensaje)
	if len(mensaje) == 0 {
		return false, "el mensaje no puede estar vacío"
	}
	if len(mensaje) > 72 {
		return false, fmt.Sprintf("muy largo (%d chars). El límite es 72 para que se vea bien en git log", len(mensaje))
	}
	if !reCommit.MatchString(mensaje) {
		return false, "no sigue el formato: tipo(scope): descripción mínimo 10 chars"
	}
	return true, "mensaje válido"
}

func main() {
	// Si se pasa un argumento, valida ese mensaje específico.
	if len(os.Args) > 1 {
		msg := strings.Join(os.Args[1:], " ")
		ok, resultado := validar(msg)
		simbolo := "✓"
		if !ok {
			simbolo = "✗"
		}
		fmt.Printf("%s %q → %s\n", simbolo, msg, resultado)
		if !ok {
			os.Exit(1)
		}
		return
	}

	// Sin argumento: muestra la referencia completa y valida ejemplos.
	fmt.Println("=== Conventional Commits: referencia rápida ===\n")
	fmt.Println("Formato: tipo(scope): descripción")
	fmt.Println("         └─ obligatorio   └─ opcional    └─ mínimo 10 chars, máx 72\n")

	fmt.Println("Tipos disponibles:")
	for tipo, desc := range tipos {
		fmt.Printf("  %-10s %s\n", tipo, desc)
	}

	fmt.Println("\n=== Validando ejemplos ===")
	ejemplos := []string{
		"feat(polizas): agregar endpoint para aplicar descuento por volumen",
		"fix(jwt): corregir validación de token expirado en middleware",
		"test(usecase): agregar table tests para CalcularPrima con casos borde",
		"chore: actualizar gin a v1.12 y pgx a v5.9",
		"docs: agregar instrucciones de Docker en módulo 13",
		// Inválidos a propósito para ver el error:
		"agregué cosas",
		"FEAT: nueva función",
		"fix: ok",
	}

	for _, e := range ejemplos {
		ok, resultado := validar(e)
		simbolo := "✓"
		if !ok {
			simbolo = "✗"
		}
		fmt.Printf("  %s %-65s → %s\n", simbolo, fmt.Sprintf("%q", e), resultado)
	}

	fmt.Println("\nUso: go run ./14-git-profesional/01-conventional-commits \"tu mensaje\"")
}
