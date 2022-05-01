package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/common/security"
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
	VerifyAccount(context.Context, string) error
}

type UseCaseMock struct {
	SaveAccount        func(accounts.Account) (uuid.UUID, error)
	ListAllAccounts    func(context.Context) ([]accounts.Account, error)
	ListBalanceByID    func(uuid.UUID) (int, error)
	ListAccountsByCPF  func(string) (accounts.Account, error)
	ListAccountByID    func(string) (accounts.Account, error)
	UpdatedAccount     func(accounts.Balance) error
	VerifyAccountByCPF func(string) error
}

// CreateAccount, mock caso de uso para criação de conta
func (m *UseCaseMock) CreateAccount(ctx context.Context, acc accounts.Account) (uuid.UUID, error) {
	return m.SaveAccount(acc)
}

// ShowBalance, mock caso de uso para exibir saldo mediante ID do usuário
func (m *UseCaseMock) ShowBalance(ctx context.Context, accID uuid.UUID) (int, error) {
	return m.ListBalanceByID(accID)
}

// ShowAccounts, mock caso de uso para listar todas as contas cadastradas
func (m *UseCaseMock) ShowAccounts(ctx context.Context) ([]accounts.Account, error) {
	return m.ListAllAccounts(ctx)
}

// NewLogin, mock caso de uso para login na api e permissão para transferências
func (m *UseCaseMock) NewLogin(ctx context.Context, l Login) (string, error) {
	acc, err := m.ListAccountsByCPF(l.CPF)
	if err != nil {
		return "", err
	}

	token, err := security.CreateToken(acc.ID)
	if err != nil {
		return "", ErrCreateToken
	}

	return token, nil
}

// GetAccount, mock caso de uso para checagem ou exibição de conta pelo ID do usuário
func (m *UseCaseMock) GetAccount(_ context.Context, accID string) (accounts.Account, error) {
	return m.ListAccountByID(accID)
}

// UpdateAccount, mock caso de uso para atualizar saldo de contas
func (m *UseCaseMock) UpdateAccount(_ context.Context, transfer transfers.Transfer) error {
	return m.UpdatedAccount(accounts.Balance{})
}

// VerifyAccount, mock caso de uso para verificar a existência de conta
func (m *UseCaseMock) VerifyAccount(_ context.Context, accCPF string) error {
	return m.VerifyAccountByCPF(accCPF)
}
