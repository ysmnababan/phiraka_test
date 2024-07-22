package handler

import (
	"backend/helper"
	"backend/models"
	"backend/repository"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

type UserHandler struct {
	UR repository.UserRepo
}

func (h *UserHandler) Login(e echo.Context) error {
	var req models.LoginReq
	err := e.Bind(&req)
	if err != nil {
		helper.ParseError(helper.ErrBindJSON, e)
	}

	// validate input
	if req.Password == "" || req.Username == "" {
		helper.ParseError(helper.ErrParam, e)
	}

	res, err := h.UR.Login(&req)
	if err != nil {
		helper.Logging(e).Error(err)
		helper.ParseError(err, e)
	}
	log.Println(res)
	return e.JSON(http.StatusOK, "LOGIN SUCCESS")
}
