package controllers

import (
	"Zadanie4/database"
	"Zadanie4/models"
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/golang-jwt/jwt"
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
	RedirectURL:  "http://localhost:1323/google-login-callback",
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.profile",
		"https://www.googleapis.com/auth/userinfo.email"},
	Endpoint: google.Endpoint,
}

var oauthStateString = "random"

func GoogleLoginHandler(w http.ResponseWriter, r *http.Request) {
	url := oauthConf.AuthCodeURL(oauthStateString)
	http.Redirect(w, r, url, http.StatusTemporaryRedirect)
}

func GoogleLoginCallbackHandler(w http.ResponseWriter, r *http.Request) {
	if r.FormValue("state") != oauthStateString {
		http.Error(w, "Invalid state", http.StatusBadRequest)
		return
	}

	code := r.FormValue("code")
	token, err := oauthConf.Exchange(context.Background(), code)
	if err != nil {
		http.Error(w, "Code exchange failed", http.StatusInternalServerError)
		return
	}

	client := oauthConf.Client(context.Background(), token)
	resp, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		http.Error(w, "Didn't manage to get user information", http.StatusInternalServerError)
		return
	}
	defer resp.Body.Close()

	var userinfo map[string]interface{}
	json.NewDecoder(resp.Body).Decode(&userinfo)

	//TODO dodać użytkownika do bazy danych

	var jwtSecret = []byte(os.Getenv("JWT_SECRET"))
	claims := jwt.MapClaims{
		"email": userinfo["email"].(string),
		"exp":   time.Now().Add(24 * time.Hour).Unix(),
		"iat":   time.Now().Unix(),
	}

	my_token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	jwtToken, err := my_token.SignedString(jwtSecret)
	if err != nil {
		http.Error(w, "Didn't manage to generate token", http.StatusInternalServerError)
		return
	}
	redirectURL := fmt.Sprintf("http://localhost:3000?token=%s", jwtToken)
	http.Redirect(w, r, redirectURL, http.StatusTemporaryRedirect)
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
