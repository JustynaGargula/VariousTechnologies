package controllers

import (
	"Zadanie4/database"
	"Zadanie4/models"
	"net/http"

	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
)

func LoginHandler(c echo.Context) error {
	var req models.LoginRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	var storedUser models.User
	result := database.DB3.Where("name = ?", req.Name).First(&storedUser)
	if result.Error != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "No such user"})
	}
	hashedStoredPassword := storedUser.HashedPassword
	err = bcrypt.CompareHashAndPassword([]byte(hashedStoredPassword), []byte(req.Password))
	if err == nil {
		return c.JSON(http.StatusOK, echo.Map{"message": "Successfully logged in", "token": "fake-jwt-token"})
	} else {
		return c.JSON(http.StatusUnauthorized, echo.Map{"error": "Invalid login credentials"})
	}

}

func SignUpHandler(c echo.Context) error {
	var req models.SignUpRequest
	err := c.Bind(&req)
	if err != nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Invalid request"})
	}

	var otherUser models.User
	result := database.DB3.Where("email = ? OR name = ?", req.Email, req.Name).First(&otherUser)
	if result.Error == nil {
		return c.JSON(http.StatusBadRequest, echo.Map{"error": "Email or name already used"})
	}

	hashedPassword, error := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	if error != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to hash password"})
	} else {
		print(hashedPassword)
	}

	user := models.User{
		Name:           req.Name,
		Email:          req.Email,
		HashedPassword: string(hashedPassword),
	}

	err = database.DB3.Create(&user).Error
	if err != nil {
		return c.JSON(http.StatusInternalServerError, echo.Map{"error": "Failed to sign up user"})
	}

	return c.JSON(http.StatusOK, echo.Map{"message": "User signed up successfully"})
}
