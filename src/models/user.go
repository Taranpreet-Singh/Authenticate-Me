package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	ID uuid.UUID `gorm:"type:uuid;primaryKey"`
	// Username string `gorm:"unique"`
	Email    string `gorm:"unique"`
	Password string
}

type UserRequestBody struct {
	Email    string
	Password string
}
