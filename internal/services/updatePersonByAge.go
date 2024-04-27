package services

import (
	"database/sql"
	"lesson23/internal/models"
	"os"
)

func UpdatePersonByAgeService(id, newAge int, db *sql.DB) (*models.Person, error) {
	fileName := "../db/updatePersonByAge.sql"
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var person models.Person

	err = db.QueryRow(string(file), newAge, id).Scan(person.Id, person.Name, person.Department)
	if err != nil {
		return nil, err
	}
	return &person, nil
}
