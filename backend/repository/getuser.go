package repository

import (
	"backend/helper"
	"backend/models"
)

func (r *Database) GetAllUser() ([]*models.User, error) {

	rows, err := r.DB.Query("SELECT UserID, Username, Password, CreateTime FROM users")
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return nil, helper.ErrQuery
	}

	var users []*models.User
	for rows.Next() {
		var u models.User
		err := rows.Scan(&u.UserID, &u.Username, &u.Password, &u.CreateTime)
		if err != nil {
			helper.Logging(nil).Error("ERR REPO: ", err)
			return nil, helper.ErrQuery
		}
		users = append(users, &u)
	}

	return users, nil
}
