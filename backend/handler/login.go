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
		return helper.ParseError(helper.ErrBindJSON, e)
	}
	log.Println(req)
	expectedCaptcha, exists := captchaStore[req.CaptchaID]
	delete(captchaStore, req.CaptchaID)

	if !exists || req.Captcha != expectedCaptcha {
		return e.JSON(http.StatusUnauthorized, map[string]string{"message": "CAPTCHA INVALID"})
	}

	// validate input
	if req.Password == "" || req.Username == "" {
		return helper.ParseError(helper.ErrParam, e)
	}

	res, err := h.UR.Login(&req)
	if err != nil {
		helper.Logging(e).Error(err)
		return helper.ParseError(err, e)
	}
	log.Println(res)
	return e.JSON(http.StatusOK, map[string]interface{}{
		"message": "LOGIN SUCCESS",
		"data":    res,
	})
}
