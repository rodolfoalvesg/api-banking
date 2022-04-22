package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

type Repository interface {
	SaveAccount(context.Context, accounts.Account) (uuid.UUID, error)
	ListAllAccounts(context.Context) ([]accounts.Account, error)
	ListBalanceByID(context.Context, uuid.UUID) (int, error)
	ListAccountsByCPF(context.Context, string) (accounts.Account, error)
	ListAccountByID(context.Context, string) (accounts.Account, error)
	UpdatedAccount(context.Context, accounts.Balance) error
}

type Usecase struct {
	repo Repository
}

//NewUsecaseAccount, cria repositório com metódos das entidades
func NewUsecase(repo Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}

}
