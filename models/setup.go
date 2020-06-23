package models

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

var DB *gorm.DB

func ConnectDatabase() {
	connStr := "host=localhost port=5432 user=postgres dbname=shower password=docker sslmode=disable"

	database, err := gorm.Open("postgres", connStr)

	if err != nil {
		panic(err)
	}

	database.AutoMigrate(&User{})
	database.AutoMigrate(&Ranking{})
	database.AutoMigrate(&ShowerRegister{})
	database.AutoMigrate(&ShowerConfig{})
	database.AutoMigrate(&FamilyGroup{})


	database.Model(&ShowerConfig{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	database.Model(&ShowerRegister{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	database.Model(&Ranking{}).AddForeignKey("user_id", "users(id)", "CASCADE", "CASCADE")
	database.Model(&FamilyGroup{}).AddForeignKey("admin_id", "users(id)", "CASCADE", "CASCADE")
	database.Model(&User{}).AddForeignKey("family_group_id", "family_groups(id)", "CASCADE", "CASCADE")

	DB = database
}
