// Middlewares personalizados. En Gin un middleware es un HandlerFunc que llama
// c.Next() para pasar al siguiente handler, o c.Abort() para cortar la cadena.

package main

import (
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

// Logger registra método, ruta, status y latencia de cada request.
func Logger() gin.HandlerFunc {
	return func(c *gin.Context) {
		inicio := time.Now()
		// c.Next() ejecuta los handlers que siguen en la cadena.
		c.Next()
		latencia := time.Since(inicio)
		log.Printf("[%d] %s %s  %v",
			c.Writer.Status(),
			c.Request.Method,
			c.Request.URL.Path,
			latencia,
		)
	}
}

// RequestID agrega un ID único a cada request y lo expone en el header de respuesta.
// Útil para trazabilidad en logs y en sistemas distribuidos.
func RequestID() gin.HandlerFunc {
	contador := 0
	return func(c *gin.Context) {
		contador++
		id := fmt.Sprintf("req-%06d", contador)
		c.Set("requestID", id)
		c.Header("X-Request-ID", id)
		c.Next()
	}
}

// APIKeyAuth verifica que el request traiga una API key válida en el header.
// Si no la trae, corta la cadena con Abort y retorna 401.
func APIKeyAuth(claveValida string) gin.HandlerFunc {
	return func(c *gin.Context) {
		clave := c.GetHeader("X-API-Key")
		if clave != claveValida {
			// Abort detiene la cadena: los handlers siguientes no se ejecutan.
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{
				"error": "API key inválida o ausente",
			})
			return
		}
		c.Next()
	}
}

// Recovery captura cualquier panic y devuelve 500 en lugar de matar el proceso.
// Gin ya incluye uno en gin.Default(), acá lo mostramos para entender cómo funciona.
func Recovery() gin.HandlerFunc {
	return func(c *gin.Context) {
		defer func() {
			if r := recover(); r != nil {
				log.Printf("panic recuperado: %v", r)
				c.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"error": "error interno del servidor",
				})
			}
		}()
		c.Next()
	}
}
