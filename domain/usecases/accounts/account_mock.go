package account

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

type RepositoryMock interface {
	CreateAccount(context.Context, accounts.Account) (uuid.UUID, error)
	ShowBalance(ctx context.Context, accID uuid.UUID) (int, error)
	ShowAccounts(ctx context.Context) ([]accounts.Account, error)
}

type UseCaseMock struct {
	SaveAccount     func(accounts.Account) (uuid.UUID, error)
	ListAllAccounts func(context.Context) ([]accounts.Account, error)
	ListBalanceByID func(context.Context, uuid.UUID) (int, error)
}

func (m *UseCaseMock) CreateAccount(ctx context.Context, acc accounts.Account) (uuid.UUID, error) {
	fmt.Println("Aqui mock")
	return m.SaveAccount(acc)
}

func (m *UseCaseMock) ShowBalance(ctx context.Context, accID uuid.UUID) (int, error) {
	return m.ListBalanceByID(ctx, accID)
}

func (m *UseCaseMock) ShowAccounts(ctx context.Context) ([]accounts.Account, error) {
	return m.ListAllAccounts(ctx)
}
