// Acá veremos el manejo completo de JSON en Gin: binding, validación, tags y respuestas.
// Correr: go run ./09-http-server/04-json  → probar en Postman en localhost:8080
//
// Rutas:
//   GET  /api/v1/polizas              → lista paginada (?page=1&limit=10)
//   GET  /api/v1/polizas/:id          → una póliza
//   POST /api/v1/polizas              → crear (body JSON requerido)
//   PUT  /api/v1/polizas/:id          → actualizar parcialmente

package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

var db = []Poliza{
	{ID: "POL-001", Tipo: "SOAT", Prima: 120.00, Activa: true},
	{ID: "POL-002", Tipo: "Vida", Prima: 1500.00, Activa: true},
	{ID: "POL-003", Tipo: "Vehicular", Prima: 890.75, Activa: false},
}

var contador = 4

func main() {
	r := gin.Default()
	v1 := r.Group("/api/v1")

	v1.GET("/polizas", listar)
	v1.GET("/polizas/:id", obtener)
	v1.POST("/polizas", crear)
	v1.PUT("/polizas/:id", actualizar)

	r.Run(":8080")
}

func listar(c *gin.Context) {
	page, _ := strconv.Atoi(c.DefaultQuery("page", "1"))
	limit, _ := strconv.Atoi(c.DefaultQuery("limit", "10"))

	inicio := (page - 1) * limit
	fin := inicio + limit
	if inicio >= len(db) {
		c.JSON(http.StatusOK, RespuestaPaginada{Data: []Poliza{}, Total: len(db), Page: page})
		return
	}
	if fin > len(db) {
		fin = len(db)
	}
	c.JSON(http.StatusOK, RespuestaPaginada{Data: db[inicio:fin], Total: len(db), Page: page})
}

func obtener(c *gin.Context) {
	id := c.Param("id")
	for _, p := range db {
		if p.ID == id {
			c.JSON(http.StatusOK, p)
			return
		}
	}
	c.JSON(http.StatusNotFound, RespuestaError{Error: "no encontrada", Detalle: id})
}

func crear(c *gin.Context) {
	var req CrearPolizaRequest
	// ShouldBindJSON valida los tags binding:"required" y binding:"gt=0".
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, RespuestaError{
			Error:   "datos inválidos",
			Detalle: err.Error(),
		})
		return
	}
	nueva := Poliza{
		ID:     fmt.Sprintf("POL-%03d", contador),
		Tipo:   req.Tipo,
		Prima:  req.Prima,
		Activa: req.Activa,
	}
	contador++
	db = append(db, nueva)
	c.JSON(http.StatusCreated, nueva)
}

func actualizar(c *gin.Context) {
	id := c.Param("id")
	var req ActualizarPolizaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, RespuestaError{Error: "datos inválidos", Detalle: err.Error()})
		return
	}
	for i, p := range db {
		if p.ID == id {
			if req.Prima > 0 {
				db[i].Prima = req.Prima
			}
			// Puntero bool: si viene en el JSON (aunque sea false), lo aplicamos.
			if req.Activa != nil {
				db[i].Activa = *req.Activa
			}
			c.JSON(http.StatusOK, db[i])
			return
		}
	}
	c.JSON(http.StatusNotFound, RespuestaError{Error: "no encontrada"})
}
