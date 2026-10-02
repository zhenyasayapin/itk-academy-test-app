package main

import (
	"sync"

	"github.com/google/uuid"
)

type MockDBRepo struct {
	mu      sync.Mutex
	Wallets map[uuid.UUID]*Wallet
}

func NewMockDBRepo() *MockDBRepo {
	return &MockDBRepo{
		Wallets: make(map[uuid.UUID]*Wallet),
	}
}

func (m *MockDBRepo) GetWallet(id uuid.UUID) (*Wallet, error) {
	m.mu.Lock()
	defer m.mu.Unlock()

	if wallet, ok := m.Wallets[id]; ok {
		return wallet, nil
	}
	return nil, nil
}

func (m *MockDBRepo) Deposit(id uuid.UUID, amount float32) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if wallet, ok := m.Wallets[id]; ok {
		wallet.Amount += amount
	}
	return nil
}

func (m *MockDBRepo) Withdraw(id uuid.UUID, amount float32) error {
	m.mu.Lock()
	defer m.mu.Unlock()

	if wallet, ok := m.Wallets[id]; ok {
		wallet.Amount -= amount
	}
	return nil
}

func (m *MockDBRepo) Connection() interface{} {
	return nil
}