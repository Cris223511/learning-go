// Acá veremos el manejo de errores básico en Go. La clave es que error es una
// interface, no una excepción: los errores se retornan, no se lanzan.

package main

import (
	"errors"
	"fmt"
)

// El patrón estándar: retornar el resultado y un error.
// Si todo salió bien, el error es nil.
func crearPoliza(id string, prima float64) (string, error) {
	// errors.New crea un error simple con un mensaje fijo.
	if id == "" {
		return "", errors.New("el id no puede estar vacío")
	}
	// fmt.Errorf crea un error con formato, igual que Sprintf pero retorna error.
	if prima <= 0 {
		return "", fmt.Errorf("prima inválida: %.2f, debe ser mayor a cero", prima)
	}
	return fmt.Sprintf("póliza %s creada con prima S/%.2f", id, prima), nil
}

func validarCliente(nombre string, edad int) error {
	if nombre == "" {
		// Retornamos nuestro tipo propio definido en errores.go
		return ErrorValidacion{Campo: "nombre", Mensaje: "no puede estar vacío"}
	}
	if edad < 18 {
		return ErrorValidacion{Campo: "edad", Mensaje: fmt.Sprintf("%d no alcanza la mayoría de edad", edad)}
	}
	return nil
}

func procesarReclamo(polizaID string) error {
	if polizaID == "POL-999" {
		return ErrorNegocio{Codigo: "POL_VENCIDA", Detalle: "la póliza no tiene cobertura vigente"}
	}
	return nil
}

func main() {
	fmt.Println("== errors.New y fmt.Errorf ==")
	msg, err := crearPoliza("POL-001", 450.50)
	if err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println(msg)
	}

	_, err = crearPoliza("", 100)
	if err != nil {
		fmt.Println("error:", err)
	}

	_, err = crearPoliza("POL-002", -50)
	if err != nil {
		fmt.Println("error:", err)
	}

	fmt.Println("\n== Tipo de error propio ==")
	if err := validarCliente("", 25); err != nil {
		fmt.Println("error:", err)
	}
	if err := validarCliente("Christopher", 16); err != nil {
		fmt.Println("error:", err)
	}
	if err := validarCliente("Christopher", 24); err == nil {
		fmt.Println("cliente válido")
	}

	fmt.Println("\n== Error de negocio ==")
	if err := procesarReclamo("POL-999"); err != nil {
		fmt.Println("error:", err)
		// errors.As extrae el tipo concreto del error para acceder a sus campos.
		// Lo veremos a fondo en el ejemplo de error wrapping.
		var negocio ErrorNegocio
		if errors.As(err, &negocio) {
			fmt.Println("código:", negocio.Codigo)
		}
	}
}
