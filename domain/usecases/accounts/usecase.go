package account

import (
	"context"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/domain/models"
)

type UsecaseAccount struct {
	repository accounts.AccountRepository
}

type UsecaseToAccount interface {
	CreateAccount(ctx context.Context, account []models.Account) ([]models.Account, error)
	ShowBalanceID(ctx context.Context, accountID string) (int, error)
	ShowListAccounts(ctx context.Context) ([]models.Account, error)
}

func NewUsecaseAccount(acc accounts.AccountRepository) *UsecaseAccount {
	return &UsecaseAccount{
		repository: acc,
	}
}
