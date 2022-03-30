package account

import (
	"context"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/domain/models"
)

func (u UsecaseAccount) CreateAccount(ctx context.Context, account models.Account) (models.Account, error) {
	_, err := accounts.NewCreateAccount(account)

	if err != nil {
		return models.Account{}, err
	}

	accCreated, err := u.repository.AddedAccount()

	if err != nil {
		return models.Account{}, err
	}

	// // modelAccount := &db.FieldsToMethodsDB{
	// // 	Accounts: acc,
	// // }

	// // accCreated, err := modelAccount.AddedAccount()
	// // if err != nil {
	// // 	return models.Account{}, err
	// // }

	return accCreated, nil
}
