package domain

type Paciente struct {
	ID        int     `json:"id"`
	Nombre    string  `json:"nombre" binding:"required"`
	Apellido  string  `json:"apellido" binding:"required"`
	Dni       float64 `json:"dni" binding:"required"`
	FechaAlta string  `json:"fechaAlta" binding:"required"`
}
