package main

import (
	"net/http"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)
  
type Product struct {
	gorm.Model
	id  string
	price uint
	name string
  }

func main() {
	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	e.POST("/users", saveUser)
	e.GET("/users/:id", getUser)

	e.Logger.Fatal(e.Start(":1323"))

}

// e.GET("/users/:id", getUser)
func getUser(c echo.Context) error {
	// User ID from path `users/:id`
	id := c.Param("id")
  return c.String(http.StatusOK, id)
}

func saveUser(c echo.Context) error {
	return c.String(http.StatusOK, "User may be saved")
}