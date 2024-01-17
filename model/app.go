package model

import (
	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

type App struct {
	DB   *gorm.DB
	Echo *echo.Echo
}
