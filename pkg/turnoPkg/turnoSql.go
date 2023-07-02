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

func NewSQLTurno() TurnoInterface {
	database, err := sql.Open("mysql", "breno:root@tcp(localhost:3306)/clinica")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlTurno{
		db: database,
	}
}

func (s *sqlTurno) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM Turnos WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlTurno) Read(id int) (domain.Turno, error) {
	turno := domain.Turno{}

	rows, err := s.db.Query("SELECT * from Turnos WHERE id=?", id)
	if err != nil {
		return domain.Turno{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&turno.ID,
			&turno.FechaHora,
			&turno.PacienteID,
			&turno.Descripcion,
			&turno.DentistaID,
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
		"UPDATE Turnos SET PacienteID = ?, FechaHora = ?, Descripcion = ?, DentistaID = ? WHERE id = ?;",
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
	return false
}

func (s *sqlTurno) Create(turno domain.Turno) error {
	_, err := s.db.Exec(
		"INSERT INTO Turnos (PacienteID, FechaHora, Descripcion, DentistaID) VALUES (?, ?, ?, ?)",
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