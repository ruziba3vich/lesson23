package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/databaseLesson/internal/services"
)

func GetPeopleFromTwentyFiveToThirtyFive(c *gin.Context, db *sql.DB) {
	results, err := services.GetPersonsFromTwentyFiveToThirtyFive(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, results)
}
