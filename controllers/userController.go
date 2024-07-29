package controllers

import (
	"github.com/gin-gonic/gin"
	"github.com/nicobrch/go-jwt/initializers"
	"github.com/nicobrch/go-jwt/models"
	"golang.org/x/crypto/bcrypt"
)

func Signup(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid parameters"})
		return
	}

	// Hash password
	hash, err := bcrypt.GenerateFromPassword([]byte(body.Password), bcrypt.DefaultCost)

	if err != nil {
		c.JSON(500, gin.H{"error": "Error hashing password"})
		return
	}

	user := models.User{Email: body.Email, Password: string(hash)}
	result := initializers.DB.Create(&user)

	if result.Error != nil {
		c.JSON(500, gin.H{"error": "Error creating user"})
		return
	}

	c.JSON(200, gin.H{"message": "User created"})
}
