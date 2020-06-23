package models

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	FirstName string 
	LastName  string 
	Username			string `gorm:"unique; not null"`
	Password 	string `gorm:"not null"`
	IsActive 	bool 	 `gorm:"default: true; not null"`
	Ranking Ranking
	ShowerConfig ShowerConfig
	FamilyGroupID uint `gorm:"default: null"`
//	ShowerRegister []ShowerRegister
//	FamilyGroupAdmin FamilyGroup `gorm:"foreignkey:admin_id;association_foreignkey:id"`
//	FamilyGroup    []FamilyGroup `gorm:"many2many:family_group_user;"`
}
