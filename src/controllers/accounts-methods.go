package controllers

import (
	"api/src/db"
	"api/src/models"
	"api/src/responses"
	"encoding/json"
	"io/ioutil"
	"net/http"
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

	defer r.Body.Close()

	db.BaseAccounts = append(db.BaseAccounts, account)

	responses.RespondJSON(w, http.StatusOK, db.BaseAccounts)
}

func ShowBallance(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Mostrando o saldo"))
}

func ShowAccounts(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando contas"))
}
