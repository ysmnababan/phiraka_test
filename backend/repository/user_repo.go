package repository

import (
	"backend/models"
	"database/sql"
)

type Database struct {
	DB *sql.DB
}

type UserRepo interface {
	Login(in *models.LoginReq) (*models.User, error)
	Register(in *models.RegisterReq) (*models.User, error)
	GetAllUser() ([]models.User, error)
	UpdateUser(in *models.EditReq) (*models.User, error)
	DeleteUser(in *models.DeleteReq) (*models.User, error)
}

func (r *Database) Close() error {
	return r.DB.Close()
}
