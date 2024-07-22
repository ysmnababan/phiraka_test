package helper

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
)

var (
	ErrNoData       = errors.New("no data in result set")
	ErrNoUser       = errors.New("no user exist")
	ErrQuery        = errors.New("query execution failed")
	ErrInvalidId    = errors.New("invalid id")
	ErrUserExists   = errors.New("user already exist")
	ErrNoUpdate     = errors.New("data already exists")
	ErrBindJSON     = errors.New("unable to bind json")
	ErrParam        = errors.New("error or missing parameter")
	ErrCredential   = errors.New("password or email doesn't match")
	ErrGeneratedPwd = errors.New("error generating password hash")
)

func ParseError(err error, ctx echo.Context) error {
	status := http.StatusOK

	switch {
	case errors.Is(err, ErrQuery):
		fallthrough
	case errors.Is(err, ErrGeneratedPwd):
		fallthrough
	case errors.Is(err, ErrNoUser):
		status = http.StatusNotFound

	case errors.Is(err, ErrNoData):
		status = http.StatusNotFound

	case errors.Is(err, ErrParam):
		status = http.StatusBadRequest

	case errors.Is(err, ErrBindJSON):
		status = http.StatusBadRequest

	case errors.Is(err, ErrInvalidId):
		status = http.StatusBadRequest

	case errors.Is(err, ErrCredential):
		status = http.StatusBadRequest

	case errors.Is(err, ErrUserExists):
		status = http.StatusBadRequest

	case errors.Is(err, ErrNoUpdate):
		status = http.StatusBadRequest

	default:
		status = http.StatusInternalServerError

	}

	return ctx.JSON(status, map[string]interface{}{"message": err.Error()})
}
