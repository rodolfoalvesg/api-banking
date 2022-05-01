package account

import (
	"context"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

// ShowAccounts, caso de uso para listar todas as contas
func (u Usecase) ShowAccounts(ctx context.Context) ([]accounts.Account, error) {
	var accountsList []accounts.Account

	accList, err := u.repo.ListAllAccounts(ctx)
	if err != nil {
		return nil, err
	}

	for _, account := range accList {
		accountsList = append(accountsList, accounts.Account{
			ID:        account.ID,
			Name:      account.Name,
			CPF:       account.CPF,
			Balance:   account.Balance,
			CreatedAt: account.CreatedAt,
		})
	}

	return accountsList, nil
}
