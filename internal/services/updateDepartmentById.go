package services

import (
	"database/sql"
	"os"

	"github.com/ruziba3vich/databaseLesson/internal/models"
)

// updateDepartmentById

func UpdateDepartmentById(id int, newDepartment string, db *sql.DB) (*models.Person, error) {
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
