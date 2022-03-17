package controllers

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/src/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/src/responses"
)

// CreateAccount cria uma conta
func (c *Controller) HandlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	acc, err := accounts.CreateAccount(r.Body)
	if err != nil {
		responses.RespondJSON(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	responses.RespondJSON(w, http.StatusOK, acc)
	return
}

// ShowBalance, exibe o saldo
func (c *Controller) HandlerShowBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accBalance, err := accounts.ShowBalance(params)
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accBalance)
	return
}

// ShowAccounts, lista as contas
func (c *Controller) HandlerShowAccounts(w http.ResponseWriter, r *http.Request) {
	accList, err := accounts.ShowListAccounts()
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	responses.RespondJSON(w, http.StatusOK, accList)
	return
}
