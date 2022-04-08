package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

type Controllers interface {
	CreateAccount(ctx context.Context, account accounts.Account) (uuid.UUID, error)
	ShowAccounts(ctx context.Context) ([]accounts.Account, error)
	ShowBalance(ctx context.Context, accID string) (int, error)
}

type Controller struct {
	account Controllers
}

func NewController(c Controllers) *Controller {
	return &Controller{
		account: c,
	}
}
