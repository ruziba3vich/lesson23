UPDATE Persons
SET department = $1 WHERE id = $2
RETURNING id, name, age;