package initializers

import "github.com/nicobrch/go-jwt/models"

func SyncDb() {
	DB.AutoMigrate(&models.User{})
}
