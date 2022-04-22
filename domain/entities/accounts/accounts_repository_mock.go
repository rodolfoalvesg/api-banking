package accounts

import (
	"context"

	"github.com/google/uuid"
)

type AccountMock struct {
	SaveAcc              func(Account) (uuid.UUID, error)
	ListAllAcc           func(context.Context) ([]Account, error)
	ListBalanceByIDAcc   func(uuid.UUID) (int, error)
	ListAccountsByCPFAcc func(string) (Account, error)
	ListAccountByIDacc   func(string) (Account, error)
	UpdatedAccountacc    func(Balance) error
}

func (m AccountMock) SaveAccount(ctx context.Context, acc Account) (uuid.UUID, error) {
	return m.SaveAcc(acc)
}

func (m AccountMock) ListBalanceByID(ctx context.Context, accID uuid.UUID) (int, error) {
	return m.ListBalanceByIDAcc(accID)
}

func (m AccountMock) ListAllAccounts(ctx context.Context) ([]Account, error) {
	return m.ListAllAcc(ctx)
}

func (m AccountMock) ListAccountsByCPF(ctx context.Context, accCPF string) (Account, error) {
	return m.ListAccountsByCPFAcc(accCPF)
}

func (m AccountMock) ListAccountByID(ctx context.Context, accID string) (Account, error) {
	return m.ListAccountByIDacc(accID)
}

func (m AccountMock) UpdatedAccount(ctx context.Context, b Balance) error {
	return m.UpdatedAccountacc(b)
}
