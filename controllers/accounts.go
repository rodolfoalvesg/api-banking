package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

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

	fmt.Println(acc)
	accCreated, err := c.account.CreateAccount(r.Context(), acc)
	if err != nil {
		responses.RespondJSON(w, http.StatusBadRequest, err)
		return
	}

	responses.RespondJSON(w, http.StatusCreated, accCreated)
}

// ShowBalance, exibe o saldo
/*func (c *Controller) ShowBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)

	accountID := params["account_id"]

	if accountID == "" {
		responses.RespondJSON(w, http.StatusBadRequest, errors.New("ID vazio, favor informar um id válido"))
	}

	accBalance, err := accounts.ShowBalance(accountID)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accBalance)
}

// ShowAccounts, lista as contas
func (c *Controller) ShowAccounts(w http.ResponseWriter, r *http.Request) {
	accList, err := accounts.ShowListAccounts()
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, accList)
}*/
