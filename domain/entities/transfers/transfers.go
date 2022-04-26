package transfers

import (
	"errors"
	"time"
)

var (
	errInvalidLengthID = errors.New("Empty ID not allowed")
	errInvalidAmount   = errors.New("Invalid amount, be must > 0")
	errEqualAccounts   = errors.New("Equal source and destination accounts are not allowed.")
)

type Transfer struct {
	ID                     string    `json:"id"`
	Account_origin_ID      string    `json:"account_origin_id"`
	Account_destination_ID string    `json:"account_destination_id"`
	Amount                 uint      `json:"amount"`
	Created_At             time.Time `json:"created_at"`
}

//ValidateTransferData, valida as informações na origem de uma transferência
func ValidateTransferData(transfer *Transfer) error {
	if len(transfer.Account_destination_ID) != 36 {
		return errInvalidLengthID
	}

	if transfer.Amount <= 0 {
		return errInvalidAmount
	}

	if transfer.Account_origin_ID == transfer.Account_destination_ID {
		return errEqualAccounts
	}

	return nil
}
