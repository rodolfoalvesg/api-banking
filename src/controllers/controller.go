package controllers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/src/db"
)

type Controls interface {
	HandlerCreateAccount(http.ResponseWriter, *http.Request)
	HandlerShowBalance(http.ResponseWriter, *http.Request)
	HandlerShowAccounts(http.ResponseWriter, *http.Request)
}
type Controller struct{}

func NewController(db db.Database) *Controller {
	return &Controller{}
}
