// Acá veremos cómo crear Pull Requests profesionales con gh CLI.
// Un buen PR tiene: título claro, descripción del cambio, plan de testing
// y referencia al ticket o issue si existe.
//
// Correr: go run ./14-git-profesional/03-prs

package main

import (
	"fmt"
	"os/exec"
)

func main() {
	fmt.Println("=== Pull Requests con gh CLI ===\n")

	fmt.Println("ESTRUCTURA DE UN BUEN PR")
	fmt.Println(`
  Título:  feat(polizas): agregar endpoint de descuento por volumen

  Cuerpo:
  ## ¿Qué hace este PR?
  Agrega PUT /api/v1/polizas/:id/descuento que aplica un porcentaje
  de descuento a la prima de una póliza activa.

  ## ¿Por qué?
  Requisito de negocio para clientes corporativos con más de 10 pólizas.

  ## Cómo testear
  - [ ] POST /api/v1/polizas para crear una póliza
  - [ ] PUT /api/v1/polizas/1/descuento body: {"porcentaje": 15}
  - [ ] Verificar que prima_final = prima_base * 0.85
  - [ ] Intentar descuento > 50% → debe rechazar con 400

  ## Checklist
  - [ ] Tests unitarios agregados
  - [ ] go vet sin errores
  - [ ] go test ./... pasa`)

	fmt.Println("\nCOMANDOS PARA CREAR EL PR")
	fmt.Println(`
  # Forma básica
  gh pr create --title "feat(polizas): agregar descuento" --body "descripción"

  # Forma interactiva (abre editor)
  gh pr create

  # Con reviewers y labels
  gh pr create --title "..." --body "..." --reviewer usuario --label "enhancement"

  # Abre el PR en el navegador al crearlo
  gh pr create --web`)

	fmt.Println("\nCOMANDOS ÚTILES DESPUÉS DE CREAR EL PR")
	fmt.Println("  gh pr list                  ver todos los PRs abiertos")
	fmt.Println("  gh pr view                  ver el PR de la rama actual")
	fmt.Println("  gh pr view --web            abrirlo en el navegador")
	fmt.Println("  gh pr checks                ver el estado del CI")
	fmt.Println("  gh pr merge --squash        mergear aplastando los commits en uno")
	fmt.Println("  gh pr close                 cerrar sin mergear")

	// Verificar si gh está instalado y mostrar el estado actual.
	fmt.Println("\n=== Estado actual del repo ===")
	if out, err := exec.Command("gh", "auth", "status").CombinedOutput(); err != nil {
		fmt.Println("gh CLI no autenticado. Correr: gh auth login")
	} else {
		fmt.Println(string(out))
	}
}
