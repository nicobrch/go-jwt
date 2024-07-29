package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/nicobrch/go-jwt/controllers"
	"github.com/nicobrch/go-jwt/initializers"
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

	app.Run()
}
