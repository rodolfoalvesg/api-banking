package db

import (
	"errors"

	"github.com/rodolfoalvesg/api-banking/api/domain/models"
)

var (
	ErrNotFound = errors.New("Conta não localizada")
)

var baseAccounts = []models.Account{}

// addedAccount, insere a conta no banco
func (f *FieldsToMethodsDB) AddedAccount() (models.Account, error) {
	baseAccounts = append(baseAccounts, f.Accounts)
	return baseAccounts[len(baseAccounts)-1], nil
}

// showBalanceId, exibe o saldo da conta, pelo id.
func (f *FieldsToMethodsDB) ShowBalanceID() (models.Account, error) {

	for _, account := range baseAccounts {
		if f.ID == account.ID {
			return account, nil
		}
	}
	return models.Account{}, ErrNotFound
}

// showAccounts, lista todas as contas
func (f *FieldsToMethodsDB) ShowAccounts() ([]models.Account, error) {
	return baseAccounts, nil
}

// findDocument Procurar se existe o cpf passado
func (f *FieldsToMethodsDB) FindDocument() (models.Account, error) {
	for _, document := range baseAccounts {
		if f.CPF == document.CPF {
			return document, nil
		}
	}

	return models.Account{}, nil
}
