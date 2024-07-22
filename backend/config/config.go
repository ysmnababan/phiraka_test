package config

import (
	repository "backend/repository"
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func Connect() (*repository.Database, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		os.Getenv("DB_USER"),
		os.Getenv("DB_PASS"),
		os.Getenv("DB_HOST"),
		os.Getenv("DB_PORT"),
		os.Getenv("DB_NAME")))

	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	return &repository.Database{DB: db}, err
}
