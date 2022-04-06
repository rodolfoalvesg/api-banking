package controllers

import (
	"net/http"

	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
)

type Controls interface {
	CreateAccount(http.ResponseWriter, *http.Request)
	ShowBalance(http.ResponseWriter, *http.Request)
	ShowAccounts(http.ResponseWriter, *http.Request)
}

type Controller struct {
	account account.UsecaseAccount
}

func NewController(c account.UsecaseAccount) *Controller {
	return &Controller{}
}
