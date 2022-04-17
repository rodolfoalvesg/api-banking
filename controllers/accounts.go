package controllers

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
)

// CreateAccount cria uma conta
func (c *Controller) CreateAccountHandler(w http.ResponseWriter, r *http.Request) {
	var acc accounts.Account

	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	// Analisa se a senha atende os critérios
	if len(acc.Secret) < 8 {
		responses.RespondError(w, http.StatusFailedDependency, fmt.Errorf("password must be at least 8 characters long"))
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
func (c *Controller) ShowBalanceHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountID := params["account_id"]

	fmt.Println(accountID)

	if accountID == "00000000-0000-0000-0000-000000000000" {
		responses.RespondJSON(w, http.StatusBadRequest, errors.New("ID cannot be empty."))
		return
	}

	accID := uuid.MustParse(accountID)

	accBalance, err := c.account.ShowBalance(r.Context(), accID)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accBalance)
}

// ShowAccounts, lista as contas
func (c *Controller) ShowAccountsHandler(w http.ResponseWriter, r *http.Request) {
	accList, err := c.account.ShowAccounts(r.Context())
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accList)
}
