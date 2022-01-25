package controllers

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/src/models"
	"github.com/rodolfoalvesg/api-banking/api/src/responses"
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

	fmt.Println(transfers)

}

func ShowTransfers(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Listando transferências"))
}
