package controllers

import (
	"net/http"
)

type Controls interface {
	HandlerCreateAccount(http.ResponseWriter, *http.Request)
	HandlerShowBalance(http.ResponseWriter, *http.Request)
	HandlerShowAccounts(http.ResponseWriter, *http.Request)
}

type Controller struct{}

func NewController(c Controls) *Controller {
	return &Controller{}
}
