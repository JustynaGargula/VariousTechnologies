package main

import (
	"Zadanie4/controllers"
	"Zadanie4/database"

	"net/http"

	"github.com/labstack/echo/v4"
)
  
func main() {
	database.InitDB()

	e := echo.New()
	e.GET("/", func(c echo.Context) error {
		return c.String(http.StatusOK, "Hello, World!")
	})

	// Products
	e.POST("/products", controllers.CreateProduct)
	e.GET("/products", controllers.GetProducts)
	e.GET("/products/:id", controllers.GetProduct)
	e.PUT("/products/:id", controllers.UpdateProduct)
	e.DELETE("/products/:id", controllers.DeleteProduct)

	e.Logger.Fatal(e.Start(":1323"))

}
