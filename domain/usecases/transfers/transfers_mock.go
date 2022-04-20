package transfer

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

type RepositoryTransferMock interface {
	SaveTransfer(context.Context, transfers.Transfer) (uuid.UUID, error)
}

type UseCaseTransferMock struct {
	onCheckAcccountID func(string) error
}

func (m *UseCaseTransferMock) CheckAccount(accID string) error {
	return m.onCheckAcccountID(accID)
}
