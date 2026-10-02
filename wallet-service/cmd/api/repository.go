package main

import (
	"github.com/google/uuid"
)

type DatabaseRepo interface {
	Connection() interface{}
	GetWallet(uuid uuid.UUID) (*Wallet, error)
	Deposit(uuid uuid.UUID, amount float32) error
	Withdraw(uuid uuid.UUID, amount float32) error
}
