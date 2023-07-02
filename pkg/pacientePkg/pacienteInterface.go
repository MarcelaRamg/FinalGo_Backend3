package pacientePkg

import "github.com/MarcelaRamg/FinalBack3.git/internal/domain"

type PacienteInterface interface {
	Read(id int) (domain.Paciente, error)
	ReadByDni(dni float64) (domain.Paciente, error)
	Create(paciente domain.Paciente) error
	Update(paciente domain.Paciente) error
	Delete(id int) error
	Exists(codeValue string) bool
}
