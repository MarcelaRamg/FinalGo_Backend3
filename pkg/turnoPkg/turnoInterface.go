package turnoPkg

import "github.com/MarcelaRamg/FinalBack3.git/internal/domain"

type TurnoInterface interface {
	Read(id int) (domain.Turno, error)
	Create(turno domain.Turno) error
	Update(turno domain.Turno) error
	Delete(id int) error
	Exists(codeValue int) bool
}
