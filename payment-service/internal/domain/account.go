package domain

import (
	"github.com/google/uuid"
)

type Account struct {
	UserID  uuid.UUID `gorm:"primaryKey"`
	Balance uint
}
