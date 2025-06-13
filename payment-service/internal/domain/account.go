package domain

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Account struct {
	gorm.Model
	UserID  uuid.UUID `gorm:"primaryKey"`
	Balance uint
}
