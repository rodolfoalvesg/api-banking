package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/domain/models"
	"github.com/rodolfoalvesg/api-banking/api/gateways/db"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/security"
)

// Login, cria o logon para a api
func (c *Controller) HandlerLogin(w http.ResponseWriter, r *http.Request) {
	bodyRequest, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.RespondError(w, http.StatusUnprocessableEntity, err)
		return
	}

	var account models.Account
	if err := json.Unmarshal(bodyRequest, &account); err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	defer r.Body.Close()

	modelFindDocument := db.FieldsToMethodsDB{}
	verifyDocument, err := modelFindDocument.FindDocument()
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	err = security.VerifyPasswd(verifyDocument.Secret, account.Secret)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	token, err := security.CreateToken(verifyDocument.Id)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
	}

	responses.RespondJSON(w, http.StatusOK, models.Authentication{ID: verifyDocument.Id, Token: token})
}
