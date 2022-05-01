package transfer_db

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

//TestSaveTransfer, teste de método para criação e registro de uma transferência no db
func TestSaveTransfer(t *testing.T) {

	type TestCase struct {
		name           string
		myTransferFake transfers.Transfer
		sourceTest     string
		err            error
	}

	testCase := []TestCase{
		{
			name: "Save the transfer successfully",
			myTransferFake: transfers.Transfer{
				ID:                   uuid.New().String(),
				AccountOriginID:      uuid.New().String(),
				AccountDestinationID: uuid.New().String(),
				Amount:               25400,
				CreatedAt:            time.Now().UTC(),
			},
		},
		{
			name: "Fail if empty transfer id",
			myTransferFake: transfers.Transfer{
				AccountOriginID:      uuid.New().String(),
				AccountDestinationID: uuid.New().String(),
				Amount:               25400,
				CreatedAt:            time.Now().UTC(),
			},
			err: nil,
		},
	}

	for _, tc := range testCase {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			repository := NewRepositoryTransfer()
			_, err := repository.SaveTransfer(context.Background(), tt.myTransferFake)
			if !errors.Is(tt.err, err) {
				t.Errorf("Expected %s, got %s", tt.err, err)
			}
		})
	}
}

//TestListAllTransfers, teste de método exibição de todas as transferências de um usuário
func TestListAllTransfers(t *testing.T) {
	t.Parallel()
	repository := NewRepositoryTransfer()

	t.Run("Accounts Listed", func(t *testing.T) {
		myTransferFake := transfers.Transfer{
			AccountOriginID:      uuid.New().String(),
			AccountDestinationID: uuid.New().String(),
			Amount:               25400,
		}

		_, err := repository.SaveTransfer(context.Background(), myTransferFake)
		if err != nil {
			t.Errorf("Save Transfer error")
		}

		got, _ := repository.ListAllTransfers(context.Background(), myTransferFake.AccountOriginID)
		want := 1
		if !reflect.DeepEqual(len(got), want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
