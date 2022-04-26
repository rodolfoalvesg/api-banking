package account

import (
	"context"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

// ShowAccounts, caso de uso para listar todas as contas
func (u Usecase) ShowAccounts(ctx context.Context) ([]accounts.Account, error) {
	accountList, err := u.repo.ListAllAccounts(ctx)
	if err != nil {
		return []accounts.Account{}, err
	}
	return accountList, nil
}
