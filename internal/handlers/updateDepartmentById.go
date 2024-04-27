package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/databaseLesson/internal/models"
	"github.com/ruziba3vich/databaseLesson/internal/services"
)

func UpdateDepartmentById(c *gin.Context, db *sql.DB) {
	var idAndDepartment models.IdAndDepartment
	if err := c.ShouldBindJSON(&idAndDepartment); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	result, err := services.UpdateDepartmentById(idAndDepartment.Id, idAndDepartment.Depatment, db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
	}
	c.JSON(http.StatusOK, result)
}
