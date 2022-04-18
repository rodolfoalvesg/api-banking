package controllers

import (
	"encoding/json"
	"net/http"

	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
)

// Login, cria o logon para a api
func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var acc account.Login
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	accountID, err := c.account.NewLogin(r.Context(), acc)
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	// modelFindDocument := db.Database{}
	// verifyDocument, err := modelFindDocument.FindDocument()
	// if err != nil {
	// 	responses.RespondError(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// err = security.VerifyPasswd(verifyDocument.Secret, account.Secret)
	// if err != nil {
	// 	responses.RespondError(w, http.StatusInternalServerError, err)
	// 	return
	// }

	// token, err := security.CreateToken(verifyDocument.ID)
	// if err != nil {
	// 	responses.RespondError(w, http.StatusInternalServerError, err)
	// }

	responses.RespondJSON(w, http.StatusOK, accountID)
}
