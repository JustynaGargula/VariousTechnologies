package controllers

import (
	"net/http"
	"strconv"

	"github.com/labstack/echo/v4"
	"gorm.io/gorm"
)

var products = []Product{
    {Id: "1", Name: "Produkt 1", Price: 100},
    {Id: "2", Name: "Produkt 2", Price: 200},
}
var lastUsedId = products[len(products)-1].Id

type Product struct {
	gorm.Model
	Id    string
	Name  string `json:"name"`
	Price uint `json:"price"`
}

func getNewId() string {
	num, _ := strconv.Atoi(lastUsedId)
	num = num+1
	lastUsedId = strconv.Itoa(num)
	return lastUsedId
}

func CreateProduct(c echo.Context) error {
	var newProduct Product
	bindError := c.Bind(&newProduct)
	if bindError == nil {
		newProduct.Id = getNewId()
		products = append(products, newProduct)
		return c.JSON(http.StatusOK, "Created product with id: "+newProduct.Id)
	} else{
		return c.JSON(http.StatusBadRequest, bindError)
	}
	
}

func GetProduct(c echo.Context) error {
	id := c.Param("id")
	for _, product := range products {
		if product.Id == id {
			return c.JSON(http.StatusOK, product)
		}
	}
	return c.String(http.StatusNotFound, "Product not found")
}

func GetProducts(c echo.Context) error {
  return c.JSON(http.StatusOK, products)
}

func UpdateProduct(c echo.Context) error {
	id := c.Param("id")
	var updatedProduct Product
	bindError := c.Bind(&updatedProduct)
	if bindError == nil {
		updatedProduct.Id = id

		for index, product := range products {
			if product.Id == id {
				products[index] = updatedProduct
				return c.JSON(http.StatusOK, "Updated product with id: "+updatedProduct.Id)
			}
		}
		return c.String(http.StatusNotFound, "Product not found")
	} else{
		return c.JSON(http.StatusBadRequest, bindError)
	}
}

func DeleteProduct (c echo.Context) error {
	id := c.Param("id")
	for index, product := range products {
		if product.Id == id {
			products = append(products[:index], products[index+1:]...)
			return c.String(http.StatusOK, "Deleted product with id: "+id)
		}
	}
	return c.String(http.StatusNotFound, "Couldn't find product with id: "+id)
}
