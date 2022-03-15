package db

import (
	"fmt"

	"github.com/rodolfoalvesg/api-banking/api/src/models"
)

var baseAccounts []models.Account

type Database interface {
	AddedAccount()
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

// addedAccount, insere a conta no banco
func (f *FieldsToMethodsDB) AddedAccount() {
	baseAccounts = append(baseAccounts, f.Accounts)
}

// showBalanceId, exibe o saldo da conta, pelo id.
func (f *FieldsToMethodsDB) ShowBalanceId() (models.Account, error) {

	for _, account := range baseAccounts {
		if f.Id == account.Id {
			return account, nil
		}
	}
	return models.Account{}, fmt.Errorf("Não encontrado")
}

// showAccounts, lista todas as contas
func (f *FieldsToMethodsDB) ShowAccounts() ([]models.Account, error) {
	return baseAccounts, nil
}

// findDocument Procurar se existe o cpf passado
func (f *FieldsToMethodsDB) FindDocument() (models.Account, error) {
	for _, document := range baseAccounts {
		if f.Cpf == document.Cpf {
			return document, nil
		}
	}

	return models.Account{}, nil
}
