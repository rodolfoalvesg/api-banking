package transfers

import (
	"context"

	"github.com/google/uuid"
)

type TransferMock struct {
	OnSaveTransfer    func(Transfer) (uuid.UUID, error)
	OnListAllTransfer func(string) ([]Transfer, error)
}

// SaveTransfer, mock para salvar transferência
func (m TransferMock) SaveTransfer(ctx context.Context, t Transfer) (uuid.UUID, error) {
	return m.OnSaveTransfer(t)
}

//ListAllTransfers, mock para listar transferências de um usuário
func (m TransferMock) ListAllTransfers(ctx context.Context, accID string) ([]Transfer, error) {
	return m.OnListAllTransfer(accID)
}
