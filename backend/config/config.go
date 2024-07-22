package config

import (
	"backend/helper"
	repository "backend/repository"
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"

	"github.com/joho/godotenv"
)

func Connect() (*repository.Database, error) {
	err := godotenv.Load()
	if err != nil {
		return nil, err
	}

	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",
		helper.DB_USER,
		helper.DB_PASS,
		helper.DB_HOST,
		helper.DB_PORT,
		helper.DB_NAME))

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
