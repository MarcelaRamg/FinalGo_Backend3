package turno

import (
	"errors"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/store"
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
	/* // Create runo por DNI del paciente y matricula del dentista
	CreateByPacienteAndDentistaID(paciente_dni string, dentista_matricula string) ([]domain.Turno, error)
	// GetByPacienteDNI
	GetByPacienteDNI(paciente_dni string) ([]domain.Turno, error)*/
}

type turnoRepository struct {
	storage store.DentistaInterface
}

// NewRepository crea un nuevo repositorio

func NewTurnoRepository(storage store.TurnoInterface) TurnoRepository {
	return &TurnoRepository{storage}
}

func (r *TurnoRepository) GetByID(id int) (domain.Turno, error) {
	turno, err := r.storage.Read(id)
	if err != nil {
		return domain.Turno{}, errors.New("turno not found")
	}
	return turno, nil
}

func (r *TurnoRepository) Create(p domain.Turno) (domain.Turno, error) {
	if r.storage.Exists(p.ID) {
		return domain.Turno{}, errors.New("turno already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return domain.Turno{}, errors.New("error creating turno")
	}
	return p, nil
}

func (r *TurnoRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *TurnoRepository) Update(id int, p domain.Turno) (domain.Turno, error) {
	if r.storage.Exists(p.ID) {
		return domain.Turno{}, errors.New("turno already exists")
	}
	err := r.storage.Update(p)
	if err != nil {
		return domain.Turno{}, errors.New("error updating turno")
	}
	return p, nil
}
