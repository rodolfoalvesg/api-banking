package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

//CreateAccount, caso de uso para criação de conta
func (u Usecase) CreateAccount(ctx context.Context, account accounts.Account) (uuid.UUID, error) {

	acc, err := accounts.GeneratePasswdHash(ctx, account)
	if err != nil {
		return uuid.Nil, err
	}

	account.Secret = string(acc)

	accID, err := u.repo.SaveAccount(ctx, account)
	if err != nil {
		return uuid.Nil, err
	}

	return accID, nil
}
