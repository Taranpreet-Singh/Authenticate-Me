package models

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type JwtToken struct {
	gorm.Model
	ID     uuid.UUID `gorm:"type:uuid;primaryKey"`
	Token  string    `gorm:"unique"`
	UserID uuid.UUID
	User   User `gorm:"foreignKey:UserID"`
}
