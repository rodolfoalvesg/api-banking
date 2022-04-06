package db

import (
	"context"
	"time"
)

// Database, metódos do banco de dados
type DatabaseMethods interface {
	AddedAccount(ctx context.Context) (Database, error)
	ShowBalanceId(ctx context.Context) (Database, error)
	ShowAccounts(ctx context.Context) ([]Database, error)
	FindDocument(ctx context.Context) (Database, error)
}

type Database struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

// NewRepositoryDB, cria um novo repositório do banco
func NewRepositoryDB(db DatabaseMethods) *Database {
	return &Database{}
}

// addedAccount, insere a conta no banco
func (db *Database) AddedAccount(ctx context.Context) ([]Database, error) {

	data, err := db.ShowAccounts(ctx)
	if err != nil {
		return []Database{}, err
	}
	data = append(data, *db)

	return data, nil
}

// showBalanceId, exibe o saldo da conta, pelo id.
/*func (db *Database) ShowBalanceID() (accounts.Account, error) {

	for _, account := range baseAccounts {
		if f.ID == account.ID {
			return account, nil
		}
	}
	return models.Account{}, errors.New("Conta não localizada")
}*/

// showAccounts, lista todas as contas
func (db *Database) ShowAccounts(ctx context.Context) ([]Database, error) {
	dataAccounts := make([]Database, 0)

	return dataAccounts, nil
}

// findDocument Procurar se existe o cpf passado
/*func (db *Database) FindDocument() (models.Account, error) {
	for _, document := range baseAccounts {
		if f.CPF == document.CPF {
			return document, nil
		}
	}

	return models.Account{}, nil
}*/
