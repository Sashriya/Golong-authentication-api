package main

import (
	"golang-auth/database"
	"golang-auth/handlers"
	"golang-auth/middleware"
	"golang-auth/models"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()

	database.Connect()

	database.DB.AutoMigrate(&models.User{})

	r.POST("/signup", handlers.Signup)
	r.POST("/login", handlers.Login)

	protected := r.Group("/")

	protected.Use(middleware.AuthMiddleware())

	protected.GET("/profile", handlers.Profile)
	protected.GET("/users", handlers.GetUsers)

	r.Run(":8080")
}