package services

import (
	"database/sql"
	"os"

	"github.com/ruziba3vich/databaseLesson/internal/models"
)

// getNameAndDepartmentOnly.sql

func GetNameAndDepartmentOnlyService(db *sql.DB) (persons []models.NameAndDepartmentObject, e error) {
	fileName := "../db/getNameAndDepartmentOnly.sql"
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(string(file))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var person models.NameAndDepartmentObject
		err := rows.Scan(&person.Name, &person.Department)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}
