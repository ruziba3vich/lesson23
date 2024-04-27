SELECT p.department, COUNT(p.id) AS counts
FROM Persons p
GROUP BY p.department;