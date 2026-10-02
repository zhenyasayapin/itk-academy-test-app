package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
)

func TestUpdateWalletValidation(t *testing.T) {
	tests := []struct {
		name       string
		payload    string
		statusCode int
	}{
		{
			name:       "Missing wallet_id",
			payload:    `{"amount": 10.5, "operation_type": "DEPOSIT"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Missing amount",
			payload:    `{"wallet_id": "123e4567-e89b-12d3-a456-426614174000", "operation_type": "DEPOSIT"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Invalid wallet_id format",
			payload:    `{"wallet_id": "not-a-uuid", "amount": 10.5, "operation_type": "DEPOSIT"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Negative amount",
			payload:    `{"wallet_id": "123e4567-e89b-12d3-a456-426614174000", "amount": -5, "operation_type": "DEPOSIT"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Invalid JSON",
			payload:    `{invalid json}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Empty body",
			payload:    ``,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Zero amount fails required validation",
			payload:    `{"wallet_id": "123e4567-e89b-12d3-a456-426614174000", "amount": 0, "operation_type": "DEPOSIT"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Invalid operation_type",
			payload:    `{"wallet_id": "123e4567-e89b-12d3-a456-426614174000", "amount": 10.5, "operation_type": "TRANSFER"}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Missing operation_type",
			payload:    `{"wallet_id": "123e4567-e89b-12d3-a456-426614174000", "amount": 10.5}`,
			statusCode: http.StatusBadRequest,
		},
		{
			name:       "Empty operation_type",
			payload:    `{"wallet_id": "123e4567-e89b-12d3-a456-426614174000", "amount": 10.5, "operation_type": ""}`,
			statusCode: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			app := &Config{}
			mux := app.routes()

			req, err := http.NewRequest(http.MethodPost, "/api/v1/wallet", bytes.NewBufferString(tt.payload))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			if rr.Code != tt.statusCode {
				t.Errorf("expected status %d, got %d", tt.statusCode, rr.Code)
			}
		})
	}
}

func TestUpdateWalletWithRepository(t *testing.T) {
	walletID := uuid.MustParse("123e4567-e89b-12d3-a456-426614174000")

	tests := []struct {
		name           string
		payload        string
		statusCode     int
		walletAmount   float32
		expectedAmount float32
	}{
		{
			name:           "Valid DEPOSIT request",
			payload:        `{"wallet_id": "123e4567-e89b-12d3-a456-426614174000", "amount": 100.5, "operation_type": "DEPOSIT"}`,
			statusCode:     http.StatusOK,
			walletAmount:   0,
			expectedAmount: 100.5,
		},
		{
			name:           "Valid WITHDRAW request with sufficient balance",
			payload:        `{"wallet_id": "123e4567-e89b-12d3-a456-426614174000", "amount": 30.0, "operation_type": "WITHDRAW"}`,
			statusCode:     http.StatusOK,
			walletAmount:   100.5,
			expectedAmount: 70.5,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockRepo := NewMockDBRepo()
			mockRepo.Wallets[walletID] = &Wallet{ID: walletID, Amount: tt.walletAmount}

			app := &Config{DB: mockRepo}
			mux := app.routes()

			req, err := http.NewRequest(http.MethodPost, "/api/v1/wallet", bytes.NewBufferString(tt.payload))
			if err != nil {
				t.Fatal(err)
			}

			rr := httptest.NewRecorder()
			mux.ServeHTTP(rr, req)

			if rr.Code != tt.statusCode {
				t.Errorf("expected status %d, got %d", tt.statusCode, rr.Code)
			}

			var response JsonResponse
			err = json.Unmarshal(rr.Body.Bytes(), &response)
			if err != nil {
				t.Fatalf("failed to unmarshal response: %v", err)
			}

			if response.Error {
				t.Errorf("expected no error, got error: %s", response.Message)
			}

			actualAmount := mockRepo.Wallets[walletID].Amount
			if actualAmount != tt.expectedAmount {
				t.Errorf("expected wallet amount %f, got %f", tt.expectedAmount, actualAmount)
			}
		})
	}
}