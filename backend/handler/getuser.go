package handler

import (
	"backend/helper"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) GetAllUser(e echo.Context) error {
	res, err := h.UR.GetAllUser()
	if err != nil {
		helper.Logging(e).Error(err)
		return helper.ParseError(err, e)
	}

	return e.JSON(http.StatusOK, res)
}
