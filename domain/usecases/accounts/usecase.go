package account

import (
	"context"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/domain/models"
)

//UsecaseToAccount, métodos das entidades
type UsecaseToAccount interface {
	CreateAccount(ctx context.Context, account models.Account) (models.Account, error)
	ShowBalanceByID(ctx context.Context, accountID string) (int, error)
	ShowListAccounts(ctx context.Context) ([]models.Account, error)
}

// UsecaseAccount, struct com campo de metódos das entidades
type UsecaseAccount struct {
	repository accounts.AccountRepository
}

//NewUsecaseAccount, cria repositório com metódos das entidades
func NewUsecaseAccount(acc accounts.AccountRepository) *UsecaseAccount {
	return &UsecaseAccount{
		repository: acc,
	}
}
