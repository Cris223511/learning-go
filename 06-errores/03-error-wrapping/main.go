// Acá veremos error wrapping: cómo agregar contexto a un error sin perder
// el error original, y cómo inspeccionarlo con errors.Is y errors.As.

package main

import (
	"errors"
	"fmt"
)

var ErrNoEncontrado = errors.New("no encontrado")

type ErrorDB struct {
	Operacion string
	Tabla     string
	Causa     error
}

// Para que errors.Is y errors.As funcionen con un tipo propio y sus errores envueltos,
// el tipo debe implementar Unwrap() error. Así Go puede desempaquetar la cadena.
func (e ErrorDB) Error() string {
	return fmt.Sprintf("DB [%s en %s]: %v", e.Operacion, e.Tabla, e.Causa)
}

func (e ErrorDB) Unwrap() error {
	return e.Causa
}

func buscarEnDB(id string) error {
	if id == "" {
		return fmt.Errorf("buscarEnDB: %w", ErrNoEncontrado)
	}
	return nil
}

func buscarPoliza(id string) error {
	if err := buscarEnDB(id); err != nil {
		// %w envuelve el error: lo incluye dentro del nuevo y permite que
		// errors.Is/As lo encuentren al desempaquetar la cadena.
		return fmt.Errorf("buscarPoliza(%q): %w", id, err)
	}
	return nil
}

func cargarPolizaDesdeRepo(id string) error {
	if id == "POL-999" {
		causa := fmt.Errorf("registro eliminado: %w", ErrNoEncontrado)
		return ErrorDB{Operacion: "SELECT", Tabla: "polizas", Causa: causa}
	}
	return nil
}

func main() {
	fmt.Println("== Wrapping con fmt.Errorf y %w ==")
	err := buscarPoliza("")
	if err != nil {
		fmt.Println("error completo:", err)
		// errors.Is recorre toda la cadena de errores envueltos buscando el sentinel.
		fmt.Println("¿es ErrNoEncontrado?", errors.Is(err, ErrNoEncontrado))
	}

	fmt.Println("\n== Cadena de wrapping ==")
	err2 := buscarPoliza("")
	// Puedes ver todos los errores de la cadena desempaquetando manualmente.
	for e := err2; e != nil; e = errors.Unwrap(e) {
		fmt.Println(" ->", e)
	}

	fmt.Println("\n== errors.As: extraer el tipo concreto ==")
	err3 := cargarPolizaDesdeRepo("POL-999")
	if err3 != nil {
		fmt.Println("error:", err3)

		// errors.As recorre la cadena buscando un error del tipo indicado
		// y lo asigna a la variable para que puedas acceder a sus campos.
		var dbErr ErrorDB
		if errors.As(err3, &dbErr) {
			fmt.Printf("operación fallida: %s en tabla %s\n", dbErr.Operacion, dbErr.Tabla)
		}
		// errors.Is también funciona porque ErrorDB implementa Unwrap().
		fmt.Println("¿contiene ErrNoEncontrado?", errors.Is(err3, ErrNoEncontrado))
	}

	fmt.Println("\n== Sin wrapping (Errorf con verb v, no w) ==")
	errSinWrap := fmt.Errorf("algo salió mal: %v", ErrNoEncontrado)
	// %v no envuelve: el error original queda como texto, no como valor.
	// errors.Is no puede encontrarlo porque no hay cadena que recorrer.
	fmt.Println("¿es ErrNoEncontrado?:", errors.Is(errSinWrap, ErrNoEncontrado)) // false
}
