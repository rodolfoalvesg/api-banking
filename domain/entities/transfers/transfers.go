package transfers

import (
	"errors"
	"time"
)

var (
	ErrInvalidLengthID = errors.New("empty ID not allowed")
	ErrInvalidAmount   = errors.New("invalid amount, be must > 0")
	ErrEqualAccounts   = errors.New("equal source and destination accounts are not allowed.")
)

type Transfer struct {
	ID                   string    `json:"id"`
	AccountOriginID      string    `json:"account_origin_id"`
	AccountDestinationID string    `json:"account_destination_id"`
	Amount               uint      `json:"amount"`
	CreatedAt            time.Time `json:"created_at"`
}

//ValidateTransferData, valida as informações na origem de uma transferência
func ValidateTransferData(transfer Transfer) error {
	if len(transfer.AccountDestinationID) != 36 {
		return ErrInvalidLengthID
	}

	if transfer.Amount <= 0 {
		return ErrInvalidAmount
	}

	if transfer.AccountOriginID == transfer.AccountDestinationID {
		return ErrEqualAccounts
	}

	return nil
}
