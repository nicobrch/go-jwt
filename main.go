package main

import (
	"fmt"

	"github.com/nicobrch/go-jwt/initializers"
)

func init() {
	fmt.Println("Initializing...")
	initializers.LoadEnv()
	initializers.ConnectDb()
}

func main() {
	fmt.Println("Hello, World!")
}
