package initializers

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() {
	// Load environment variables
	err := godotenv.Load()

	if err != nil {
		log.Fatal("Error loading .env file")
	}

	envVars := []string{
		"PORT",
		"JWT_SECRET",
	}

	for _, envVar := range envVars {
		if value := os.Getenv(envVar); value == "" {
			log.Fatalf("Error loading environment variable %s", envVar)
		}
	}
}
