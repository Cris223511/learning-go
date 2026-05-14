// Acá veremos JWT en Go con Gin: login que genera el token y middleware que lo valida.
// Correr: go run ./10-seguridad/01-jwt  → probar en Postman en localhost:8080
//
// Flujo:
//   1. POST /login  body: {"email":"admin@correo.com","password":"1234"}
//   2. Copiar el token de la respuesta
//   3. GET /api/v1/polizas  header: Authorization: Bearer <token>

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type LoginRequest struct {
	Email    string `json:"email"    binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

// Usuarios hardcodeados solo para el ejemplo. En producción van en BD con bcrypt.
var usuarios = map[string]string{
	"admin@correo.com": "1234",
	"user@correo.com":  "abcd",
}

func main() {
	r := gin.Default()

	r.POST("/login", login)

	api := r.Group("/api/v1")
	// El middleware AuthJWT protege todas las rutas de este grupo.
	api.Use(AuthJWT())
	{
		api.GET("/polizas", func(c *gin.Context) {
			// Los claims del token fueron inyectados en el contexto por el middleware.
			claims, _ := c.Get("claims")
			c.JSON(http.StatusOK, gin.H{
				"mensaje":  "acceso autorizado",
				"usuario":  claims.(*Claims).Email,
				"rol":      claims.(*Claims).Rol,
				"polizas":  []string{"POL-001", "POL-002"},
			})
		})

		api.GET("/perfil", func(c *gin.Context) {
			claims := c.MustGet("claims").(*Claims)
			c.JSON(http.StatusOK, gin.H{
				"usuario_id": claims.UsuarioID,
				"email":      claims.Email,
				"rol":        claims.Rol,
			})
		})
	}

	r.Run(":8080")
}

func login(c *gin.Context) {
	var req LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	pass, existe := usuarios[req.Email]
	if !existe || pass != req.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "credenciales inválidas"})
		return
	}
	token, err := GenerarToken("USR-001", req.Email, "admin")
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "no se pudo generar el token"})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token, "tipo": "Bearer"})
}

// AuthJWT es el middleware que valida el token en cada request.
func AuthJWT() gin.HandlerFunc {
	return func(c *gin.Context) {
		header := c.GetHeader("Authorization")
		if len(header) < 8 || header[:7] != "Bearer " {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token requerido"})
			return
		}
		tokenStr := header[7:]
		claims, err := ValidarToken(tokenStr)
		if err != nil {
			c.AbortWithStatusJSON(http.StatusUnauthorized, gin.H{"error": "token inválido: " + err.Error()})
			return
		}
		// Inyectamos los claims en el contexto para que los handlers los lean.
		c.Set("claims", claims)
		c.Next()
	}
}
