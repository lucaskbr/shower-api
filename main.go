package main

import (
	"github.com/rahmanfadhil/gin-bookstore/controllers"
	"github.com/rahmanfadhil/gin-bookstore/models"

	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// Connect to database
	models.ConnectDatabase()

	// Routes
	router.GET("/users", controllers.FindUsers)
	router.POST("/users", controllers.CreateUser)

	router.GET("/shower-registers/user/:id", controllers.FindShowerRegisters)
	router.POST("/shower-registers/user/:id", controllers.CreateShowerRegister)
	router.DELETE("/shower-registers/:id", controllers.DeleteShowerRegisters)

	router.GET("/rankings", controllers.FindRankings)
	
	router.POST("/auth/login", controllers.Login)


	// Run the server
	router.Run()
}
