package controllers

import (
	"Zadanie4/models"
	"net/http"

	"github.com/labstack/echo/v4"
)

func LoginHandler(c echo.Context) error{
	var req models.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	if req.Name !="" && req.Password != "" {
		return c.JSON(http.StatusOK, echo.Map{"message":"Successfully logged in", "token": "fake-jwt-token"})
	} else {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid login credentials"})
	}

}