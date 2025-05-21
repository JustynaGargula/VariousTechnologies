package database

import (
	"Zadanie4/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB
var DB2 *gorm.DB
var DB3 *gorm.DB

func InitDB() {
	var err error

	// Products
	DB, err = gorm.Open(sqlite.Open("file:database/products.db?_pragma=foreign_keys(ON)"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	DB.AutoMigrate(&models.Product{})

	// Carts
	DB2, err = gorm.Open(sqlite.Open("file:database/carts.db?_pragma=foreign_keys(ON)"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	DB2.AutoMigrate(&models.Cart{}, &models.CartItem{})

	// Users
	DB3, err = gorm.Open(sqlite.Open("file:database/users.db"), &gorm.Config{})
	if err != nil {
		log.Fatal("Failed to connect database:", err)
	}
	DB3.AutoMigrate(&models.User{})
}
