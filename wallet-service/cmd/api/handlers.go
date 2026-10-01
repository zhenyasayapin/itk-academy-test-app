package main

import (
	"errors"
	"net/http"

	"github.com/go-playground/validator/v10"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type UpdateWalletRequest struct {
	WalletID      string `json:"wallet_id" validate:"required,uuid"`
	Amount        float32 `json:"amount" validate:"required,gte=0"`
	OperationType string `json:"operation_type" validate:"required,oneof=DEPOSIT WITHDRAW"`
}

func (app *Config) UpdateWallet(w http.ResponseWriter, r *http.Request) {
	var request UpdateWalletRequest
	err := app.readJSON(w, r, &request)
	if err != nil {
		app.errorJSON(w, errors.New("Failed to read JSON"), http.StatusBadRequest)
		return
	}

	validate := validator.New()

	if err := validate.Struct(request); err != nil {
		app.errorJSON(w, errors.New("Validation error"), http.StatusBadRequest)
		return
	}

	var response JsonResponse
	response.Error = false
	response.Message = "The wallet was updated successfully!"

	app.writeJSON(w, http.StatusOK, response)
}
