package account

import (
	"context"
	"fmt"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

func (u Usecase) ShowAccounts(ctx context.Context) ([]accounts.Account, error) {
	accountLits, err := u.repo.ShowAccounts(ctx)
	if err != nil {
		return []accounts.Account{}, fmt.Errorf("Error showing accounts: %v", err)
	}
	return accountLits, nil
}
