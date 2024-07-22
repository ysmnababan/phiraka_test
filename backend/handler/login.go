package handler

import (
	"backend/repository"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UR repository.UserRepo
}

func (h *UserHandler) Login(e echo.Context) error {

	return e.JSON(http.StatusOK, "")
}
