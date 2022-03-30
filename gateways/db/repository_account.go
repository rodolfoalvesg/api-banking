package db

import "github.com/rodolfoalvesg/api-banking/api/domain/models"

type Database interface {
	AddedAccount() (models.Account, error)
	ShowBalanceId() (models.Account, error)
	ShowAccounts() ([]models.Account, error)
	FindDocument() (models.Account, error)
}

type FieldsToMethodsDB struct {
	Accounts models.Account
	Id       string
	Cpf      string
	Balance  int `json:"balance,omitempty"`
}

// NewRepositoryDB, cria um novo repositório do banco
func NewRepositoryDB(db Database) *FieldsToMethodsDB {
	return &FieldsToMethodsDB{}
}
