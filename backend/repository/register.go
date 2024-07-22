package repository

import (
	"backend/helper"
	"backend/models"
	"time"

	"golang.org/x/crypto/bcrypt"
)

// Function to check if a username exists
func (r *Database) usernameExists(username string) (bool, error) {
	var count int
	err := r.DB.QueryRow("SELECT COUNT(*) FROM users WHERE Username = ?", username).Scan(&count)
	if err != nil {
		helper.Logging(nil).Error("error query: ", err)
		return false, err
	}
	return count > 0, nil
}
func (r *Database) Register(in *models.RegisterReq) (*models.User, error) {
	var u models.User

	exists, err := r.usernameExists(in.Username)
	if err != nil {
		return nil, helper.ErrQuery
	}
	if exists {
		return nil, helper.ErrUserExists
	}

	// hashing pwd
	hashedpwd, err := bcrypt.GenerateFromPassword([]byte(in.Password), bcrypt.DefaultCost)
	if err != nil {
		helper.Logging(nil).Error("error generating pwd: ", err)
		return nil, helper.ErrGeneratedPwd
	}

	u.CreateTime = time.Now()
	res, err := r.DB.Exec("INSERT INTO users (Username, Password, CreateTime) VALUES (?, ?, ? )", in.Username, hashedpwd, u.CreateTime)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return nil, helper.ErrQuery
	}
	u.Username = in.Username
	id, _ := res.LastInsertId()
	u.UserID = int(id)

	return &u, nil
}
