package main

import (
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"github.com/lsig/OtpWebAPI/database"
	"github.com/lsig/OtpWebAPI/handler"
	"log"
)

func main() {
	err := database.ConnectDatabase()

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

	err = database.Migrate()

	if err != nil {
		log.Fatalf("Server failed: %v", err)
	}

	e := echo.New()
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())

	e.GET("/users/:id/otp", handler.GetOTP)
	e.POST("/users", handler.CreateUser)

	e.Logger.Fatal(e.Start(":5050"))
}
