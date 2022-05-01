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
			AccountDestinationID: uuid.New().String(),
			Amount:               2500,
		}

		err := ValidateTransferData(TestCaseA)

		if err != nil {
			t.Errorf("got %v, want %v", err, nil)
		}
	})

	t.Run("Lengeth", func(t *testing.T) {

		TestCaseA := Transfer{
			AccountDestinationID: "xxxx-xxxxxxxx",
			Amount:               2500,
		}

		err := ValidateTransferData(TestCaseA)

		if !errors.Is(err, ErrInvalidLengthID) {
			t.Errorf("got %v, want %v", err, ErrInvalidLengthID)
		}
	})

	t.Run("Invalid Amount", func(t *testing.T) {
		TestCaseB := Transfer{
			AccountDestinationID: uuid.New().String(),
			Amount:               0,
		}

		err := ValidateTransferData(TestCaseB)

		if !errors.Is(err, ErrInvalidAmount) {
			t.Errorf("want %v, got %v", ErrInvalidAmount, err)
		}
	})

	t.Run("Conflict ID", func(t *testing.T) {
		accID := uuid.New().String()
		TestCaseC := Transfer{
			AccountOriginID:      accID,
			AccountDestinationID: accID,
			Amount:               2500,
		}

		err := ValidateTransferData(TestCaseC)

		if !errors.Is(err, ErrEqualAccounts) {
			t.Errorf("want %v, got %v", ErrEqualAccounts, err)
		}
	})
}
