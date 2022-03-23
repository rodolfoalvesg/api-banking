package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/src/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/src/models"
	"github.com/rodolfoalvesg/api-banking/api/src/responses"
)

// CreateAccount cria uma conta
func (c *Controller) HandlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	acc, err := accounts.CreateAccount(account)
	if err != nil {
		responses.RespondJSON(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	responses.RespondJSON(w, http.StatusOK, acc)
}

// ShowBalance, exibe o saldo
func (c *Controller) HandlerShowBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	accountID := params["account_id"]

	accBalance, err := accounts.ShowBalance(accountID)
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accBalance)
}

// ShowAccounts, lista as contas
func (c *Controller) HandlerShowAccounts(w http.ResponseWriter, r *http.Request) {
	accList, err := accounts.ShowListAccounts()
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	responses.RespondJSON(w, http.StatusOK, accList)
}
