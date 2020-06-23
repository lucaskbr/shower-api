package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	services "github.com/rahmanfadhil/gin-bookstore/services/users"
)

type LoginInput struct {
	Username    string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// POST /login
func Login(c *gin.Context) {

	var input LoginInput
	err := c.ShouldBindJSON(&input); 
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	user, err := services.FindUserByUsername(input.Username); 
	
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if (user.Password != input.Password){
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid login!"})
		return
	}

	c.JSON(http.StatusOK, user)


}