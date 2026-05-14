// Acá veremos el servidor HTTP de la librería estándar: net/http.
// Antes de usar Gin conviene entender qué hay debajo.
// Correr: go run ./09-http-server/01-net-http  → probar en Postman en localhost:8080

package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Poliza struct {
	ID     string  `json:"id"`
	Tipo   string  `json:"tipo"`
	Prima  float64 `json:"prima"`
	Activa bool    `json:"activa"`
}

var polizas = []Poliza{
	{ID: "POL-001", Tipo: "SOAT", Prima: 120.00, Activa: true},
	{ID: "POL-002", Tipo: "Vida", Prima: 1500.00, Activa: true},
	{ID: "POL-003", Tipo: "Vehicular", Prima: 890.75, Activa: false},
}

// Un handler es cualquier función con la firma (ResponseWriter, *Request).
// ResponseWriter escribe la respuesta, Request tiene todo lo del request entrante.
func listarPolizas(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "método no permitido", http.StatusMethodNotAllowed)
		return
	}
	// Le decimos al cliente que la respuesta es JSON.
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(polizas)
}

func healthCheck(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, `{"status":"ok","servicio":"polizas-api"}`)
}

func main() {
	// ServeMux es el router de net/http. Asocia rutas con handlers.
	mux := http.NewServeMux()
	mux.HandleFunc("/health", healthCheck)
	mux.HandleFunc("/polizas", listarPolizas)

	fmt.Println("servidor en http://localhost:8080")
	fmt.Println("rutas disponibles:")
	fmt.Println("  GET /health")
	fmt.Println("  GET /polizas")

	// ListenAndServe bloquea. Ctrl+C para detener.
	if err := http.ListenAndServe(":8080", mux); err != nil {
		log.Fatal(err)
	}
}
