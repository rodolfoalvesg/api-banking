package account

import (
	"context"
	"fmt"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

// ShowAccounts, caso de uso para listar todas as contas
func (u Usecase) ShowAccounts(ctx context.Context) ([]accounts.Account, error) {
	accountList, err := u.repo.ListAllAccounts(ctx)
	if err != nil {
		return []accounts.Account{}, fmt.Errorf("Error showing accounts: %v", err)
	}
	return accountList, nil
}
