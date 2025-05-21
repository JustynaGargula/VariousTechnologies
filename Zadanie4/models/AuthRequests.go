package models

import "gorm.io/gorm"

type LoginRequest struct {
	Name     string
	Password string
}

type SignUpRequest struct {
	Name     string
	Email    string
	Password string
}
type User struct {
	gorm.Model
	Name           string
	Email          string
	HashedPassword string
}
