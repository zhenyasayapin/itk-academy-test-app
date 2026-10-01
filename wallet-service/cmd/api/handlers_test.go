package main

import (
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestUpdateWallet(t *testing.T) {
	tests := []struct {
		name       string
		payload    string
		statusCode int
	}{
		{
			name:       "Valid DEPOSIT request",
			payload:    `{"wallet_id": "123e4567-e89b-12d3-a456-426614174000", "amount": 10.5, "operation_type": "DEPOSIT"}`,
			statusCode: http.StatusOK,
		},
		{
			name:       "Valid WITHDRAW request",
			payload:    `{"wallet_id": "123e4567-e89b-12d3-a456-426614174000", "amount": 10.5, "operation_type": "WITHDRAW"}`,
			statusCode: http.StatusOK,
		},
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

	app := &Config{}
	mux := app.routes()

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
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
