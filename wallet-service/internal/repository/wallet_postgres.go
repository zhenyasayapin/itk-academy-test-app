package repository

import (
	"context"
	"database/sql"
	"time"

	"github.com/google/uuid"
	"wallet-service/internal/model"
)

const dbTimeout = time.Second * 3

type PostgresDBRepo struct {
	DB *sql.DB
}

func (m *PostgresDBRepo) Connection() interface{} {
	return m.DB
}

func (m *PostgresDBRepo) GetWallet(id uuid.UUID) (*model.Wallet, error) {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		select 
			w.id, w.amount
		from 
			wallets w
		where 
		    w.id = $1`

	var wallet model.Wallet
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

func (m *PostgresDBRepo) Deposit(id uuid.UUID, amount float32) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		update wallets
		set amount = amount + $1
		where id = $2`

	_, err := m.DB.ExecContext(ctx, query, amount, id)
	return err
}

func (m *PostgresDBRepo) Withdraw(id uuid.UUID, amount float32) error {
	ctx, cancel := context.WithTimeout(context.Background(), dbTimeout)
	defer cancel()

	query := `
		update wallets
		set amount = amount - $1
		where id = $2`

	_, err := m.DB.ExecContext(ctx, query, amount, id)
	return err
}