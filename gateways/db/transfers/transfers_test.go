package transfer_db

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

func TestSaveTransfer(t *testing.T) {
	t.Parallel()

	repository := NewRepositoryTransfer()

	myTransferFake := transfers.Transfer{
		Account_origin_ID:      uuid.New().String(),
		Account_destination_ID: uuid.New().String(),
		Amount:                 25400,
	}

	got, _ := repository.SaveTransfer(context.Background(), myTransferFake)
	if got == (uuid.UUID{}) {
		t.Errorf("got %v, want !=  %v", got, uuid.UUID{})
	}
}

func TestListAllTransfers(t *testing.T) {
	t.Parallel()
	repository := NewRepositoryTransfer()

	t.Run("Accounts Listed", func(t *testing.T) {
		myTransferFake := transfers.Transfer{
			Account_origin_ID:      uuid.New().String(),
			Account_destination_ID: uuid.New().String(),
			Amount:                 25400,
		}

		_, err := repository.SaveTransfer(context.Background(), myTransferFake)
		if err != nil {
			t.Errorf("Save Transfer error")
		}

		got, _ := repository.ListAllTransfers(context.Background(), myTransferFake.Account_origin_ID)
		want := 1
		if !reflect.DeepEqual(len(got), want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})

}
