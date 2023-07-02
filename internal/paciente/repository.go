package paciente

import (
	"errors"
	"fmt"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/pacientePkg"
)

type Repository interface {
	GetByID(id int) (domain.Paciente, error)
	GetByDni(dni float64) (domain.Paciente, error)
	Create(p domain.Paciente) (domain.Paciente, error)
	Update(id int, p domain.Paciente) (domain.Paciente, error)
	Delete(id int) error
}

type repository struct {
	storage pacientePkg.PacienteInterface
}

func NewRepository(storage pacientePkg.PacienteInterface) Repository {
	return &repository{storage}
}

func (r *repository) GetByID(id int) (domain.Paciente, error) {
	paciente, err := r.storage.Read(id)
	if err != nil {
		return domain.Paciente{}, errors.New("Paciente not found")
	}
	return paciente, nil

}
func (r *repository) GetByDni(dni float64) (domain.Paciente, error) {
	paciente, err := r.storage.ReadByDni(dni)
	if err != nil {
		return domain.Paciente{}, errors.New("Paciente not found")
	}
	return paciente, nil

}

func (r *repository) Create(p domain.Paciente) (domain.Paciente, error) {
	err := r.storage.Create(p)
	if err != nil {
		return domain.Paciente{}, errors.New(fmt.Sprintf("error creating Paciente: %s", err.Error()))
	}
	return p, nil
}

func (r *repository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *repository) Update(id int, p domain.Paciente) (domain.Paciente, error) {
	err := r.storage.Update(p)
	if err != nil {
		return domain.Paciente{}, errors.New("error .. updating Paciente")
	}
	return p, nil
}
