package handlers

import (
	"database/sql"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/ruziba3vich/databaseLesson/internal/services"
)

func GetOlderThanThirty(c *gin.Context, db *sql.DB) {
	olderThanThirty, err := services.GetOlderThanThirtyService(db)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	c.JSON(http.StatusOK, olderThanThirty)
}
