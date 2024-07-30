package main

import (
	"fmt"

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

	app.GET("/", func(c *gin.Context) {
		c.JSON(200, gin.H{"message": "Hello, World!"})
	})

	app.POST("/signup", controllers.Signup)

	app.POST("/login", controllers.Login)

	app.GET("/validate", middleware.RequireAuth, controllers.Validate)

	app.Run()
}
