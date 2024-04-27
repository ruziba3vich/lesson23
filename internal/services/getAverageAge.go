package services

import (
	"database/sql"
	"os"
)

func GetAverageAge(db *sql.DB) (float64, error) {
	var averageAge float64
	fileName := "../db/deletePerson.sql"
	file, err := os.ReadFile(fileName)

	if err != nil {
		return 0, err
	}

	err = db.QueryRow(string(file)).Scan(&averageAge)
	if err != nil {
		return 0, err
	}

	return averageAge, nil
}
