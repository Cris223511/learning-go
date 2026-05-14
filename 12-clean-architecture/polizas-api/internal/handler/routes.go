// Registro de rutas separado del handler para mantener cada archivo con una responsabilidad.
// Si la API crece con más recursos (clientes, pagos), cada uno tendría su propio handler
// y se registraría acá sin tocar los otros archivos.

package handler

import "github.com/gin-gonic/gin"

func RegistrarRutas(r *gin.Engine, ph *PolizaHandler) {
	v1 := r.Group("/api/v1")

	polizas := v1.Group("/polizas")
	{
		polizas.GET("", ph.Listar)
		polizas.GET("/:id", ph.ObtenerPorID)
		polizas.POST("", ph.Crear)
		polizas.PUT("/:id/desactivar", ph.Desactivar)
		polizas.PUT("/:id/descuento", ph.AplicarDescuento)
	}
}
