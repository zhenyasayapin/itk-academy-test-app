package main

import (
	"database/sql"

	"github.com/google/uuid"
)

type DatabaseRepo interface {
	Connection() *sql.DB
	GetWallet(uuid uuid.UUID) (*Wallet, error)
}
