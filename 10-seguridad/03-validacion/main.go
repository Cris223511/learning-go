// Acá veremos validación de inputs en Gin con go-playground/validator.
// La validación se declara en los tags de la struct y Gin la ejecuta automáticamente.
// Correr: go run ./10-seguridad/03-validacion  → probar en Postman en localhost:8080
//
// Rutas:
//   POST /api/v1/clientes   body: {"nombre":"Ana","email":"ana@mail.com","dni":"12345678","edad":25}
//   POST /api/v1/polizas    body: {"cliente_id":"CLI-001","tipo":"SOAT","prima":120,"duracion":12}

package main

import (
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.POST("/clientes", crearCliente)
	v1.POST("/polizas", crearPoliza)

	r.Run(":8080")
}

func crearCliente(c *gin.Context) {
	var req CrearClienteRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatearErrores(err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "cliente creado",
		"datos":   req,
	})
}

func crearPoliza(c *gin.Context) {
	var req CrearPolizaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, formatearErrores(err))
		return
	}
	c.JSON(http.StatusCreated, gin.H{
		"mensaje": "póliza creada",
		"datos":   req,
	})
}

// formatearErrores convierte los errores de validación en un mapa campo→mensaje
// legible para el cliente de la API.
func formatearErrores(err error) RespuestaError {
	// ValidationErrors es un slice de FieldError, uno por cada campo que falló.
	// Intentamos el cast al tipo concreto de errores de validación.
	errs, ok := err.(validator.ValidationErrors)
	if !ok {
		return RespuestaError{Error: err.Error()}
	}
	campos := make(map[string]string)
	for _, e := range errs {
		// Field() retorna el nombre del campo. Tag() retorna la regla que falló.
		campos[strings.ToLower(e.Field())] = mensajeValidacion(e)
	}
	return RespuestaError{Error: "datos inválidos", Campos: campos}
}

func mensajeValidacion(e validator.FieldError) string {
	switch e.Tag() {
	case "required":
		return "campo obligatorio"
	case "email":
		return "formato de email inválido"
	case "min":
		return "valor mínimo: " + e.Param()
	case "max":
		return "valor máximo: " + e.Param()
	case "len":
		return "longitud exacta requerida: " + e.Param()
	case "numeric":
		return "solo se permiten dígitos"
	case "gte":
		return "debe ser mayor o igual a " + e.Param()
	case "lte":
		return "debe ser menor o igual a " + e.Param()
	case "oneof":
		return "valores permitidos: " + e.Param()
	default:
		return "validación fallida: " + e.Tag()
	}
}
