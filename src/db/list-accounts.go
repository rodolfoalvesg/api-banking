package db

import (
	"fmt"

	"github.com/rodolfoalvesg/api-banking/api/src/models"
)

var BaseAccounts []models.Account

func CreatedAccount(account models.Account) {
	BaseAccounts = append(BaseAccounts, account)
	fmt.Println(BaseAccounts)
}

func ListAccount(id string) (models.Account, error) {
	for _, account := range BaseAccounts {
		if id == account.Id {
			return account, nil
		}
	}
	return models.Account{}, fmt.Errorf("Não encontrado")
}
