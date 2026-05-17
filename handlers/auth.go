package handlers

import (
	"golang-auth/database"
	"golang-auth/models"
	"golang-auth/utils"
	"net/http"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {

	var user models.User

	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid input",
		})
		return
	}

	hashedPassword, _ := bcrypt.GenerateFromPassword(
		[]byte(user.Password),
		14,
	)

	user.Password = string(hashedPassword)

	if user.Role == "" {
		user.Role = "user"
	}

	database.DB.Create(&user)

	token, _ := utils.GenerateJWT(user.ID, user.Role)

	c.JSON(http.StatusOK, gin.H{
		"message": "Signup successful",
		"token":   token,
	})
}

func Login(c *gin.Context) {

	var input models.User
	var user models.User

	c.ShouldBindJSON(&input)

	database.DB.Where("email = ?", input.Email).First(&user)

	if user.ID == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid email",
		})
		return
	}

	err := bcrypt.CompareHashAndPassword(
		[]byte(user.Password),
		[]byte(input.Password),
	)

	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{
			"error": "Invalid password",
		})
		return
	}

	token, _ := utils.GenerateJWT(user.ID, user.Role)

	c.JSON(http.StatusOK, gin.H{
		"message": "Login successful",
		"token":   token,
	})
}