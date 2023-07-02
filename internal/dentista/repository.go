package dentista

import (
	"errors"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
	"github.com/MarcelaRamg/FinalBack3.git/pkg/dentistaPkg"
)

type DentistaRepository interface {

	// GetByID busca un dentista por su id
	GetByID(id int) (domain.Dentista, error)
	// Create agrega un nuevo dentista
	Create(p domain.Dentista) (domain.Dentista, error)
	// Update actualiza un dentista
	Update(id int, p domain.Dentista) (domain.Dentista, error)
	// Delete elimina un dentista
	Delete(id int) error
}

type dentistaRepository struct {
	storage dentistaPkg.DentistaInterface
}

// NewRepository crea un nuevo repositorio
func NewRepository(storage dentistaPkg.DentistaInterface) DentistaRepository {
	return &dentistaRepository{storage}
}

func (r *dentistaRepository) GetByID(id int) (domain.Dentista, error) {
	product, err := r.storage.Read(id)
	if err != nil {
		return domain.Dentista{}, errors.New("dentista not found")
	}
	return product, nil

}

func (r *dentistaRepository) Create(p domain.Dentista) (domain.Dentista, error) {
	if r.storage.Exists(p.Matricula) {
		return domain.Dentista{}, errors.New("matricula already exists")
	}
	err := r.storage.Create(p)
	if err != nil {
		return domain.Dentista{}, errors.New("error creating dentista")
	}
	return p, nil
}

func (r *dentistaRepository) Delete(id int) error {
	err := r.storage.Delete(id)
	if err != nil {
		return err
	}
	return nil
}

func (r *dentistaRepository) Update(id int, p domain.Dentista) (domain.Dentista, error) {
	if r.storage.Exists(p.Matricula) {
		return domain.Dentista{}, errors.New("matricula already exists")
	}
	err := r.storage.Update(p)
	if err != nil {
		return domain.Dentista{}, errors.New("error updating dentista")
	}
	return p, nil
}
