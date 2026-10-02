package main

import (
	"errors"
	"log"
	"net/http"

	"github.com/go-playground/validator/v10"
	"github.com/google/uuid"
)

const (
	OperationDeposit  = "DEPOSIT"
	OperationWithdraw = "WITHDRAW"
)

type JsonResponse struct {
	Error   bool   `json:"error"`
	Message string `json:"message"`
	Data    any    `json:"data,omitempty"`
}

type UpdateWalletRequest struct {
	WalletID      uuid.UUID `json:"wallet_id" validate:"required,uuid"`
	Amount        float32   `json:"amount" validate:"required,gte=0"`
	OperationType string    `json:"operation_type" validate:"required,oneof=DEPOSIT WITHDRAW"`
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

	wallet, err := app.DB.GetWallet(request.WalletID)
	if err != nil {
		app.errorJSON(w, errors.New("Failed to retrieve the wallet"), http.StatusBadRequest)
		log.Println(err)
		return
	}

	switch request.OperationType {
	case OperationDeposit:
		err = app.DB.Deposit(request.WalletID, request.Amount)
	case OperationWithdraw:
		if wallet.Amount-request.Amount <= 0 {
			app.errorJSON(w, errors.New("Insufficient funds"), http.StatusBadRequest)
			return
		} else {
			err = app.DB.Withdraw(request.WalletID, request.Amount)
		}
	}

	if err != nil {
		app.errorJSON(w, errors.New("Failed to update wallet"), http.StatusBadRequest)
		log.Println(err)
		return
	}

	var response JsonResponse
	response.Error = false
	response.Message = "The wallet was updated successfully!"

	app.writeJSON(w, http.StatusOK, response)
}
