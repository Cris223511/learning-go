// Acá veremos CI (Integración Continua) con GitHub Actions.
// El archivo .github/workflows/ci.yml (en la raíz del repo) corre automáticamente
// en cada push y en cada PR, verificando que el código compile, pase los tests
// y no tenga problemas de análisis estático.
//
// Correr: go run ./14-git-profesional/04-ci

package main

import "fmt"

func main() {
	fmt.Println("=== CI con GitHub Actions: learning-go ===\n")

	fmt.Println("El archivo .github/workflows/ci.yml en la raíz del repo")
	fmt.Println("define el pipeline que GitHub ejecuta automáticamente.\n")

	fmt.Println("¿QUÉ HACE EL PIPELINE?")
	pasos := []string{
		"Checkout del código (actions/checkout)",
		"Instalar Go 1.25 (actions/setup-go)",
		"Descargar dependencias (go mod download)",
		"Análisis estático (go vet ./...)",
		"Correr todos los tests con coverage (go test -cover ./...)",
		"Compilar todo el módulo (go build ./...)",
	}
	for i, p := range pasos {
		fmt.Printf("  %d. %s\n", i+1, p)
	}

	fmt.Println("\n¿CUÁNDO CORRE?")
	fmt.Println("  • En cada push a main")
	fmt.Println("  • En cada Pull Request hacia main")
	fmt.Println("  • Puedes verlo en la pestaña Actions de tu repo en GitHub\n")

	fmt.Println("¿CÓMO SE VE EN UN PR?")
	fmt.Println("  GitHub muestra ✓ o ✗ en el PR antes de poder mergearlo.")
	fmt.Println("  Si el CI falla, no se puede mergear (si está configurado como required).\n")

	fmt.Println("COMANDOS PARA VER EL ESTADO DEL CI")
	fmt.Println("  gh pr checks               ver checks del PR actual")
	fmt.Println("  gh run list                ver ejecuciones recientes del workflow")
	fmt.Println("  gh run view                ver detalles de la última ejecución")
	fmt.Println("  gh run view --log          ver los logs completos")
}
