// Acá veremos la estrategia de ramas para este proyecto.
// Usamos un flujo simple basado en trunk (main) con ramas de feature cortas.
//
// Correr: go run ./14-git-profesional/02-branching

package main

import "fmt"

func main() {
	fmt.Println("=== Estrategia de ramas: learning-go ===\n")

	fmt.Println("RAMA PRINCIPAL")
	fmt.Println("  main → siempre estable, solo recibe PRs aprobados\n")

	fmt.Println("CONVENCIÓN DE NOMBRES")
	fmt.Println("  feat/nombre-corto     nueva funcionalidad")
	fmt.Println("  fix/nombre-del-bug    corrección de bug")
	fmt.Println("  docs/que-documentas   cambios de documentación")
	fmt.Println("  refactor/que-cambias  refactor sin nuevas features\n")

	fmt.Println("FLUJO DE TRABAJO")
	pasos := []struct {
		cmd  string
		desc string
	}{
		{"git checkout main", "Siempre partir desde main actualizado"},
		{"git pull origin main", "Traer los últimos cambios"},
		{"git checkout -b feat/mi-feature", "Crear la rama de trabajo"},
		{"# ... hacer cambios ...", ""},
		{"git add archivo.go", "Agregar solo los archivos relevantes (nunca git add .)"},
		{`git commit -m "feat(modulo): descripción clara"`, "Commitear con conventional commits"},
		{"git push origin feat/mi-feature", "Subir la rama a GitHub"},
		{"gh pr create --title \"feat: mi feature\" --body \"...\"", "Crear el PR con gh CLI"},
	}
	for i, p := range pasos {
		if p.desc != "" {
			fmt.Printf("  %d. %-55s # %s\n", i+1, p.cmd, p.desc)
		} else {
			fmt.Printf("     %s\n", p.cmd)
		}
	}

	fmt.Println("\nREGLAS")
	reglas := []string{
		"Nunca hacer push directo a main",
		"Las ramas deben ser cortas: idealmente menos de 2 días de trabajo",
		"Un PR = una cosa. No mezclar features con refactors",
		"Resolver conflictos con rebase, no con merge commit: git rebase main",
		"Borrar la rama después de que el PR se mergea",
	}
	for _, r := range reglas {
		fmt.Printf("  • %s\n", r)
	}

	fmt.Println("\nCOMANDOS ÚTILES")
	fmt.Println("  git log --oneline --graph --all   ver el historial visual")
	fmt.Println("  git stash                          guardar cambios temporalmente")
	fmt.Println("  git stash pop                      recuperar cambios guardados")
	fmt.Println("  git rebase -i HEAD~3               editar los últimos 3 commits")
	fmt.Println("  gh pr list                         ver PRs abiertos")
	fmt.Println("  gh pr view --web                   abrir el PR actual en el navegador")
}
