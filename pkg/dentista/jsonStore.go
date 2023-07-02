package dentista

import (
	"encoding/json"
	"errors"
	"os"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
)

type jsonStore struct {
	pathToFile string
}

// loadProducts carga los productos desde un archivo json
func (s *jsonStore) loadDentistas() ([]domain.Dentista, error) {
	var dentista []domain.Dentista
	file, err := os.ReadFile(s.pathToFile)
	if err != nil {
		return nil, err
	}
	err = json.Unmarshal([]byte(file), &dentista)
	if err != nil {
		return nil, err
	}
	return dentista, nil
}

// saveDentista guarda los dentistas en un archivo json
func (s *jsonStore) saveDentistas(dentista []domain.Dentista) error {
	bytes, err := json.Marshal(dentista)
	if err != nil {
		return err
	}
	return os.WriteFile(s.pathToFile, bytes, 0644)
}

// NewJsonStore crea un nuevo store de dentistas
func NewJsonStore(path string) DentistaInterface {
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	return &jsonStore{
		pathToFile: path,
	}
}

func (s *jsonStore) GetAll() ([]domain.Dentista, error) {
	return nil, nil
}

func (s *jsonStore) Read(id int) (domain.Dentista, error) {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return domain.Dentista{}, err
	}
	for _, dentista := range dentistas {
		if dentista.ID == id {
			return dentista, nil
		}
	}
	return domain.Dentista{}, errors.New("dentista not found")
}

func (s *jsonStore) Create(dentista domain.Dentista) error {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return err
	}
	dentista.ID = len(dentistas) + 1
	dentistas = append(dentistas, dentista)
	return s.saveDentistas(dentistas)
}

func (s *jsonStore) Update(dentista domain.Dentista) error {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return err
	}
	for i, p := range dentistas {
		if p.ID == dentista.ID {
			dentistas[i] = dentista
			return s.saveDentistas(dentistas)
		}
	}
	return errors.New("dentista not found")
}

func (s *jsonStore) Delete(id int) error {
	dentistas, err := s.loadDentistas()
	if err != nil {
		return err
	}
	for i, p := range dentistas {
		if p.ID == id {
			dentistas = append(dentistas[:i], dentistas[i+1:]...)
			return s.saveDentistas(dentistas)
		}
	}
	return errors.New("dentista not found")
}

func (s *jsonStore) Exists(codeValue string) bool {
	products, err := s.loadDentistas()
	if err != nil {
		return false
	}
	for _, p := range products {
		if p.Matricula == codeValue {
			return true
		}
	}
	return false
}
