package paciente

import (
	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
)

type Service interface {
	// GetByID busca un producto por su id
	GetByID(id int) (domain.Paciente, error)
	// Create agrega un nuevo producto
	Create(p domain.Paciente) (domain.Paciente, error)
	// Delete elimina un producto
	Delete(id int) error
	// Update actualiza un producto
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
	if u.Name == "" {
		p.Name = u.Name
	}
	if u.Apellido == "" {
		p.Apellido = u.Apellido
	}
	if u.Dni == "" {
		p.Dni = u.Dni
	}
	if u.FechaNacimiento == "" {
		p.FechaNacimiento = u.FechaNacimiento
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
