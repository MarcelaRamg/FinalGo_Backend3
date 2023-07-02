package dentistaPkg

import (
	"database/sql"
	"log"

	"github.com/MarcelaRamg/FinalBack3.git/internal/domain"
)

type sqlDentista struct {
	DB *sql.DB
}

func NewSQLDentista(db *sql.DB) DentistaInterface {

	err := db.Ping()
	if err != nil {
		log.Fatalln(err)
	}

	return &sqlDentista{
		DB: db,
	}
}

func (s *sqlDentista) Read(id int) (domain.Dentista, error) {

	var dentistaReturn domain.Dentista

	query := "SELECT * FROM dentista WHERE ID = ?;"
	row := s.DB.QueryRow(query, id)
	err := row.Scan(&dentistaReturn.ID, &dentistaReturn.Apellido, &dentistaReturn.Nombre, &dentistaReturn.Matricula)
	if err != nil {
		return domain.Dentista{}, err
	}
	return dentistaReturn, nil //va aqui un puntero?
}

func (s *sqlDentista) Create(dentista domain.Dentista) error {

	query := "INSERT INTO dentista(Apellido, Nombre, Matricula) VALUES(?, ?, ?)"
	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(dentista.Apellido, dentista.Nombre, dentista.Matricula)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlDentista) Update(dentista domain.Dentista) error {

	query := "UPDATE dentista SET Apellido = ?, Nombre = ?, Matricula = ?  WHERE ID = ?;"

	stmt, err := s.DB.Prepare(query)
	if err != nil {
		return err
	}
	defer stmt.Close()
	res, err := stmt.Exec(dentista.Apellido, dentista.Nombre, dentista.Matricula, dentista.ID)
	if err != nil {
		return err
	}

	_, err = res.RowsAffected()
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlDentista) Delete(id int) error {

	query := "DELETE FROM dentista WHERE ID = ?"
	_, err := s.DB.Exec(query, id)
	if err != nil {
		return err
	}
	return nil
}

func (s *sqlDentista) Exists(codeValue string) bool {

	query := "SELECT ID FROM dentista WHERE Matricula like(%?%)"
	row := s.DB.QueryRow(query, codeValue)
	var id int
	err := row.Scan(&id)
	if err != nil {
		return false
	}

	if id > 0 {
		return true
	}
	return false
}
