package helper

import (
	"os"

	"github.com/joho/godotenv"
)

var (
	DB_USER string
	DB_PASS string
	DB_HOST string
	DB_PORT string
	DB_NAME string
	PORT    string
)

func LoadEnv() {
	err := godotenv.Load()
	if err != nil {
		panic("unable to get .env")
	}

	DB_USER = os.Getenv("DB_USER")
	DB_PASS = os.Getenv("DB_PASS")
	DB_HOST = os.Getenv("DB_HOST")
	DB_PORT = os.Getenv("DB_PORT")
	DB_NAME = os.Getenv("DB_NAME")
	PORT = os.Getenv("PORT")
}
