package database

import (
	"Zadanie4/models"
	"log"

	"github.com/glebarez/sqlite"
	"gorm.io/gorm"
)

var DB *gorm.DB

func InitDB() {
    var err error
    DB, err = gorm.Open(sqlite.Open("file:database/products.db?_pragma=foreign_keys(ON)"), &gorm.Config{})
    if err != nil {
        log.Fatal("Failed to connect database:", err)
    }

    DB.AutoMigrate(&models.Product{})
}
