package turno

import "github.com/MarcelaRamg/FinalBack3.git/internal/domain"

type TurnoService interface {
	GetByID(id int) (domain.Turno, error)
	GetByDni(dni float64) ([]domain.Turno, error)
	Create(p domain.Turno) (domain.Turno, error)
	Delete(id int) error
	Update(id int, p domain.Turno) (domain.Turno, error)
	/* CreateByPacienteAndDentistaID(paciente_dni string, dentista_matricula string) ([]domain.Turno, error)
	GetByPacienteDNI(paciente_dni string) ([]domain.Turno, error) */
}

type turnoService struct {
	r TurnoRepository
}

// NewService crea un nuevo servicio
func NewService(r TurnoRepository) TurnoService {
	return &turnoService{r}
}

func (s *turnoService) GetByID(id int) (domain.Turno, error) {
	turno, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (s *turnoService) GetByDni(dni float64) ([]domain.Turno, error) {
	turno, err := s.r.GetByDni(dni)
	if err != nil {
		return []domain.Turno{}, err
	}
	return turno, nil
}

func (s *turnoService) Create(p domain.Turno) (domain.Turno, error) {
	turno, err := s.r.Create(p)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}
func (s *turnoService) Update(id int, u domain.Turno) (domain.Turno, error) {
	turno, err := s.r.GetByID(id)
	if err != nil {
		return domain.Turno{}, err
	}
	if u.PacienteID != "" {
		turno.PacienteID = u.PacienteID
	}
	if u.DentistaID != "" {
		turno.DentistaID = u.DentistaID
	}
	if u.FechaHora != "" {
		turno.FechaHora = u.FechaHora
	}
	if u.Descripcion != "" {
		turno.Descripcion = u.Descripcion
	}

	turno, err = s.r.Update(id, turno)
	if err != nil {
		return domain.Turno{}, err
	}
	return turno, nil
}

func (s *turnoService) Delete(id int) error {
	err := s.r.Delete(id)
	if err != nil {
		return err
	}
	return nil
}
