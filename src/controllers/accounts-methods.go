package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/src/db"
	"github.com/rodolfoalvesg/api-banking/api/src/models"
	"github.com/rodolfoalvesg/api-banking/api/src/responses"
	"github.com/rodolfoalvesg/api-banking/api/src/security"
)

type RouteMethods interface {
	HandleCreateAccount(w http.ResponseWriter, r *http.Request)
}

type Controller struct {
	db db.Database
}

type BalanceAccount struct {
	Balance int `json:"balance,omitempty"`
}

// CreateAccount cria uma conta
func (c *Controller) HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.RespondError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var account models.Account
	if err := json.Unmarshal(bodyRequest, &account); err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	account.FormatDocumentNumber() //FormatDocumentNumber formata os valores do cpf tirando caractees
	if account.Cpf == "" || len(account.Cpf) > 11 {
		responses.RespondError(w, http.StatusBadRequest, errors.New("O número do documento deve possuir 11 caracteres"))
		return
	}

	defer r.Body.Close()

	passwdHash, err := security.SecurityHash(account.Secret) //Cria um hash da senha passada
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	account.Secret = string(passwdHash) //account.Secret, atribui ao campo de senha do modelo o HASH
	account.Id = uuid.New().String()    //account.Id, cria um id único e atribui ao campo Id

	c.db.AddedAccount(account) //db.CreatedAccount, insere os dados na Base de dados

	responses.RespondJSON(w, http.StatusOK, c.db.ShowAccounts)

}

// func HandleCreateAccount(w http.ResponseWriter, r *http.Request) {
// }

// ShowBalance, exibe o saldo
func (c *Controller) ShowBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["account_id"]

	accountPerson, err := c.db.ShowBalanceId(accountId)
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	responseAccount := BalanceAccount{
		Balance: accountPerson.Balance,
	}

	responses.RespondJSON(w, http.StatusOK, responseAccount)
}

// func ShowBalance(w http.ResponseWriter, r *http.Request) {
// }

// ShowAccounts, lista as contas
func (c *Controller) ShowAccounts(w http.ResponseWriter, r *http.Request) {

	accountLits, err := c.db.ShowAccounts()
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	responses.RespondJSON(w, http.StatusOK, accountLits)
}
