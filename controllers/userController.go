package controllers

import (
	"net/http"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
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

func Login(c *gin.Context) {
	var body struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}

	if c.Bind(&body) != nil {
		c.JSON(400, gin.H{"error": "Invalid parameters"})
		return
	}

	// Search DB for user
	var user models.User
	result := initializers.DB.Where("email = ?", body.Email).First(&user)

	if result.Error != nil {
		c.JSON(404, gin.H{"error": "User not found"})
		return
	}

	// Compare password
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(body.Password))

	if err != nil {
		c.JSON(401, gin.H{"error": "Invalid password"})
		return
	}

	// Generate JWT token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"sub": user.ID,
		"exp": time.Now().Add(time.Hour * 24 * 30).Unix(),
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))

	if err != nil {
		c.JSON(500, gin.H{"error": "Error generating token"})
		return
	}

	c.SetSameSite(http.SameSiteNoneMode)
	c.SetCookie("token", tokenString, 3600, "/", "", false, true)

	c.JSON(200, gin.H{})
}

func Validate(c *gin.Context) {

}
