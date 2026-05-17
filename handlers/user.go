package handlers

import (
	"golang-auth/database"
	"golang-auth/models"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Profile(c *gin.Context) {

	userID := c.GetFloat64("userID")

	var user models.User

	database.DB.First(&user, uint(userID))

	c.JSON(http.StatusOK, user)
}

func GetUsers(c *gin.Context) {

	role := c.GetString("role")

	if role != "admin" {
		c.JSON(http.StatusForbidden, gin.H{
			"error": "Access denied",
		})
		return
	}

	var users []models.User

	database.DB.Find(&users)

	c.JSON(http.StatusOK, users)
}