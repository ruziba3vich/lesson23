package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/databaseLesson/internal/models"
	"github.com/ruziba3vich/databaseLesson/internal/services"
)

func UpdatePersonByAge(c *gin.Context, db *sql.DB) {
	var idAndAge models.IdAndAge

	person, err := services.UpdatePersonByAgeService(idAndAge.Id, idAndAge.Age, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, person)
}
