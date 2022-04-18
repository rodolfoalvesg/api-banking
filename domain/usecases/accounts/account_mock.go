package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

type RepositoryMock interface {
	CreateAccount(context.Context, accounts.Account) (uuid.UUID, error)
	ShowBalance(ctx context.Context, accID uuid.UUID) (int, error)
	ShowAccounts(ctx context.Context) ([]accounts.Account, error)
	NewLogin(ctx context.Context, l Login) (string, error)
}

type UseCaseMock struct {
	SaveAccount       func(accounts.Account) (uuid.UUID, error)
	ListAllAccounts   func(context.Context) ([]accounts.Account, error)
	ListBalanceByID   func(uuid.UUID) (int, error)
	ListAccountsByCPF func(accCPF string) (accounts.Account, error)
}

func (m *UseCaseMock) CreateAccount(ctx context.Context, acc accounts.Account) (uuid.UUID, error) {
	return m.SaveAccount(acc)
}

func (m *UseCaseMock) ShowBalance(ctx context.Context, accID uuid.UUID) (int, error) {
	return m.ListBalanceByID(accID)
}

func (m *UseCaseMock) ShowAccounts(ctx context.Context) ([]accounts.Account, error) {
	return m.ListAllAccounts(ctx)
}

func (m *UseCaseMock) NewLogin(ctx context.Context, l Login) (accounts.Account, error) {
	return m.ListAccountsByCPF(l.CPF)
}
