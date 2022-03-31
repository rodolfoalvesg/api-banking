package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/domain/models"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
)

var (
	invalidID = errors.New("ID vazio, favor informar um id válido")
)

// CreateAccount cria uma conta
func (c *Controller) HandlerCreateAccount(w http.ResponseWriter, r *http.Request) {
	var account models.Account

	if err := json.NewDecoder(r.Body).Decode(&account); err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	acc, err := accounts.NewCreateAccount(account)
	if err != nil {
		responses.RespondJSON(w, http.StatusBadRequest, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, acc)
}

// ShowBalance, exibe o saldo
func (c *Controller) HandlerShowBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	accountID := params["account_id"]

	if accountID == "" {
		responses.RespondJSON(w, http.StatusBadRequest, invalidID)
	}

	accBalance, err := accounts.ShowBalance(accountID)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accBalance)
}

// ShowAccounts, lista as contas
func (c *Controller) HandlerShowAccounts(w http.ResponseWriter, r *http.Request) {
	accList, err := accounts.ShowListAccounts()
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accList)
}
