package repository

import (
	"backend/helper"
	"backend/models"
	"database/sql"
	"log"

	"golang.org/x/crypto/bcrypt"
)

func (r *Database) UpdateUser(in *models.EditReq) error {
	var u models.User
	err := r.DB.QueryRow("SELECT Username, Password FROM users WHERE UserID = ?", in.UserID).Scan(&u.Username, &u.Password)
	if err != nil {
		if err == sql.ErrNoRows {
			return helper.ErrNoUser
		}
		helper.Logging(nil).Error("error query: ", err)
		return helper.ErrQuery
	}

	if u.Username != in.Username {
		exists, err := r.usernameExists(in.Username)
		if err != nil {
			helper.Logging(nil).Error("error query: ", err)
			return helper.ErrQuery
		}
		if exists {
			return helper.ErrUserExists
		}
	}

	var hashedpwd []byte
	err = bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(in.Password))
	if err != nil {
		hashedpwd, err = bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
		if err != nil {
			return helper.ErrGeneratedPwd
		}
		log.Println(hashedpwd)
	} else {
		hashedpwd = []byte(u.Password)
	}

	_, err = r.DB.Exec("UPDATE users SET Username = ? , Password = ? WHERE UserID = ? ", in.Username, hashedpwd, in.UserID)
	if err != nil {
		helper.Logging(nil).Error("error query: ", err)
		return helper.ErrQuery
	}

	return nil
}
