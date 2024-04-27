package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/databaseLesson/internal/services"
)

func UpdatePersonByAge(c *gin.Context, id int, db *sql.DB) {
	var age int

	if err := c.ShouldBindJSON(&age); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}

	person, err := services.UpdatePersonByAgeService(id, age, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, person)
}
