package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rahmanfadhil/gin-bookstore/models"
	services "github.com/rahmanfadhil/gin-bookstore/services/users"
)

type CreateUserInput struct {
	FirstName  string `json:"firstName" binding:"required"`
	LastName string `json:"lastName" binding:"required"`
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}


// GET /books
// Find all books
func FindUsers(c *gin.Context) {
	var users []models.User
	models.DB.Find(&users)

	c.JSON(http.StatusOK, gin.H{"data": users})
}

// GET /books/:id
// Find a book
func FindUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// POST /books
// Create new book
func CreateUser(c *gin.Context) {
	// Validate input
	print(c.Request.Body)
	var input CreateUserInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	isUserExists, _ := services.FindUserByUsername(input.Username); 
	
	if isUserExists != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "User already exists"})
		return
	}

	user := models.User{ FirstName: input.FirstName, LastName: input.LastName, Username: input.Username, Password: input.Password }
	models.DB.Create(&user)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

/*

// PATCH /books/:id
// Update a book
func UpdateUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	// Validate input
	var input UpdateBookInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	models.DB.Model(&user).Updates(input)

	c.JSON(http.StatusOK, gin.H{"data": user})
}

// DELETE /books/:id
// Delete a book
func DeleteUser(c *gin.Context) {
	// Get model if exist
	var user models.User
	if err := models.DB.Where("id = ?", c.Param("id")).First(&user).Error; err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Record not found!"})
		return
	}

	models.DB.Delete(&user)

	c.JSON(http.StatusOK, gin.H{"data": true})
}
*/