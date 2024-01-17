package handler

import (
	"errors"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/lsig/OtpWebAPI/database"
	"github.com/lsig/OtpWebAPI/model"
	"gorm.io/gorm"
)

func CreateUser(c echo.Context) error {
	db := database.GetDB()

	name := c.QueryParam("name")
	otp := c.QueryParam("otp")
	user := model.User{
		Name: name,
		OTP:  otp,
	}

	result := db.Create(&user)

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusCreated, user)

}

// /users/:id/otp
func GetOTP(c echo.Context) error {
	var user model.User
	db := database.GetDB()

	id := c.Param("id")
	result := db.Select("otp").First(&user, "id = $1", id)

	if errors.Is(result.Error, gorm.ErrRecordNotFound) {
		return c.JSON(http.StatusNotFound, echo.Map{"error": "User not found"})
	}

	if result.Error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": result.Error.Error()})
	}

	return c.JSON(http.StatusOK, echo.Map{"otp": user.OTP})
}
