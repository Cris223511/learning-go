// Handlers HTTP de clientes-svc.
// El endpoint /clientes/:id/polizas muestra la comunicación entre microservicios:
// clientes-svc llama a polizas-svc por HTTP para obtener las pólizas del cliente.

package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type crearClienteRequest struct {
	ID     string `json:"id"     binding:"required"`
	Nombre string `json:"nombre" binding:"required"`
	Email  string `json:"email"  binding:"required,email"`
	DNI    string `json:"dni"    binding:"required"`
}

func registrarRutas(r *gin.Engine, repo ClienteRepo, polizasClient *PolizasClient) {
	r.GET("/health", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{"status": "ok", "servicio": "clientes-svc"})
	})

	v1 := r.Group("/api/v1")

	v1.GET("/clientes", func(c *gin.Context) {
		clientes, err := repo.BuscarTodos()
		if err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, clientes)
	})

	v1.GET("/clientes/:id", func(c *gin.Context) {
		cliente, err := repo.BuscarPorID(c.Param("id"))
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusOK, cliente)
	})

	v1.POST("/clientes", func(c *gin.Context) {
		var req crearClienteRequest
		if err := c.ShouldBindJSON(&req); err != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
			return
		}
		cliente := &Cliente{ID: req.ID, Nombre: req.Nombre, Email: req.Email, DNI: req.DNI}
		if err := cliente.Validar(); err != nil {
			c.JSON(http.StatusUnprocessableEntity, gin.H{"error": err.Error()})
			return
		}
		if err := repo.Guardar(cliente); err != nil {
			c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		c.JSON(http.StatusCreated, cliente)
	})

	// Este endpoint es la comunicación entre microservicios.
	// clientes-svc llama a polizas-svc via HTTP para obtener las pólizas del cliente.
	// Si polizas-svc está caído, devuelve el cliente con pólizas vacías en lugar de fallar.
	v1.GET("/clientes/:id/polizas", func(c *gin.Context) {
		id := c.Param("id")
		cliente, err := repo.BuscarPorID(id)
		if err != nil {
			c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
			return
		}

		// Llamada HTTP a polizas-svc. Si falla, responde gracefully con array vacío.
		polizas, err := polizasClient.ObtenerPorCliente(id)
		if err != nil {
			c.JSON(http.StatusOK, gin.H{
				"cliente": cliente,
				"polizas": []PolizaResumen{},
				"warning": "no se pudo contactar a polizas-svc: " + err.Error(),
			})
			return
		}

		c.JSON(http.StatusOK, gin.H{
			"cliente": cliente,
			"polizas": polizas,
		})
	})
}
