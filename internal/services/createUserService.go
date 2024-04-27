package services

import (
	"database/sql"
	"github.com/ruziba3vich/databaseLesson/internal/models"
	"os"
)

func CreatePersonService(person models.Person, db *sql.DB) (*models.Person, error) {
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
