package controllers

import (
	"encoding/json"
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

	// ExtractUserID, extrai o id do usuário do token
	userIDToken, err := security.ExtractUserID(r)
	if err != nil {
		responses.RespondError(w, http.StatusUnauthorized, err)
		return
	}

	transfer.AccountOriginID = userIDToken

	// ValidateTransferData, valida os dados de entrada
	err = transfers.ValidateTransferData(transfer)
	if err != nil {
		responses.RespondError(w, http.StatusPreconditionFailed, err)
		return
	}

	// GetAccount, verifica se a conta de destino existe
	_, err = c.account.GetAccount(r.Context(), transfer.AccountDestinationID)
	if err != nil {
		responses.RespondError(w, http.StatusNotFound, err)
		return
	}

	// CreateTransfer, salva a transferencia
	transferID, err := c.transfer.CreateTransfer(r.Context(), transfer)
	if err != nil {
		responses.RespondError(w, http.StatusConflict, err)
		return
	}

	// UpdateAccount, atualiza o saldo das contas
	err = c.account.UpdateAccount(r.Context(), transfer)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusCreated, transferID)

}

// ListTransferHandler, lista transferências de um usuário autenticado
func (c *Controller) ListTransferHandler(w http.ResponseWriter, r *http.Request) {

	userIDToken, err := security.ExtractUserID(r)
	if err != nil {
		responses.RespondError(w, http.StatusUnauthorized, err)
		return
	}

	TransferList, err := c.transfer.ShowTransfers(r.Context(), userIDToken)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
		return
	}

	responses.RespondJSON(w, http.StatusOK, TransferList)
}
