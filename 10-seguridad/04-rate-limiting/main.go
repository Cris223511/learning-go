// Acá veremos rate limiting: limitar cuántos requests puede hacer una IP en un período.
// Protege la API de abuso, scraping y ataques de fuerza bruta.
// Correr: go run ./10-seguridad/04-rate-limiting  → probar en Postman en localhost:8080
//
// Rutas:
//   GET /api/v1/polizas   → limitado a 5 requests por minuto por IP
//   GET /health           → sin rate limiting

package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	// 5 requests por minuto por IP.
	limiter := NuevoRateLimiter(5, time.Minute)

	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	api := r.Group("/api/v1")
	api.Use(RateLimitMiddleware(limiter))
	{
		api.GET("/polizas", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{
				"data": []string{"POL-001", "POL-002", "POL-003"},
			})
		})

		api.GET("/cotizar", func(c *gin.Context) {
			c.JSON(http.StatusOK, gin.H{"prima": 450.50})
		})
	}

	fmt.Println("Rate limiting: máx 5 requests/minuto por IP")
	r.Run(":8080")
}

// RateLimitMiddleware corta el request con 429 Too Many Requests si la IP excedió el límite.
func RateLimitMiddleware(rl *RateLimiter) gin.HandlerFunc {
	return func(c *gin.Context) {
		// ClientIP() retorna la IP real del cliente, considerando headers de proxy.
		ip := c.ClientIP()
		if !rl.Permitir(ip) {
			c.Header("Retry-After", "60")
			c.AbortWithStatusJSON(http.StatusTooManyRequests, gin.H{
				"error":       "demasiadas solicitudes",
				"retry_after": "60 segundos",
			})
			return
		}
		c.Next()
	}
}
