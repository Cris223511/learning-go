// Acá viven los handlers HTTP: traducen requests de Gin a llamadas al usecase
// y traducen las respuestas del usecase a JSON. Nada más.
// Si mañana cambias de Gin a otro framework, solo tocas esta capa.

package handler

import (
	"net/http"
	"strconv"

	"github.com/cpillihuaman/learning-go/12-clean-architecture/polizas-api/internal/usecase"
	"github.com/gin-gonic/gin"
)

// Los DTOs (structs de request/response) viven acá porque son parte de la capa HTTP.
// Son distintos de las entidades del dominio: el cliente de la API no tiene por qué
// ver exactamente igual la entidad que usa internamente el negocio.

type crearPolizaRequest struct {
	ClienteID string  `json:"cliente_id" binding:"required"`
	Tipo      string  `json:"tipo"       binding:"required"`
	Prima     float64 `json:"prima"      binding:"required,gt=0"`
}

type descuentoRequest struct {
	Porcentaje float64 `json:"porcentaje" binding:"required,gt=0,lte=50"`
}

type errorResponse struct {
	Error string `json:"error"`
}

type PolizaHandler struct {
	uc *usecase.PolizaUseCase
}

func NuevoPolizaHandler(uc *usecase.PolizaUseCase) *PolizaHandler {
	return &PolizaHandler{uc: uc}
}

func (h *PolizaHandler) Listar(c *gin.Context) {
	polizas, err := h.uc.Listar()
	if err != nil {
		c.JSON(http.StatusInternalServerError, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, polizas)
}

func (h *PolizaHandler) ObtenerPorID(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: "id debe ser un número"})
		return
	}
	poliza, err := h.uc.ObtenerPorID(id)
	if err != nil {
		c.JSON(http.StatusNotFound, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, poliza)
}

func (h *PolizaHandler) Crear(c *gin.Context) {
	var req crearPolizaRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}
	poliza, err := h.uc.Crear(req.ClienteID, req.Tipo, req.Prima)
	if err != nil {
		c.JSON(http.StatusUnprocessableEntity, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusCreated, poliza)
}

func (h *PolizaHandler) Desactivar(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: "id inválido"})
		return
	}
	poliza, err := h.uc.Desactivar(id)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, poliza)
}

func (h *PolizaHandler) AplicarDescuento(c *gin.Context) {
	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: "id inválido"})
		return
	}
	var req descuentoRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}
	poliza, err := h.uc.AplicarDescuento(id, req.Porcentaje)
	if err != nil {
		c.JSON(http.StatusBadRequest, errorResponse{Error: err.Error()})
		return
	}
	c.JSON(http.StatusOK, poliza)
}
