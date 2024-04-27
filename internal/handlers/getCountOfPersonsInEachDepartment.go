package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/databaseLesson/internal/services"
)

func GetCountOfPersonsInEachDepartment(c *gin.Context, db *sql.DB) {
	mp, err := services.GetNumberOfPersonsInEachDepartment(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, mp)
}
