package services

import (
	"database/sql"
	"os"

	"github.com/ruziba3vich/databaseLesson/internal/models"
)

// getEldestFivePersons.sql

func GetEldestFivePersons(db *sql.DB) ([]models.Person, error) {
	fileName := "../db/getEldestFivePersons.sql"
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}
	rows, err := db.Query(string(file))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var persons []models.Person

	for rows.Next() {
		var person models.Person
		err := rows.Scan(&person.Id, &person.Name, &person.Age, &person.Department)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return persons, nil
}
