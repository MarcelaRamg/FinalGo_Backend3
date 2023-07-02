package domain

type Turno struct {
	ID          int    `json:"id"`
	FechaHora   string `json:"fechaHora" binding:"required"`
	PacienteID  string `json:"paciente" binding:"required"`
	DentistaID  string `json:"odontologo" binding:"required"`
	Descripcion string `json:"descripcion" binding:"required"`
}

type TurnoByMatriculaAndDni struct {
	FechaHora   string  `json:"fechaHora" binding:"required"`
	Descripcion string  `json:"descripcion" binding:"required"`
	Dni         float64 `json:"dni_paciente" binding:"required"`
	Matricula   string  `json:"matricula_odontologo" binding:"required"`
}
