package handler

import (
	"backend/helper"
	"backend/models"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
)

func (h *UserHandler) Register(e echo.Context) error {
	var req models.RegisterReq
	err := e.Bind(&req)
	if err != nil {
		helper.ParseError(helper.ErrBindJSON, e)
	}

	// validate input
	if req.Password == "" || req.Username == "" {
		helper.ParseError(helper.ErrParam, e)
	}

	res, err := h.UR.Register(&req)
	if err != nil {
		helper.Logging(e).Error(err)
		helper.ParseError(err, e)
	}
	log.Println(res)
	return e.JSON(http.StatusCreated, res)
}
