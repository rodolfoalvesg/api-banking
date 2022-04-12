package usecase

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

type Repository interface {
	SaveAccount(context.Context, accounts.Account) (uuid.UUID, error)
	ListAllAccounts(context.Context) ([]accounts.Account, error)
	ListBalanceByID(context.Context, uuid.UUID) (int, error)
}
