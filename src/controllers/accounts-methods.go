package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/src/db"
	"github.com/rodolfoalvesg/api-banking/api/src/models"
	"github.com/rodolfoalvesg/api-banking/api/src/responses"
	"github.com/rodolfoalvesg/api-banking/api/src/security"

	"github.com/google/uuid"
)

type BalanceAccount struct {
	Balance int `json:"balance,omitempty"`
}

// CreateAccount cria uma conta
func CreateAccount(w http.ResponseWriter, r *http.Request) {
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

	db.CreatedAccount(account) //db.CreatedAccount, insere os dados na Base de dados

	responses.RespondJSON(w, http.StatusOK, db.BaseAccounts)
}

// ShowBalance, exibe o saldo
func ShowBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["account_id"]

	accountPerson, err := db.ListBalance(accountId)
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	responseAccount := BalanceAccount{
		Balance: accountPerson.Balance,
	}

	responses.RespondJSON(w, http.StatusOK, responseAccount)
}

// ShowAccounts, lista as contas
func ShowAccounts(w http.ResponseWriter, r *http.Request) {

	accountLits, err := db.ListAccounts()
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	responses.RespondJSON(w, http.StatusOK, accountLits)
}
