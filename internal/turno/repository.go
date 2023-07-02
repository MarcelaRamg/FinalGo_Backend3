package turno

import (
	"errors"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/turnoPkg"
)

type TurnoRepository interface {

	// GetByID busca un turno por su id
	GetByID(id int) (domain.Turno, error)
	// Create agrega un nuevo turno
	Create(p domain.Turno) (domain.Turno, error)
	// Update actualiza un turno
	Update(id int, p domain.Turno) (domain.Turno, error)
	// Delete elimina un turno
	Delete(id int) error
}

type turnoRepository struct {
	storage turnoPkg.TurnoInterface
}

// NewRepository crea un nuevo repositorio

func NewTurnoRepository(storage turnoPkg.TurnoInterface) TurnoRepository {
	return &turnoRepository{storage}
}

func (r *turnoRepository) GetByID(id int) (domain.Turno, error) {
	turno, err := r.storage.Read(id)
	if err != nil {
		return domain.Turno{}, errors.New("turno not found")
	}
	return turno, nil
}

func (r *turnoRepository) Create(p domain.Turno) (domain.Turno, error) {
	if r.storage.Exists(p.ID) {
		return domain.Turno{}, errors.New("turno already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return domain.Turno{}, errors.New("error creating turno")
	}
	return p, nil
}

func (r *turnoRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *turnoRepository) Update(id int, p domain.Turno) (domain.Turno, error) {
	if r.storage.Exists(p.ID) {
		return domain.Turno{}, errors.New("turno already exists")
	}
	err := r.storage.Update(p)
	if err != nil {
		return domain.Turno{}, errors.New("error updating turno")
	}
	return p, nil
}
