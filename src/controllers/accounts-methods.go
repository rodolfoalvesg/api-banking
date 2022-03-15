package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/src/db"
	"github.com/rodolfoalvesg/api-banking/api/src/models"
	"github.com/rodolfoalvesg/api-banking/api/src/responses"
	"github.com/rodolfoalvesg/api-banking/api/src/security"
)

// CreateAccount cria uma conta
func (c *Controller) HandlerCreateAccount(w http.ResponseWriter, r *http.Request) {
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

	defer r.Body.Close()

	passwdHash, err := security.SecurityHash(account.Secret) //Cria um hash da senha passada
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	account.Secret = string(passwdHash)  //account.Secret, atribui ao campo de senha do modelo o HASH
	account.Id = uuid.New().String()     //account.Id, cria um id único e atribui ao campo Id
	account.CreatedAt = time.Now().UTC() //account.CreatedAt, data e hora
	modelAccount := &db.FieldsToMethodsDB{
		Accounts: account,
	}

	modelAccount.AddedAccount()

	modelList := &db.FieldsToMethodsDB{}
	data, err := modelList.ShowAccounts()
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	responses.RespondJSON(w, http.StatusOK, data)
}

// ShowBalance, exibe o saldo
func (c *Controller) HandlerShowBalance(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	accountId := params["account_id"]

	modelListId := &db.FieldsToMethodsDB{
		Id: accountId,
	}

	accountPerson, err := modelListId.ShowBalanceId()
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	responseAccount := &db.FieldsToMethodsDB{
		Balance: accountPerson.Balance,
	}

	responses.RespondJSON(w, http.StatusOK, responseAccount.Balance)
}

// ShowAccounts, lista as contas
func (c *Controller) HandlerShowAccounts(w http.ResponseWriter, r *http.Request) {
	modelShowAccounts := &db.FieldsToMethodsDB{}
	accountLits, err := modelShowAccounts.ShowAccounts()
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	responses.RespondJSON(w, http.StatusOK, accountLits)
}
