package db

import (
	"github.com/rodolfoalvesg/api-banking/api/src/models"
)

var BaseAccounts []models.Account

func CreatedAccount(account models.Account) {
	BaseAccounts = append(BaseAccounts, account)
}
