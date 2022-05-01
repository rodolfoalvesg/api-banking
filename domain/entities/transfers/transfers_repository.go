package transfers

import (
	"context"

	"github.com/google/uuid"
)

type TransferRepository interface {
	SaveTransfer(context.Context, Transfer) (uuid.UUID, error)
	ListAllTransfers(context.Context, string) ([]Transfer, error)
}
