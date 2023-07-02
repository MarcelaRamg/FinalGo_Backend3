package domain

type Turno struct {
	ID          int    `json:"id"`
	FechaHora   string `json:"fechaHora" binding:"required"`
	PacienteID  string `json:"paciente" binding:"required"`
	DentistaID  string `json:"odontologo" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}
