package services

import (
	"database/sql"
	"lesson23/internal/models"
	"os"
)

func CreatePersonHandler(person models.Person, db *sql.DB) (*models.Person, error) {
	fileName := "../db/createNewPerson.sql"
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	err = db.QueryRow(string(file), person.Name, person.Age, person.Department).Scan(&person.Id)
	if err != nil {
		return nil, err
	}

	return &person, nil
}
