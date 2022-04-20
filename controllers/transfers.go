package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/common/security"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
)

// CreateTransferHandler, cria uma transfência entre usuários cadastrados
func (c *Controller) CreateTransferHandler(w http.ResponseWriter, r *http.Request) {
	var transfer transfers.Transfer

	if err := json.NewDecoder(r.Body).Decode(&transfer); err != nil {
		responses.RespondError(w, http.StatusBadRequest, err)
		return
	}

	defer r.Body.Close()

	err := transfers.ValidateTransferData(transfer.Account_destination_ID, transfer.Amount)
	if err != nil {
		responses.RespondError(w, http.StatusPreconditionFailed, err)
		return
	}

	listAccounts, err := c.account.ShowAccounts(r.Context())
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	fmt.Println(listAccounts)

	userIDToken, err := security.ExtractUserID(r)
	if err != nil {
		responses.RespondError(w, http.StatusUnauthorized, err)
		return
	}

	fmt.Println(userIDToken)

}

// ListTransferHandler, lista transferências de um usuário autenticado
func (c *Controller) ListTransferHandler(w http.ResponseWriter, r *http.Request) {

}
