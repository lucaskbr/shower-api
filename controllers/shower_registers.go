package controllers

import (
	"net/http"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/models"
)

type CreateShowerRegisterInput struct {
	BeginDate  string `json:"beginDate" binding:"required"`
	EndDate 	 string `json:"endDate" binding:"required"`
}


// GET /shower-registers/user/:id
// Find all shower registers by user
func FindShowerRegisters(c *gin.Context) {
	// Get model if exist
	var showerRegisters []models.ShowerRegister
	if err := models.DB.Where("user_id = ?", c.Param("id")).Find(&showerRegisters).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Records not found!"})
		return
	}

	c.JSON(http.StatusOK, showerRegisters)
}

// POST /shower-registers/user/:id
// Create a shower register for a user
func CreateShowerRegister(c *gin.Context) {
	var input CreateShowerRegisterInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	userID, err := strconv.ParseUint(c.Param("id"), 10, 8)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	
	layout := "2006-01-02T15:04:05"

	beginDate, err := time.Parse(layout, input.BeginDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	endDate, err := time.Parse(layout, input.EndDate)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	showerRegister := models.ShowerRegister{ BeginDate: beginDate, EndDate: endDate, UserID: uint(userID) }

	err = models.DB.Create(&showerRegister).Error

	if (err != nil) {
		println(err.Error())
	}

	var ranking models.Ranking

	if err := models.DB.Where("user_id = ?", c.Param("id")).Find(&ranking).Error; err != nil {
		
		ranking.UserID = uint(userID)
		ranking.Score = 5

		models.DB.Save(&ranking)

	} else {
		ranking.Score += 10;
		models.DB.Update(&ranking);
	}

	c.Writer.WriteHeader(http.StatusCreated);

}

// DELETE /shower-registers/:id
// Delete a shower regiser
func DeleteShowerRegisters(c *gin.Context) {
	// Get model if exist
	var showerRegister models.ShowerRegister
	if err := models.DB.Where("id = ?", c.Param("id")).First(&showerRegister).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&showerRegister)

	c.JSON(http.StatusOK, gin.H{})
}