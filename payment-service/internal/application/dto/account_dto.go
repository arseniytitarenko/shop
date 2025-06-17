package dto

import "github.com/google/uuid"

type AccountRequest struct {
	UserID uuid.UUID `json:"user_id"`
}

type ReplenishAccountRequest struct {
	UserID uuid.UUID `json:"user_id"`
	Amount uint      `json:"amount"`
}

type AccountResponse struct {
	UserID  uuid.UUID `json:"user_id"`
	Balance uint      `json:"balance"`
}
