UPDATE Persons
SET age = $1 WHERE id = $2
RETURNING id, name, department;