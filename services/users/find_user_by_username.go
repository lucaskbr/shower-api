package services

import (
	"errors"

	"github.com/rahmanfadhil/gin-bookstore/models"
)

func FindUserByUsername(username string) (*models.User, error) {
	var user = &models.User{}

	err := models.DB.Where(models.User{ Username: username }).First(&user).Error; 
	
	if (err != nil) {
		return nil, errors.New("Username not found in the database")
	}
	return user, nil

}