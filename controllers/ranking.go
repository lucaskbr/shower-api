package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

// GET /rankings
// Find all rankig scores
func FindRankings(c *gin.Context) {
	// Get model if exist
	var rankings []models.Ranking
	if err := models.DB.Find(&rankings).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Records not found!"})
		return
	}

	c.JSON(http.StatusOK, rankings)
}
