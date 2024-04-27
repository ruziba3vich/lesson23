package services

import (
	"database/sql"
	"github.com/ruziba3vich/databaseLesson/internal/models"
	"os"
)

func GetPeopleStartingWithA(db *sql.DB) (persons []models.Person, e error) {
	fileName := "../db/getPersonsByNameA.sql"
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(string(file))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var person models.Person
		err := rows.Scan(&person.Id, &person.Name, &person.Age, &person.Department)
		if err != nil {
			return nil, err
		}
		persons = append(persons, person)
	}
	return persons, nil
}
