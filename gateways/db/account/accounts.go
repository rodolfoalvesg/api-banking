package account_db

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
)

var (
	ErrIDExists = errors.New("ID already exists!")
	ErrNotFound = errors.New("account not found")
	ErrInternal = errors.New("internal Server Error")
)

var _ account.Repository = (*Database)(nil)

type Database struct {
	data map[uuid.UUID]*accounts.Account
}

// NewRepository, cria um novo repositório do banco
func NewRepository() *Database {
	return &Database{
		data: make(map[uuid.UUID]*accounts.Account),
	}
}

// SaveAccount, insere a conta no banco
func (db *Database) SaveAccount(_ context.Context, account accounts.Account) (uuid.UUID, error) {
	var uuID = uuid.New()

	if _, ok := db.data[uuID]; ok {
		return uuid.Nil, ErrIDExists
	}

	account.ID = uuID.String()
	account.CreatedAt = time.Now().UTC()
	db.data[uuID] = &account

	return uuID, nil
}

// ListBalanceByID, exibe o saldo da conta, pelo id.
func (db *Database) ListBalanceByID(_ context.Context, accID uuid.UUID) (int, error) {

	if acc, ok := db.data[accID]; ok {
		return acc.Balance, nil
	}

	return 0, ErrNotFound
}

// ListAllAccounts, lista todas as contas
func (db *Database) ListAllAccounts(_ context.Context) ([]accounts.Account, error) {

	var accountsList []accounts.Account

	for _, account := range db.data {
		accountsList = append(accountsList, *account)
	}

	return accountsList, nil
}

// ListAccountsByCPF, lista conta pelo CPF
func (db *Database) ListAccountsByCPF(ctx context.Context, accCPF string) (accounts.Account, error) {

	listAcc, _ := db.ListAllAccounts(ctx)

	for _, account := range listAcc {
		if account.CPF == accCPF {
			return account, nil
		}
	}

	return accounts.Account{}, ErrNotFound
}

// ListAccountsByID, verifica a existência de uma conta pelo ID
func (db *Database) ListAccountByID(_ context.Context, acc string) (accounts.Account, error) {

	accID, err := uuid.Parse(acc)
	if err != nil {
		return accounts.Account{}, err
	}

	if account, ok := db.data[accID]; ok {
		acc := accounts.Account{
			ID:      account.ID,
			Name:    account.Name,
			CPF:     account.CPF,
			Balance: account.Balance,
		}
		return acc, nil
	}

	return accounts.Account{}, ErrNotFound
}

// UpdatedAccount, atualiza o saldo da conta
func (db *Database) UpdatedAccount(ctx context.Context, b accounts.Balance) error {
	accID, err := uuid.Parse(b.ID)
	if err != nil {
		return err
	}

	if _, ok := db.data[accID]; ok {
		db.data[accID].Balance = b.Balance
		return nil
	}

	return ErrInternal
}
