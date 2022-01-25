package db

import (
	"fmt"

	"github.com/rodolfoalvesg/api-banking/api/src/models"
)

var baseAccounts []models.Account

type Database interface {
	AddedAccount(models.Account)
	ShowBalanceId(id string) (models.Account, error)
	ShowAccounts() ([]models.Account, error)
	FindDocument(cpf string) (models.Account, error)
}

type CreateAccount struct {
	account models.Account
}

type ListBalance struct {
	id string
}

type ListAccounts []models.Account

type SearchDocument struct {
	cpf string
}

func (c *CreateAccount) addedAccount() {
	baseAccounts = append(baseAccounts, c.account)
}

func (l ListBalance) showBalanceId() (models.Account, error) {
	for _, account := range baseAccounts {
		if l.id == account.Id {
			return account, nil
		}
	}
	return models.Account{}, fmt.Errorf("Não encontrado")
}

func (l ListAccounts) showAccounts() ([]models.Account, error) {
	return baseAccounts, nil
}

func (s SearchDocument) findDocument() (models.Account, error) {
	for _, document := range baseAccounts {
		if s.cpf == document.Cpf {
			return document, nil
		}
	}

	return models.Account{}, nil
}
