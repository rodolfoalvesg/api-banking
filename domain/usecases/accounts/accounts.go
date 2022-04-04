package account

import (
	"context"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/domain/models"
	"github.com/rodolfoalvesg/api-banking/api/gateways/db"
)

var database db.Database
var dba = db.NewRepositoryDB(database)

func (u UsecaseAccount) CreateAccount(ctx context.Context, account models.Account) (models.Account, error) {
	acc, err := accounts.CreateNewAccount(ctx, account)

	if err != nil {
		return models.Account{}, err
	}

	dba.Accounts = acc // Salva a conta no campo Accounts do dbFields

	accCreated, err := dba.AddedAccount(ctx)
	if err != nil {
		return models.Account{}, err
	}

	return accCreated, nil
}
