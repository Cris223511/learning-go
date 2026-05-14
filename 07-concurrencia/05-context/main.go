// Acá veremos context: la forma estándar de Go para propagar cancelación,
// timeouts y valores a través de una cadena de llamadas de funciones.

package main

import (
	"context"
	"fmt"
	"time"
)

// Por convención, ctx siempre es el primer parámetro de cualquier función
// que haga trabajo cancelable (llamadas a BD, HTTP, procesamiento largo).
func buscarPoliza(ctx context.Context, id string) (string, error) {
	select {
	case <-time.After(80 * time.Millisecond): // simula consulta a BD
		return fmt.Sprintf("póliza %s encontrada", id), nil
	case <-ctx.Done():
		// ctx.Err() dice por qué se canceló: context.Canceled o context.DeadlineExceeded.
		return "", fmt.Errorf("búsqueda cancelada: %w", ctx.Err())
	}
}

func procesarSiniestro(ctx context.Context, id string) error {
	fmt.Println("  iniciando procesamiento de", id)

	// Pasas el mismo contexto a todas las llamadas internas.
	// Si el contexto padre se cancela, todos los hijos se enteran.
	poliza, err := buscarPoliza(ctx, "POL-001")
	if err != nil {
		return fmt.Errorf("procesarSiniestro: %w", err)
	}
	fmt.Println(" ", poliza)
	return nil
}

// context.WithValue agrega un valor al contexto. Se usa para datos transversales
// como el ID del request o el usuario autenticado, no para parámetros de negocio.
type claveCtx string

const claveUsuario claveCtx = "usuario"

func manejarRequest(ctx context.Context) {
	usuario, ok := ctx.Value(claveUsuario).(string)
	if !ok {
		fmt.Println("  usuario no encontrado en contexto")
		return
	}
	fmt.Println("  request manejado por:", usuario)
}

func main() {
	// context.Background() es el contexto raíz. Nunca se cancela.
	// Es el punto de partida para todos los demás contextos.
	fmt.Println("== WithCancel: cancelación manual ==")
	ctx, cancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(40 * time.Millisecond)
		cancel() // cancela el contexto después de 40ms
	}()

	_, err := buscarPoliza(ctx, "POL-001")
	if err != nil {
		fmt.Println("error:", err)
	}

	// WithTimeout: cancela automáticamente después del tiempo indicado.
	// Siempre llama cancel() con defer para liberar recursos aunque termine antes.
	fmt.Println("\n== WithTimeout: cancelación automática ==")
	ctx2, cancel2 := context.WithTimeout(context.Background(), 200*time.Millisecond)
	defer cancel2()

	if err := procesarSiniestro(ctx2, "SIN-007"); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("  siniestro procesado con éxito")
	}

	// Timeout demasiado corto: el contexto expira antes que la operación.
	fmt.Println("\n== WithTimeout muy corto ==")
	ctx3, cancel3 := context.WithTimeout(context.Background(), 30*time.Millisecond)
	defer cancel3()

	if err := procesarSiniestro(ctx3, "SIN-008"); err != nil {
		fmt.Println("error:", err)
	}

	// WithDeadline: igual que Timeout pero con una hora absoluta.
	fmt.Println("\n== WithDeadline ==")
	deadline := time.Now().Add(150 * time.Millisecond)
	ctx4, cancel4 := context.WithDeadline(context.Background(), deadline)
	defer cancel4()
	fmt.Println("deadline en:", time.Until(deadline).Round(time.Millisecond))

	if _, err := buscarPoliza(ctx4, "POL-002"); err != nil {
		fmt.Println("error:", err)
	} else {
		fmt.Println("  póliza encontrada antes del deadline")
	}

	// WithValue: agrega metadata al contexto. La clave debe ser un tipo propio
	// (no string ni int) para evitar colisiones entre paquetes.
	fmt.Println("\n== WithValue: datos del request ==")
	ctx5 := context.WithValue(context.Background(), claveUsuario, "christopher@correo.com")
	manejarRequest(ctx5)
	manejarRequest(context.Background()) // sin valor
}
