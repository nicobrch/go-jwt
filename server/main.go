package main

import (
	"fmt"

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

	// React App routing
	app.Use(cors.Default())

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	// API routing
	app.POST("/api/v1/signup", controllers.Signup)
	app.POST("/api/v1/login", controllers.Login)
	app.GET("/api/v1/validate", middleware.RequireAuth, controllers.Validate)

	app.Run()
}
