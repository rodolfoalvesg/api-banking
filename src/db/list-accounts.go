package db

import (
	"fmt"

	"github.com/rodolfoalvesg/api-banking/api/src/models"
)

var BaseAccounts []models.Account

func CreatedAccount(account models.Account) {
	BaseAccounts = append(BaseAccounts, account)
}

func ListBalance(id string) (models.Account, error) {
	for _, account := range BaseAccounts {
		if id == account.Id {
			return account, nil
		}
	}
	return models.Account{}, fmt.Errorf("Não encontrado")
}

func ListAccounts() ([]models.Account, error) {
	return BaseAccounts, nil
}

func SearchDocument(cpf string) (models.Account, error) {
	for _, document := range BaseAccounts {
		if cpf == document.Cpf {
			return document, nil
		}
	}

	return models.Account{}, nil
}
