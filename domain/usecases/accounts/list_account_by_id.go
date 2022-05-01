package account

import (
	"context"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

// GetAccount, obtém dados de uma conta pelo ID da conta
func (u Usecase) GetAccount(ctx context.Context, accID string) (accounts.Account, error) {

	account, err := u.repo.ListAccountByID(ctx, accID)
	if err != nil {
		return accounts.Account{}, err
	}

	return account, nil

}
