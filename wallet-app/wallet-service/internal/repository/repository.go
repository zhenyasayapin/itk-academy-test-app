package repository

import (
	"github.com/google/uuid"
	"wallet-service/internal/model"
)

type DatabaseRepo interface {
	Connection() interface{}
	GetWallet(uuid uuid.UUID) (*model.Wallet, error)
	Deposit(uuid uuid.UUID, amount float32) error
	Withdraw(uuid uuid.UUID, amount float32) error
}