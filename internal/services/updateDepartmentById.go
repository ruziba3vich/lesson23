package services

import (
	"database/sql"
	"lesson23/internal/models"
	"os"
)

// updateDepartmentById

func updateDepartmentById(id int, newDepartment string, db *sql.DB) (*models.Person, error) {
	fileName := "../db/updatePersonByAge.sql"
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	var person models.Person

	err = db.QueryRow(string(file), newDepartment, id).Scan(person.Id, person.Name, person.Age)
	if err != nil {
		return nil, err
	}
	return &person, nil
}
