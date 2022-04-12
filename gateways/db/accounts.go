package db

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	usecase "github.com/rodolfoalvesg/api-banking/api/domain/usecases"
)

var _ usecase.Repository = (*Database)(nil)

type Database struct {
	data map[uuid.UUID]accounts.Account
}

// NewRepositoryDB, cria um novo repositório do banco
func NewRepository() *Database {
	return &Database{
		data: make(map[uuid.UUID]accounts.Account),
	}
}

// addedAccount, insere a conta no banco
func (db *Database) SaveAccount(_ context.Context, account accounts.Account) (uuid.UUID, error) {
	var uuID = uuid.New()

	if _, ok := db.data[uuID]; ok {
		return uuID, fmt.Errorf("ID already exists!")
	}

	account.ID = uuID.String()
	db.data[uuID] = account

	return uuID, nil
}

// showBalanceId, exibe o saldo da conta, pelo id.
func (db *Database) ListBalanceByID(_ context.Context, accID uuid.UUID) (int, error) {

	if balance, ok := db.data[accID]; ok {
		return balance.Balance, nil
	}

	return 0, errors.New("Account not found")
}

// showAccounts, lista todas as contas
func (db *Database) ListAllAccounts(_ context.Context) ([]accounts.Account, error) {

	var accountsList []accounts.Account

	for _, account := range db.data {
		accountsList = append(accountsList, account)
	}

	return accountsList, nil
}
