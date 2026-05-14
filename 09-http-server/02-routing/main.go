// Acá veremos Gin: el framework HTTP más usado en Go en producción.
// Gin simplifica routing, params, query strings y middleware respecto a net/http.
// Correr: go run ./09-http-server/02-routing  → probar en Postman en localhost:8080

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
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

func main() {
	// gin.Default() incluye Logger y Recovery (manejo de panics) por defecto.
	r := gin.Default()

	r.GET("/health", func(c *gin.Context) {
		// c.JSON serializa a JSON y setea el Content-Type automáticamente.
		c.JSON(http.StatusOK, gin.H{"status": "ok"})
	})

	// Grupo de rutas: todas empiezan con /api/v1
	v1 := r.Group("/api/v1")
	{
		v1.GET("/polizas", listarPolizas)
		v1.GET("/polizas/:id", obtenerPoliza) // :id es un path param
		v1.POST("/polizas", crearPoliza)
		v1.DELETE("/polizas/:id", eliminarPoliza)
	}

	// Rutas con query params: GET /api/v1/buscar?tipo=SOAT&activa=true
	v1.GET("/buscar", buscarPolizas)

	r.Run(":8080")
}

func listarPolizas(c *gin.Context) {
	c.JSON(http.StatusOK, polizas)
}

func obtenerPoliza(c *gin.Context) {
	// c.Param recupera el path param :id
	id := c.Param("id")
	for _, p := range polizas {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "póliza no encontrada", "id": id})
}

func crearPoliza(c *gin.Context) {
	var nueva Poliza
	// ShouldBindJSON deserializa el body JSON a la struct. Retorna error si falla.
	if err := c.ShouldBindJSON(&nueva); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	polizas = append(polizas, nueva)
	c.JSON(http.StatusCreated, nueva)
}

func eliminarPoliza(c *gin.Context) {
	id := c.Param("id")
	for i, p := range polizas {
		if p.ID == id {
			polizas = append(polizas[:i], polizas[i+1:]...)
			c.JSON(http.StatusOK, gin.H{"mensaje": "eliminada", "id": id})
			return
		}
	}
	c.JSON(http.StatusNotFound, gin.H{"error": "no encontrada"})
}

func buscarPolizas(c *gin.Context) {
	// c.Query recupera query params. c.DefaultQuery pone un valor por defecto si no viene.
	tipo := c.Query("tipo")
	soloActivas := c.DefaultQuery("activa", "false") == "true"

	var resultado []Poliza
	for _, p := range polizas {
		if tipo != "" && p.Tipo != tipo {
			continue
		}
		if soloActivas && !p.Activa {
			continue
		}
		resultado = append(resultado, p)
	}
	c.JSON(http.StatusOK, resultado)
}
