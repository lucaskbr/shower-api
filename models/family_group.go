package models

import (
	"github.com/jinzhu/gorm"
)

type FamilyGroup struct {
	gorm.Model
	AdminID uint	`json:"adminId"`
	Users []User	`json:"users"`
}
