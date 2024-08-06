package main

import (
	"fmt"
	"time"

	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/nicobrch/go-jwt/controllers"
	"github.com/nicobrch/go-jwt/initializers"
	"github.com/nicobrch/go-jwt/middleware"
)

func init() {
	fmt.Println("Initializing...")
	initializers.LoadEnv()
	initializers.ConnectDb()
	initializers.SyncDb()
}

func main() {
	app := gin.Default()

	app.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"http://localhost:5173"}, // Cambia esto a tu origen
		AllowMethods:     []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowHeaders:     []string{"Origin", "Content-Type", "Accept", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
	}))

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	app.POST("/signup", controllers.Signup)

	app.POST("/login", controllers.Login)

	app.POST("/validate", middleware.RequireAuth, controllers.Validate)

	app.POST("/logout", middleware.RequireAuth, controllers.Logout)

	app.Run()
}
