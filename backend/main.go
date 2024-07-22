package main

import (
	"backend/config"
	"backend/handler"
	"backend/helper"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func main() {
	db, err := config.Connect()
	if err != nil {
		log.Fatalf("Error connecting to db:", err)
	}

	defer db.Close()

	e := echo.New()

	e.Use(middleware.Recover())
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			helper.Logging(c).Info("Calling Api")

			return next(c)
		}
	})

	userhandler := &handler.UserHandler{UR: db}

	e.POST("/login", userhandler.Login)
	e.POST("/register", userhandler.Register)
	e.GET("/users", userhandler.GetAllUser)
	e.PUT("/user", userhandler.EditUser)
	e.DELETE("/user", userhandler.DeleteUser)

	e.Logger.Fatal(e.Start(":" + helper.PORT))
}
