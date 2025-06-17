package domain

import (
	"github.com/google/uuid"
	"time"
)

type Outbox struct {
	ID          uuid.UUID `gorm:"primaryKey"`
	Type        string
	Payload     string
	CreatedAt   time.Time
	ProcessedAt *time.Time
}
