package main

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	DB *sql.DB
}

func (m *PostgresDBRepo) Connection() *sql.DB {
	return m.DB
}

func (m *PostgresDBRepo) GetWallet(id uuid.UUID) (*Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
			w.id, w.amount
		from 
			wallets w
		where 
		    w.id = $1`

	var wallet Wallet
	row := m.DB.QueryRowContext(ctx, query, id)

	err := row.Scan(
		&wallet.ID,
		&wallet.Amount,
	)

	if err != nil {
		return nil, err
	}

	return &wallet, nil
}
