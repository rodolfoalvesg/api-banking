package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

type RepositoryMock interface {
	CreateAccount(context.Context, accounts.Account) (uuid.UUID, error)
	ShowBalance(context.Context, uuid.UUID) (int, error)
	ShowAccounts(ctx context.Context) ([]accounts.Account, error)
	NewLogin(context.Context, Login) (string, error)
	GetAccount(context.Context, string) (accounts.Account, error)
	UpdateAccount(context.Context, transfers.Transfer) error
}

type UseCaseMock struct {
	SaveAccount       func(accounts.Account) (uuid.UUID, error)
	ListAllAccounts   func(context.Context) ([]accounts.Account, error)
	ListBalanceByID   func(uuid.UUID) (int, error)
	ListAccountsByCPF func(string) (accounts.Account, error)
	ListAccountByID   func(string) (accounts.Account, error)
	UpdatedAccount    func(accounts.Balance) error
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

func (m *UseCaseMock) NewLogin(ctx context.Context, l Login) (string, error) {
	return "", nil
}

func (m *UseCaseMock) GetAccount(_ context.Context, accID string) (accounts.Account, error) {
	return m.ListAccountByID(accID)
}

func (m *UseCaseMock) UpdateAccount(_ context.Context, transfer transfers.Transfer) error {
	return m.UpdatedAccount(accounts.Balance{})
}
