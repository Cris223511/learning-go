// Modelos con validaciones usando go-playground/validator a través de Gin.
// Los tags `binding:` activan la validación automática en ShouldBindJSON.

package main

// CrearClienteRequest valida todos los campos de entrada del request.
type CrearClienteRequest struct {
	Nombre   string `json:"nombre"   binding:"required,min=2,max=100"`
	Email    string `json:"email"    binding:"required,email"`
	DNI      string `json:"dni"      binding:"required,len=8,numeric"`
	Edad     int    `json:"edad"     binding:"required,gte=18,lte=120"`
	Telefono string `json:"telefono" binding:"omitempty,min=9,max=15"`
}

// CrearPolizaRequest valida los datos de una nueva póliza.
type CrearPolizaRequest struct {
	ClienteID string  `json:"cliente_id" binding:"required"`
	// oneof valida que el valor sea uno de los listados exactamente.
	Tipo      string  `json:"tipo"       binding:"required,oneof=SOAT Vida Vehicular Hogar"`
	Prima     float64 `json:"prima"      binding:"required,gt=0"`
	// Duración en meses, entre 1 y 120 (10 años).
	Duracion  int     `json:"duracion"   binding:"required,min=1,max=120"`
}

type RespuestaError struct {
	Error   string            `json:"error"`
	Campos  map[string]string `json:"campos,omitempty"`
}
