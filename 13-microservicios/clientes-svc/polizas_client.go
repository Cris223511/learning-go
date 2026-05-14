// Cliente HTTP para comunicarse con polizas-svc.
// Esto es la comunicación entre microservicios: clientes-svc llama a polizas-svc
// por HTTP para obtener las pólizas de un cliente. Nunca accede a su BD directamente.

package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type PolizasClient struct {
	baseURL    string
	httpClient *http.Client
}

func nuevoPolizasClient(baseURL string) *PolizasClient {
	return &PolizasClient{
		baseURL: baseURL,
		// Timeout obligatorio: sin él, una llamada colgada bloquea la goroutine para siempre.
		httpClient: &http.Client{Timeout: 5 * time.Second},
	}
}

// ObtenerPorCliente llama al endpoint GET /api/v1/polizas?cliente_id=X de polizas-svc.
// Si polizas-svc está caído, retorna error y clientes-svc lo maneja gracefully.
func (c *PolizasClient) ObtenerPorCliente(clienteID string) ([]PolizaResumen, error) {
	url := fmt.Sprintf("%s/api/v1/polizas?cliente_id=%s", c.baseURL, clienteID)

	resp, err := c.httpClient.Get(url)
	if err != nil {
		// Error de red: polizas-svc no está disponible.
		return nil, fmt.Errorf("no se pudo contactar a polizas-svc: %w", err)
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("polizas-svc respondió con status %d", resp.StatusCode)
	}

	var polizas []PolizaResumen
	if err := json.NewDecoder(resp.Body).Decode(&polizas); err != nil {
		return nil, fmt.Errorf("error decodificando respuesta de polizas-svc: %w", err)
	}
	return polizas, nil
}
