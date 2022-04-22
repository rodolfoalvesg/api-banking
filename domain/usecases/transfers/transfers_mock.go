package transfer

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

type RepositoryTransferMock interface {
	CreateTransfer(context.Context, transfers.Transfer) (uuid.UUID, error)
	ShowTransfers(context.Context, string) ([]transfers.Transfer, error)
}

type UseCaseTransferMock struct {
	SaveTransfer     func(transfers.Transfer) (uuid.UUID, error)
	ListAllTransfers func(string) ([]transfers.Transfer, error)
}

func (m *UseCaseTransferMock) CreateTransfer(ctx context.Context, t transfers.Transfer) (uuid.UUID, error) {
	return m.SaveTransfer(t)
}

func (m *UseCaseTransferMock) ShowTransfers(ctx context.Context, accID string) ([]transfers.Transfer, error) {
	return m.ListAllTransfers(accID)
}
