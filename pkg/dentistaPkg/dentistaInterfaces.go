package dentistaPkg

import "github.com/MarcelaRamg/FinalBack3.git/internal/domain"

type DentistaInterface interface {
	// Read devuelve un dentista por su id
	Read(id int) (domain.Dentista, error)
	// Read devuelve un dentista por su id
	ReadByMatricula(matricula string) (domain.Dentista, error)
	// Create agrega un nuevo dentista
	Create(dentista domain.Dentista) error
	// Update actualiza un dentista
	Update(dentista domain.Dentista) error
	//Delete elimina un dentista
	Delete(id int) error
	// Exists verifica si un dentista existe
	Exists(codeValue string) bool
}
