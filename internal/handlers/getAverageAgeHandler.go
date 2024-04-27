package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/databaseLesson/internal/services"
)

func GetAverageAgeHandler(c *gin.Context, db *sql.DB) {
	avgAge, err := services.GetAverageAge(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, gin.H{"averageAge": avgAge})
}
