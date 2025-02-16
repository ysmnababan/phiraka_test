package handler

import (
	"backend/helper"
	"backend/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) DeleteUser(e echo.Context) error {
	var req models.DeleteReq
	err := e.Bind(&req)
	if err != nil {
		return helper.ParseError(helper.ErrBindJSON, e)
	}

	// validate input
	if req.Username == "" {
		return helper.ParseError(helper.ErrParam, e)
	}

	err = h.UR.DeleteUser(&req)
	if err != nil {
		helper.Logging(e).Error(err)
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, "USER DELETED SUCCESSFULLY")
}
