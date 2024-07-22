package repository

import (
	"backend/helper"
	"backend/models"

	"golang.org/x/crypto/bcrypt"
)

func (r *Database) UpdateUser(in *models.EditReq) error {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM users WHERE UserID = ?", in.UserID).Scan(&count)
	if err != nil {
		helper.Logging(nil).Error("error query: ", err)
		return helper.ErrQuery
	}

	if count == 0 {
		return helper.ErrNoUser
	}

	exists, err := r.usernameExists(in.Username)
	if err != nil {
		helper.Logging(nil).Error("error query: ", err)
		return  helper.ErrQuery
	}
	if exists {
		return  helper.ErrUserExists
	}

	hashedpwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		return  helper.ErrGeneratedPwd
	}

	_, err = r.DB.Exec("UPDATE users SET Username = ? , Password = ? WHERE UserID = ? ", in.Username, hashedpwd, in.UserID)
	if err != nil {
		helper.Logging(nil).Error("error query: ", err)
		return helper.ErrQuery
	}

	return nil
}
