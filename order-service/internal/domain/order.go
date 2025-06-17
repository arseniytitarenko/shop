package domain

import "github.com/google/uuid"

type Order struct {
	OrderID     uuid.UUID `gorm:"primaryKey"`
	UserID      uuid.UUID `gorm:"index"`
	Amount      uint
	Description string
	Status      Status
}

type Status string

const (
	StatusNew       Status = "NEW"
	StatusFinished  Status = "FINISHED"
	StatusCancelled Status = "CANCELLED"
)

func IsValidStatus(s Status) bool {
	switch s {
	case StatusNew, StatusFinished, StatusCancelled:
		return true
	default:
		return false
	}
}
