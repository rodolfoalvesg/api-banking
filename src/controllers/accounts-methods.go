package controllers

import (
	"encoding/json"
	"errors"
	"io/ioutil"
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/src/db"
	"github.com/rodolfoalvesg/api-banking/api/src/models"
	"github.com/rodolfoalvesg/api-banking/api/src/responses"
	"github.com/rodolfoalvesg/api-banking/api/src/security"

	"github.com/google/uuid"
)

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

	account.FormatDocumentNumber() // FormatDocumentNumber formata os valores do cpf tirando caractees
	if account.Cpf == "" || len(account.Cpf) > 11 {
		responses.RespondError(w, http.StatusBadRequest, errors.New("O número do documento deve possuir 11 caracteres"))
		return
	}

	defer r.Body.Close()

	passwdHash, err := security.SecurityHash(account.Secret) // Cria um hash da senha passada
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
	}

	account.Secret = string(passwdHash) //account.Secret, atribui ao campo de senha do modelo o HASH
	account.Id = uuid.New().String()    //account.Id, cria um id único e atribui ao campo Id

	db.CreatedAccount(account) // db.CreatedAccount, insere os dados na Base de dados

	responses.RespondJSON(w, http.StatusOK, db.BaseAccounts)
}

// ShowBallance, exibe o saldo
func ShowBallance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mostrando o saldo"))
}

// ShowAccounts, lista as contas
func ShowAccounts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando contas"))
}
