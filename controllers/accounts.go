package controllers

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
)

// CreateAccount cria uma conta
func (c *Controller) CreateAccount(w http.ResponseWriter, r *http.Request) {
	var acc accounts.Account

	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	// Analisa se a senha atende os critérios
	if len(acc.Secret) < 8 {
		responses.RespondError(w, http.StatusFailedDependency, errors.New("A senha nao atende aos requisitos"))
		return
	}

	accCreated, err := c.account.CreateAccount(r.Context(), acc)
	if err != nil {
		responses.RespondJSON(w, http.StatusBadRequest, err)
		return
	}

	responses.RespondJSON(w, http.StatusCreated, accCreated)
}

// ShowBalance, exibe o saldo
func (c *Controller) ShowBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID := params["account_id"]

	if accountID == "" {
		responses.RespondJSON(w, http.StatusBadRequest, errors.New("ID vazio, favor informar um id válido"))
	}

	accBalance, err := c.account.ShowBalance(r.Context(), accountID)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accBalance)
}

// ShowAccounts, lista as contas
func (c *Controller) ShowAccounts(w http.ResponseWriter, r *http.Request) {
	accList, err := c.account.ShowAccounts(r.Context())
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accList)
}
