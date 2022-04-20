package transfers

import (
	"errors"
	"time"
)

var (
	errInvalidLengthID = errors.New("Empty ID not allowed")
	errInvalidAmount   = errors.New("Invalid amount, be must > 0")
)

type Transfer struct {
	ID                     string    `json:"id"`
	Account_origin_ID      string    `json:"account_origin_id"`
	Account_destination_ID string    `json:"account_destination_id"`
	Amount                 uint      `json:"amount"`
	Created_At             time.Time `json:"created_at"`
}

func ValidateTransferData(accID string, amount uint) error {
	if len(accID) == 0 {
		return errInvalidLengthID
	}

	if amount <= 0 {
		return errInvalidAmount
	}

	return nil
}
