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
	FindDocument(cpf string) (models.Account, error)
}

type CreateAccount struct {
	Account models.Account
}

type ListBalance struct {
	Id string
}

type SearchDocument struct {
	cpf string
}

// addedAccount, insere a conta no banco
func (c CreateAccount) AddedAccount() {
	fmt.Println("cheguei aqui")
	baseAccounts = append(baseAccounts, c.Account)
	fmt.Println(baseAccounts)
}

// showBalanceId, exibe o saldo da conta, pelo id.
func (l ListBalance) ShowBalanceId() (models.Account, error) {
	for _, account := range baseAccounts {
		if l.Id == account.Id {
			return account, nil
		}
	}
	return models.Account{}, fmt.Errorf("Não encontrado")
}

// showAccounts, lista todas as contas
func (l ListBalance) ShowAccounts() ([]models.Account, error) {
	return baseAccounts, nil
}

// findDocument Procurar se existe o cpf passado
func (s SearchDocument) FindDocument() (models.Account, error) {
	for _, document := range baseAccounts {
		if s.cpf == document.Cpf {
			return document, nil
		}
	}

	return models.Account{}, nil
}
