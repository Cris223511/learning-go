// Modelos con tags de JSON y validación. Los tags `json:` controlan la serialización,
// `binding:` activa la validación de Gin al hacer ShouldBindJSON.

package main

type CrearPolizaRequest struct {
	// binding:"required" hace que ShouldBindJSON falle si el campo no viene o está vacío.
	Tipo   string  `json:"tipo"   binding:"required"`
	Prima  float64 `json:"prima"  binding:"required,gt=0"`
	Activa bool    `json:"activa"`
}

type ActualizarPolizaRequest struct {
	Prima  float64 `json:"prima"  binding:"omitempty,gt=0"`
	Activa *bool   `json:"activa" binding:"omitempty"` // puntero para distinguir false de "no enviado"
}

type Poliza struct {
	ID     string  `json:"id"`
	Tipo   string  `json:"tipo"`
	Prima  float64 `json:"prima"`
	Activa bool    `json:"activa"`
}

// RespuestaError es el formato estándar de error de la API.
type RespuestaError struct {
	Error   string `json:"error"`
	Detalle string `json:"detalle,omitempty"` // omitempty: no aparece en el JSON si está vacío
}

// RespuestaPaginada envuelve una lista con metadata de paginación.
type RespuestaPaginada struct {
	Data  []Poliza `json:"data"`
	Total int      `json:"total"`
	Page  int      `json:"page"`
}
