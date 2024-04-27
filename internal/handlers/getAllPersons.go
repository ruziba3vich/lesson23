package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/databaseLesson/internal/services"
)

func GetAllPersons(c *gin.Context, db *sql.DB) {
	result, err := services.GetPersonsService(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, result)
}
