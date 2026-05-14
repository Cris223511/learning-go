// Expone los archivos SQL embebidos para que cmd/main.go los use al arrancar.
// El //go:embed solo puede referenciar archivos en el mismo directorio o subdirectorios,
// por eso el embed vive acá y no en cmd/main.go.

package migrations

import "embed"

//go:embed *.sql
var FS embed.FS
