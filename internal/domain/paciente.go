package domain

type Paciente struct {
	ID              int    `json:"id"`
	Name            string `json:"name" binding:"required"`
	Apellido        string `json:"apellido" binding:"required"`
	Dni             string `json:"dni" binding:"required"`
	FechaNacimiento string `json:"fechaNacimiento" binding:"required"`
}
