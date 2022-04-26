package controllers

import (
	"context"
	"encoding/json"
	"errors"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
)

var (
	errEmptyID = errors.New("ID cannot be empty")
)

// CreateAccountHandler, cria uma requisição para criação de conta
func (c *Controller) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var acc accounts.Account

	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	err := accounts.ValidateCreateAccountData(acc)
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	err = c.account.VerifyAccount(context.Background(), acc.CPF)
	if err != nil {
		responses.RespondError(w, http.StatusConflict, err)
		return
	}

	accCreated, err := c.account.CreateAccount(r.Context(), acc)
	if err != nil {
		responses.RespondError(w, http.StatusConflict, err)
		return
	}

	responses.RespondJSON(w, http.StatusCreated, accCreated)

}

// ShowBalanceHandler, cria uma requisição para listar um saldo
func (c *Controller) ShowBalanceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID := params["account_id"]

	emptyID := uuid.UUID{}
	accID := uuid.MustParse(accountID)

	if accID == emptyID {
		responses.RespondError(w, http.StatusBadRequest, errEmptyID)
		return
	}

	accBalance, err := c.account.ShowBalance(r.Context(), accID)
	if err != nil {
		responses.RespondError(w, http.StatusNotFound, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accBalance)
}

// ShowAccountsHandler, cria uma requisição para listagem de todas as contas
func (c *Controller) ShowAccountsHandler(w http.ResponseWriter, r *http.Request) {
	accList, err := c.account.ShowAccounts(r.Context())
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accList)
}
