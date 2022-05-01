package accounts

import (
	"context"

	"github.com/google/uuid"
)

type AccountRepository interface {
	SaveAccount(context.Context, Account) (uuid.UUID, error)
	ListAllAccounts(context.Context) ([]Account, error)
	ListBalanceByID(context.Context, uuid.UUID) (int, error)
	ListAccountsByCPF(context.Context, string) (Account, error)
	ListAccountByID(context.Context, string) (Account, error)
	UpdatedAccount(context.Context, Balance) error
}
