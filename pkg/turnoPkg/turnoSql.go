package turnoPkg

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type sqlTurno struct {
	db *sql.DB
}

func NewSQLTurno(db *sql.DB) TurnoInterface {
	if err := db.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlTurno{
		db: db,
	}
}

func (s *sqlTurno) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM Turno WHERE ID=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlTurno) Read(id int) (domain.Turno, error) {
	turno := domain.Turno{}

	rows, err := s.db.Query("SELECT * from Turno WHERE ID=?", id)
	if err != nil {
		return domain.Turno{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&turno.ID,
			&turno.FechaHora,
			&turno.PacienteID,
			&turno.DentistaID,
			&turno.Descripcion,
		)
		if err != nil {
			return domain.Turno{}, err
		}
	}
	return turno, nil
}

func (s *sqlTurno) Update(turno domain.Turno) error {
	fmt.Println("updating Turno")
	_, err := s.db.Exec(
		"UPDATE Turno SET PacienteID = ?, FechaHora = ?, Descripcion = ?, DentistaID = ? WHERE ID = ?;",
		turno.PacienteID,
		turno.FechaHora,
		turno.Descripcion,
		turno.DentistaID,
		turno.ID,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *sqlTurno) Exists(Descripcion int) bool {
	//TODO implementar
	return false
}

func (s *sqlTurno) Create(turno domain.Turno) error {
	_, err := s.db.Exec(
		"INSERT INTO Turno (PacienteID, FechaHora, Descripcion, DentistaID) VALUES (?, ?, ?, ?)",
		turno.PacienteID,
		turno.FechaHora,
		turno.Descripcion,
		turno.DentistaID,
	)
	if err != nil {
		return err
	}
	return nil
}
