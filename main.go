package main

import (
	"fmt"

	"github.com/nicobrch/go-jwt/initializers"
)

func init() {
	fmt.Println("Initializing...")
	initializers.LoadEnv()
	initializers.ConnectDb()
	initializers.SyncDb()
}

func main() {
	fmt.Println("Hello, World!")
}
