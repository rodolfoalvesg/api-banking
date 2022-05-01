package controllers

import (
	"encoding/json"
	"net/http"

	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
)

// LoginHandler, cria o login para a api
func (c *Controller) LoginHandler(w http.ResponseWriter, r *http.Request) {

	var acc account.Login
	if err := json.NewDecoder(r.Body).Decode(&acc); err != nil {
		responses.RespondError(w, http.StatusUnprocessableEntity, err)
		return
	}

	defer r.Body.Close()

	tokenID, err := c.account.NewLogin(r.Context(), acc)
	if err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, tokenID)
}
