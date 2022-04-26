package transfers

import (
	"errors"
	"testing"

	"github.com/google/uuid"
)

//TestValidateTransferData, teste de validação de informações da transferência
func TestValidateTransferData(t *testing.T) {
	t.Parallel()

	t.Run("No error", func(t *testing.T) {

		TestCaseA := Transfer{
			Account_destination_ID: uuid.New().String(),
			Amount:                 2500,
		}

		err := ValidateTransferData(&TestCaseA)

		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
	})

	t.Run("Invalid Amount", func(t *testing.T) {
		TestCaseB := Transfer{
			Account_destination_ID: uuid.New().String(),
			Amount:                 0,
		}

		err := ValidateTransferData(&TestCaseB)

		if !errors.Is(err, errInvalidAmount) {
			t.Errorf("want %v, got %v", errInvalidAmount, err)
		}
	})

	t.Run("Conflict ID", func(t *testing.T) {
		accID := uuid.New().String()
		TestCaseC := Transfer{
			Account_origin_ID:      accID,
			Account_destination_ID: accID,
			Amount:                 2500,
		}

		err := ValidateTransferData(&TestCaseC)

		if !errors.Is(err, errEqualAccounts) {
			t.Errorf("want %v, got %v", errEqualAccounts, err)
		}
	})
}
