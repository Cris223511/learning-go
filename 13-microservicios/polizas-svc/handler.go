// Handlers HTTP de polizas-svc. Traduce requests Gin → lógica de negocio → JSON.
// Incluye un endpoint /health que docker-compose usa para saber si el servicio está listo.

package main

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

type crearRequest struct {
	ClienteID string  `json:"cliente_id" binding:"required"`
	Tipo      string  `json:"tipo"       binding:"required"`
	Prima     float64 `json:"prima"      binding:"required,gt=0"`
}

func registrarRutas(r *gin.Engine, repo PolizaRepo) {
	// Health check: docker-compose lo usa para esperar a que el servicio esté listo
	// antes de arrancar los servicios que dependen de este.
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "servicio": "polizas-svc"})
	})

	v1 := r.Group("/api/v1")
	v1.GET("/polizas", func(c *gin.Context) {
		// Si viene ?cliente_id=CLI-001 filtra por cliente, si no devuelve todos.
		// clientes-svc usa este query param para pedir las pólizas de un cliente específico.
		clienteID := c.Query("cliente_id")
		var (
			polizas []*Poliza
			err     error
		)
		if clienteID != "" {
			polizas, err = repo.BuscarPorCliente(clienteID)
		} else {
			polizas, err = repo.BuscarTodos()
		}
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, polizas)
	})

	v1.GET("/polizas/:id", func(c *gin.Context) {
		id, err := strconv.Atoi(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": "id inválido"})
			return
		}
		p, err := repo.BuscarPorID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	})

	v1.POST("/polizas", func(c *gin.Context) {
		var req crearRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		p := &Poliza{ClienteID: req.ClienteID, Tipo: req.Tipo, Prima: req.Prima, Activa: true}
		if err := p.Validar(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		creada, err := repo.Guardar(p)
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, creada)
	})

	v1.PUT("/polizas/:id/desactivar", func(c *gin.Context) {
		id, _ := strconv.Atoi(c.Param("id"))
		p, err := repo.BuscarPorID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		if err := p.Desactivar(); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		if err := repo.Actualizar(p); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, p)
	})
}
