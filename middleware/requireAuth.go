package middleware

import (
	"fmt"
	"os"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/nicobrch/go-jwt/initializers"
	"github.com/nicobrch/go-jwt/models"
)

func RequireAuth(c *gin.Context) {
	// Get the cookie off the request
	tokenString, err := c.Cookie("Authorization")

	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	// Decode/validate the JWT
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		// Check if the signing method is correct
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("SECRET")), nil
	})

	if err != nil {
		c.AbortWithStatus(401)
		return
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {
		// Check if the expiration date is valid
		if int64(claims["exp"].(float64)) < time.Now().Unix() {
			c.AbortWithStatus(401)
			return
		}

		// Find the user with token sub
		var user models.User
		initializers.DB.Where("id = ?", claims["sub"]).First(&user)

		if user.ID == 0 {
			c.AbortWithStatus(401)
			return
		}

		// Attach the user to the context
		c.Set("user", user)

		// Continue
		c.Next()

	} else {
		c.AbortWithStatus(401)
		return
	}
}
