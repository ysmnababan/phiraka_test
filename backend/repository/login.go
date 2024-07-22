package repository

import (
	"backend/helper"
	"backend/models"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (r *Database) Login(in *models.LoginReq) (*models.User, error) {
	u := models.User{}
	err := r.DB.QueryRow("SELECT UserID, Username, Password, CreateTime FROM users WHERE Username = ?", in.Username).Scan(&u.UserID, &u.Username, u.Password, u.CreateTime)
	if err != nil {
		if err == sql.ErrNoRows {
			return nil, helper.ErrCredential
		}
		helper.Logging(nil).Error("ERROR REPO: ", err)
		return nil, helper.ErrQuery
	}

	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password))
	if err != nil {
		return nil, helper.ErrCredential
	}

	log.Println(u)
	return &u, nil
}
