package model

import (
	"github.com/google/uuid"
)

type Wallet struct {
	ID     uuid.UUID `json:"id"`
	Amount float32   `json:"amount"`
}