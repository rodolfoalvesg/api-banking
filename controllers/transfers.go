package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/domain/models"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
)

//HandleTransfers, realiza transferencia entre contas
func HandleTransfers(w http.ResponseWriter, r *http.Request) {
	requestBody, err := ioutil.ReadAll(r.Body)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
	}

	var transfers models.Transfers
	err = json.Unmarshal(requestBody, &transfers)
	if err != nil {
		responses.RespondError(w, http.StatusInternalServerError, err)
	}

	defer r.Body.Close()
}

func ShowTransfers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando transferências"))
}
