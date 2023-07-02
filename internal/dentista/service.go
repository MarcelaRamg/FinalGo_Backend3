package dentista

import "github.com/MarcelaRamg/FinalBack3.git/internal/domain"

type DentistaService interface {
	GetByID(id int) (domain.Dentista, error)
	GetByMatricula(matricula string) (domain.Dentista, error)
	Create(p domain.Dentista) (domain.Dentista, error)
	Delete(id int) error
	Update(id int, p domain.Dentista) (domain.Dentista, error)
}

type dentistaService struct {
	r DentistaRepository
}

// NewService crea un nuevo servicio
func NewService(r DentistaRepository) DentistaService {
	return &dentistaService{r}
}

func (s *dentistaService) GetByID(id int) (domain.Dentista, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

func (s *dentistaService) GetByMatricula(matricula string) (domain.Dentista, error) {
	p, err := s.r.GetByMatricula(matricula)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

func (s *dentistaService) Create(p domain.Dentista) (domain.Dentista, error) {
	p, err := s.r.Create(p)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}
func (s *dentistaService) Update(id int, u domain.Dentista) (domain.Dentista, error) {
	p, err := s.r.GetByID(id)
	if err != nil {
		return domain.Dentista{}, err
	}
	if u.Apellido != "" {
		p.Apellido = u.Apellido
	}
	if u.Nombre != "" {
		p.Nombre = u.Nombre
	}
	if u.Matricula != "" {
		p.Matricula = u.Matricula
	}
	p, err = s.r.Update(id, p)
	if err != nil {
		return domain.Dentista{}, err
	}
	return p, nil
}

func (s *dentistaService) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
