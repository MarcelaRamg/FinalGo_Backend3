package paciente

import (
	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
)

type Service interface {
	// GetByID busca un paciente por su id
	GetByID(id int) (domain.Paciente, error)
	// GetByDni busca un paciente por DNI
	GetByDni(dni float64) (domain.Paciente, error)
	// Create agrega un nuevo paciente
	Create(p domain.Paciente) (domain.Paciente, error)
	// Delete elimina un paciente
	Delete(id int) error
	// Update actualiza un paciente
	Update(id int, p domain.Paciente) (domain.Paciente, error)
}

type service struct {
	r Repository
}

// NewService crea un nuevo servicio
func NewService(r Repository) Service {
	return &service{r}
}

func (s *service) GetByID(id int) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) GetByDni(dni float64) (domain.Paciente, error) {
	p, err := s.r.GetByDni(dni)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) Create(p domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}
func (s *service) Update(id int, u domain.Paciente) (domain.Paciente, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Paciente{}, err
	}
	if u.Nombre != "" {
		p.Nombre = u.Nombre
	}
	if u.Apellido != "" {
		p.Apellido = u.Apellido
	}
	if u.Dni != 0 {
		p.Dni = u.Dni
	}
	if u.FechaAlta != "" {
		p.FechaAlta = u.FechaAlta
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Paciente{}, err
	}
	return p, nil
}

func (s *service) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
