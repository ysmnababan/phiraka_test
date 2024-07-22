package repository

import (
	"backend/helper"
	"backend/models"
)

func (r *Database) DeleteUser(in *models.DeleteReq) error {
	exists, err := r.usernameExists(in.Username)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return helper.ErrQuery
	}
	if !exists {
		return  helper.ErrNoUser
	}

	_, err = r.DB.Exec("DELETE FROM users WHERE Username = ? ", in.Username)
	if err != nil {
		helper.Logging(nil).Error("ERR REPO: ", err)
		return helper.ErrQuery
	}

	return nil
}
