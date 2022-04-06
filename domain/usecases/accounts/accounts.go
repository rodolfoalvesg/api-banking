package account

import (
	"context"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/gateways/db"
)

//UsecaseToAccount, métodos das entidades
type UsecaseToAccount interface {
	CreateAccount(ctx context.Context, account accounts.Account) (db.Database, error)
}

// UsecaseAccount, struct com campo de metódos das entidades
type UsecaseAccount struct {
	db *db.Database
}

//NewUsecaseAccount, cria repositório com metódos das entidades
func NewUsecaseAccount(a *db.Database) *UsecaseAccount {
	return &UsecaseAccount{
		db: a,
	}
}

func (u UsecaseAccount) CreateAccount(ctx context.Context, account accounts.Account) ([]db.Database, error) {

	err := account.CreateNewAccount(ctx, account)
	if err != nil {
		return []db.Database{}, err
	}

	acc := &db.Database{
		ID:        uuid.New().String(),
		Name:      account.Name,
		CPF:       account.CPF,
		Secret:    account.Secret,
		Balance:   account.Balance,
		CreatedAt: time.Now(),
	}

	accCreated, err := acc.AddedAccount(ctx)
	if err != nil {
		return []db.Database{}, err
	}

	return accCreated, nil
}
