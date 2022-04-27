package transfer

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

//CreateTransfer, caso de uso para criação e registro de uma transferência
func (u UsecaseTransfers) CreateTransfer(ctx context.Context, acc transfers.Transfer) (uuid.UUID, error) {
	transferID, err := u.repo.SaveTransfer(ctx, acc)
	if err != nil {
		return uuid.Nil, err
	}

	return transferID, nil
}
