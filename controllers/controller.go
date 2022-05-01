package controllers

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
)

type ControllersTransfers interface {
	CreateTransfer(context.Context, transfers.Transfer) (uuid.UUID, error)
	ShowTransfers(context.Context, string) ([]transfers.Transfer, error)
}

type ControllersAccount interface {
	CreateAccount(context.Context, accounts.Account) (uuid.UUID, error)
	ShowAccounts(context.Context) ([]accounts.Account, error)
	ShowBalance(context.Context, uuid.UUID) (int, error)
	NewLogin(context.Context, account.Login) (string, error)
	GetAccount(context.Context, string) (accounts.Account, error)
	UpdateAccount(context.Context, transfers.Transfer) error
	VerifyAccount(context.Context, string) error
}

type Controller struct {
	account  ControllersAccount
	transfer ControllersTransfers
}

func NewController(c ControllersAccount, t ControllersTransfers) *Controller {
	return &Controller{
		account:  c,
		transfer: t,
	}
}
