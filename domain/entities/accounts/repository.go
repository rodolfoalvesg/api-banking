package accounts

import (
	"context"

	"github.com/rodolfoalvesg/api-banking/api/domain/models"
)

// AccountRepository, repositórios da entidade accounts
type AccountRepository interface {
	CreateAccount(ctx context.Context, account models.Account) ([]models.Account, error)
	ShowBalanceByID(ctx context.Context, accountID string) (int, error)
	ShowListAccounts(ctx context.Context) ([]models.Account, error)
}
