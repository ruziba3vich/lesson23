INSERS INTO Persons(
    name,
    age,
    department
) VALUES ($1, $2, $3)
    RETURNING id;