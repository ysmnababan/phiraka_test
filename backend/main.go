package main

import (
	"backend/config"
	"backend/handler"
	"backend/helper"
	"log"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
)

func init() {
	helper.LoadEnv()
}

func main() {
	db, err := config.Connect()
	if err != nil {
		log.Fatalf("Error connecting to db: %v", err)
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
	// Configure CORS middleware
	e.Use(middleware.CORSWithConfig(middleware.CORSConfig{
		AllowOrigins: []string{"http://127.0.0.1:5500"}, // Replace with your frontend URL
		AllowMethods: []string{echo.GET, echo.POST, echo.PUT, echo.DELETE},
		AllowHeaders: []string{echo.HeaderOrigin, echo.HeaderContentType, echo.HeaderAccept},
	}))

	userhandler := &handler.UserHandler{UR: db}

	e.GET("/captcha", handler.GetCaptcha)
	e.POST("/login", userhandler.Login)
	e.POST("/register", userhandler.Register)
	e.GET("/users", userhandler.GetAllUser)
	e.PUT("/user", userhandler.EditUser)
	e.DELETE("/user", userhandler.DeleteUser)

	e.Logger.Fatal(e.Start(":" + helper.PORT))
}
