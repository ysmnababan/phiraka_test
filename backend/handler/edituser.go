package handler

import (
	"backend/helper"
	"backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) EditUser(e echo.Context) error {
	var req models.EditReq
	err := e.Bind(&req)
	if err != nil {
		helper.ParseError(helper.ErrBindJSON, e)
	}

	// validate input
	if req.Password == "" || req.Username == "" || req.UserID <= 0 {
		helper.ParseError(helper.ErrParam, e)
	}

	err = h.UR.UpdateUser(&req)
	if err != nil {
		helper.Logging(e).Error(err)
		helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, "USER UPDATED SUCCESSFULLY")
}
