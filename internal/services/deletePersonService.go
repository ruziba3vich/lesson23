package services

import (
	"database/sql"
	"os"
)

func DeletePersonByIdService(id int, db *sql.DB) error {
	fileName := "../db/deletePerson.sql"
	file, err := os.ReadFile(fileName)

	if err != nil {
		return err
	}

	_, err = db.Exec(string(file), id)
	if err != nil {
		return err
	}
	return nil
}
