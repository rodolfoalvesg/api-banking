package transfer

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

func (u UsecaseTransfers) CreateTransfer(ctx context.Context, acc transfers.Transfer) (uuid.UUID, error) {
	transferID, err := u.repo.SaveTransfer(ctx, acc)
	if err != nil {
		return uuid.UUID{}, err
	}

	return transferID, nil
}
