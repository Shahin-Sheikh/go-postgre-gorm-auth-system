package main

import (
	"fmt"
	"github.com/Shahin-Sheikh/go-postgre-gorm-auth-system/initializers"
	"github.com/Shahin-Sheikh/go-postgre-gorm-auth-system/models"
	"log"
)

func init() {
	config, err := initializers.LoadConfig(".")
	if err != nil {
		log.Fatal("🚀 Could not load environment variables", err)
	}

	initializers.ConnectDB(&config)
}

func main() {
	initializers.DB.Exec("CREATE EXTENSION IF NOT EXISTS \"uuid-ossp\"")
	initializers.DB.AutoMigrate(&models.User{}, &models.Post{})
	fmt.Println("👍 Migration complete")
}
