package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/databaseLesson/internal/models"
	"github.com/ruziba3vich/databaseLesson/internal/services"
)

func CreatePersonHandler(c *gin.Context, db *sql.DB) {
	var person models.Person
	if err := c.ShouldBindJSON(&person); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	newUser, err := services.CreatePersonService(person, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusCreated, gin.H{"error": newUser})
}
