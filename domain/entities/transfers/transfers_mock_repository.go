package transfers

import (
	"context"

	"github.com/google/uuid"
)

type TransferMock struct {
	onSaveTransfer    func(Transfer) (uuid.UUID, error)
	onListAllTransfer func(string) ([]Transfer, error)
}

func (m *TransferMock) SaveTransfer(ctx context.Context, t Transfer) (uuid.UUID, error) {
	return m.onSaveTransfer(t)
}

func (m *TransferMock) ListAllTransfers(ctx context.Context, accID string) ([]Transfer, error) {
	return m.onListAllTransfer(accID)
}
