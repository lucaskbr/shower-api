package models

import (
	"github.com/jinzhu/gorm"
)

type ShowerConfig struct {
	gorm.Model
  UserID   uint `json:"userId"`
}
