package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
)

type Controllers interface {
	CreateAccount(context.Context, accounts.Account) (uuid.UUID, error)
	ShowAccounts(context.Context) ([]accounts.Account, error)
	ShowBalance(context.Context, uuid.UUID) (int, error)
	NewLogin(context.Context, account.Login) (string, error)
}

type Controller struct {
	account Controllers
}

func NewController(c Controllers) *Controller {
	return &Controller{
		account: c,
	}
}
