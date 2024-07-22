package handler

import (
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) Register(e echo.Context) error {

	return e.JSON(http.StatusCreated, "")
}
