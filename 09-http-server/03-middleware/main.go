// Acá veremos middleware en Gin: funciones que se ejecutan antes o después de los handlers.
// Se usan para logging, autenticación, recovery, CORS, rate limiting, etc.
// Correr: go run ./09-http-server/03-middleware  → probar en Postman en localhost:8080
//
// Rutas disponibles:
//   GET  /health                    → sin auth
//   GET  /api/v1/polizas            → requiere header X-API-Key: secreto-aprendizago
//   GET  /api/v1/panic              → demuestra el middleware de recovery

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type Poliza struct {
	ID    string  `json:"id"`
	Tipo  string  `json:"tipo"`
	Prima float64 `json:"prima"`
}

func main() {
	// gin.New() crea un router sin middlewares. Los agregamos manualmente.
	r := gin.New()

	// Use() agrega middlewares globales: se ejecutan en todas las rutas.
	r.Use(Recovery())
	r.Use(Logger())
	r.Use(RequestID())

	r.GET("/health", func(c *gin.Context) {
		reqID, _ := c.Get("requestID")
		c.JSON(http.StatusOK, gin.H{"status": "ok", "request_id": reqID})
	})

	// Group con middleware específico: solo las rutas de este grupo requieren API key.
	api := r.Group("/api/v1")
	api.Use(APIKeyAuth("secreto-aprendizago"))
	{
		api.GET("/polizas", func(c *gin.Context) {
			c.JSON(http.StatusOK, []Poliza{
				{ID: "POL-001", Tipo: "SOAT", Prima: 120},
				{ID: "POL-002", Tipo: "Vida", Prima: 1500},
			})
		})

		// Esta ruta paniquea a propósito para demostrar que Recovery lo captura.
		api.GET("/panic", func(c *gin.Context) {
			panic("este panic fue capturado por el middleware Recovery")
		})
	}

	r.Run(":8080")
}
