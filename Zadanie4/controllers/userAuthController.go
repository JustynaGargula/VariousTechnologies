package controllers

import (
	"Zadanie4/database"
	"Zadanie4/models"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	"github.com/labstack/echo/v4"
	"golang.org/x/crypto/bcrypt"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
)

func init() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
}

var oauthConf = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_CLIENT_SECRET_KEY"),
	RedirectURL:  "http://localhost:3000",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: google.Endpoint,
}

var oauthStateString = "random"

func GoogleLoginHandler(c echo.Context, w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
	//TODO zrobić handleCallback i w ogóle skleić to
}

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
