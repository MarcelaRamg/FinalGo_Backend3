package domain

type Paciente struct {
	ID        int    `json:"id"`
	Nombre    string `json:"name" binding:"required"`
	Apellido  string `json:"apellido" binding:"required"`
	Dni       string `json:"dni" binding:"required"`
	FechaAlta string `json:"fechaNacimiento" binding:"required"`
}
