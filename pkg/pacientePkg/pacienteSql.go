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
	database, err := sql.Open("mysql", "root:toor@tcp(localhost:3306)/clinica")
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
	_, err := s.db.Exec("DELETE FROM Pacientes WHERE ID=?", id)
	if err != nil {
		return err
	}

	return nil
}

func (s *sqlPaciente) Read(id int) (domain.Paciente, error) {
	var paciente domain.Paciente

	query := "SELECT * from Pacientes WHERE ID=?"
	row := s.db.QueryRow(query, id)
	err := row.Scan(&paciente.ID, &paciente.Nombre, &paciente.Apellido, &paciente.Dni, &paciente.FechaAlta)
	if err != nil {
		return domain.Paciente{}, err
	}
	return paciente, nil
}

func (s *sqlPaciente) Update(paciente domain.Paciente) error {
	fmt.Println("updating Paciente")
	_, err := s.db.Exec(
		"UPDATE Pacientes SET Nombre = ?, Apellido = ?, Dni = ?, FechaAlta = ? WHERE ID = ?;",
		paciente.Nombre,
		paciente.Apellido,
		paciente.Dni,
		paciente.FechaAlta,
		paciente.ID,
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
		"INSERT INTO Pacientes (Nombre, Apellido, Dni, FechaAlta) VALUES (?, ?, ?, ?)",
		paciente.Nombre,
		paciente.Apellido,
		paciente.Dni,
		paciente.FechaAlta,
	)
	if err != nil {
		return err
	}
	return nil
}
