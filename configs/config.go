package configs

import (
	"final-project/models"
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadEnv() models.Config {
	var config models.Config

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	config.Port = os.Getenv("PORT")
	config.Database.Host = os.Getenv("DB_HOST")
	config.Database.Port = os.Getenv("DB_PORT")
	config.Database.User = os.Getenv("DB_USER")
	config.Database.Pass = os.Getenv("DB_PASS")
	config.Database.Name = os.Getenv("DB_NAME")

	return config
}
