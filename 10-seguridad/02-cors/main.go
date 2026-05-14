// Acá veremos CORS en Gin con gin-contrib/cors.
// CORS define qué orígenes (dominios) pueden hacer requests a la API desde un navegador.
// Correr: go run ./10-seguridad/02-cors  → probar en Postman en localhost:8080
//
// Rutas:
//   GET  /api/v1/polizas   → permitido desde orígenes configurados
//   GET  /api/v1/publico   → CORS permisivo (cualquier origen)

package main

import (
	"net/http"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// Configuración de CORS para producción: solo los orígenes que deben acceder.
	configProd := cors.Config{
		// Lista explícita de orígenes permitidos. Nunca uses * en producción con cookies/auth.
		AllowOrigins: []string{
			"https://app.miapi.com",
			"https://admin.miapi.com",
		},
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Authorization", "Content-Type", "X-Request-ID"},
		ExposeHeaders:    []string{"X-Request-ID"},
		AllowCredentials: true,
		// MaxAge: cuánto tiempo el navegador cachea la respuesta del preflight OPTIONS.
		MaxAge: 12 * time.Hour,
	}

	api := r.Group("/api/v1")
	api.Use(cors.New(configProd))
	{
		api.GET("/polizas", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"data": []string{"POL-001", "POL-002"}})
		})
	}

	// Rutas públicas con CORS permisivo (ej: documentación, health check públicos).
	publico := r.Group("/api/v1")
	publico.Use(cors.Default()) // cors.Default() permite cualquier origen con GET/POST/PUT
	{
		publico.GET("/publico", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"mensaje": "endpoint público, cualquier origen permitido"})
		})
	}

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	r.Run(":8080")
}
