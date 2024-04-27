package services

import (
	"database/sql"
	"os"
)

/// getCountOfPersonsInEachDepartment.sql

func GetNumberOfPersonsInEachDepartment(db *sql.DB) (map[string]int, error) {
	fileName := "../db/getCountOfPersonsInEachDepartment.sql"
	file, err := os.ReadFile(fileName)

	if err != nil {
		return nil, err
	}

	rows, err := db.Query(string(file))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	counts := make(map[string]int)

	for rows.Next() {
		var department string
		var count int
		err := rows.Scan(&department, &count)
		if err != nil {
			return nil, err
		}
		counts[department] = count
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return counts, nil
}
