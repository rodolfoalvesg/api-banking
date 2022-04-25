package transfers

import (
	"context"

	"github.com/google/uuid"
)

type TransferMock struct {
	OnSaveTransfer    func(Transfer) (uuid.UUID, error)
	OnListAllTransfer func(string) ([]Transfer, error)
}

func (m TransferMock) SaveTransfer(ctx context.Context, t Transfer) (uuid.UUID, error) {
	return m.OnSaveTransfer(t)
}

func (m TransferMock) ListAllTransfers(ctx context.Context, accID string) ([]Transfer, error) {
	return m.OnListAllTransfer(accID)
}
