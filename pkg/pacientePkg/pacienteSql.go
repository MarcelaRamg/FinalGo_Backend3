package pacientePkg

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"

	_ "github.com/go-sql-driver/mysql"
)

type sqlPaciente struct {
	db *sql.DB
}

func NewSQLPaciente() PacienteInterface {
	database, err := sql.Open("mysql", "breno:root@tcp(localhost:3306)/clinica")
	if err != nil {
		log.Fatalln(err)
	}

	if err := database.Ping(); err != nil {
		log.Fatalln(err)
	}

	return &sqlPaciente{
		db: database,
	}
}

func (s *sqlPaciente) Delete(id int) error {
	_, err := s.db.Exec("DELETE FROM Pacientes WHERE id=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlPaciente) Read(id int) (domain.Paciente, error) {
	paciente := domain.Paciente{}

	rows, err := s.db.Query("SELECT * from Pacientes WHERE id=?", id)
	if err != nil {
		return domain.Paciente{}, err
	}
	for rows.Next() {
		err := rows.Scan(
			&paciente.Id,
			&paciente.Name,
			&paciente.Apellido,
			&paciente.Dni,
			&paciente.FechaNacimiento,
		)
		if err != nil {
			return domain.Paciente{}, err
		}
	}
	return paciente, nil
}

func (s *sqlPaciente) Update(paciente domain.Paciente) error {
	fmt.Println("updating Paciente")
	_, err := s.db.Exec(
		"UPDATE pacientes SET name = ?, apellido = ?, dni = ?, fecha_nacimiento = ? WHERE id = ?;",
		paciente.Name,
		paciente.Apellido,
		paciente.Dni,
		paciente.FechaNacimiento,
		paciente.Id,
	)
	if err != nil {
		fmt.Println(err)
		return err
	}
	return nil
}

func (s *sqlPaciente) Exists(RG string) bool {
	return false
}

func (s *sqlPaciente) Create(paciente domain.Paciente) error {
	_, err := s.db.Exec(
		"INSERT INTO Pacientes (name, apellido, dni, fecha_nacimiento) VALUES (?, ?, ?, ?)",
		paciente.Name,
		paciente.Apellido,
		paciente.Dni,
		paciente.FechaNacimiento,
	)
	if err != nil {
		return err
	}
	return nil
}
