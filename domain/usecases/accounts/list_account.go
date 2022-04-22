package account

import (
	"context"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

func (u Usecase) GetAccount(ctx context.Context, accID string) (accounts.Account, error) {

	account, err := u.repo.ListAccountByID(ctx, accID)
	if err != nil {
		return account, err
	}

	return account, nil

}
