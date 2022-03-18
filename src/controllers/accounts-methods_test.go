package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rodolfoalvesg/api-banking/api/src/models"
)

// TestHandlerCreateAccount, teste do handler de criação de conta
func TestHandlerCreateAccount(t *testing.T) {
	t.Parallel()

	personAccountA := models.Account{
		Secret: "12345",
	}
	personAccountB := models.Account{
		Secret: "123456794",
	}

	controller := NewController(nil)

	testCreateAccount := map[string]struct {
		personACC models.Account
		want      int
	}{

		"Status Bad": {personAccountA, http.StatusBadRequest},
		"Status OK":  {personAccountB, http.StatusOK},
	}

	for name, tt := range testCreateAccount {
		response := httptest.NewRecorder()

		bodyBytes, err := json.Marshal(tt.personACC)
		if err != nil {
			t.Fatal(err)
		}

		body := bytes.NewReader(bodyBytes)

		request := httptest.NewRequest(http.MethodPost, "/accounts", body)

		controller.HandlerCreateAccount(response, request)

		responseCode := response.Result().StatusCode

		if responseCode != tt.want {
			t.Errorf("%s: got %v, want %v", name, responseCode, tt.want)
		}
	}

}

//TestHandlerShowBalance, teste do handler para exibição de saldo
func TestHandlerShowBalance(t *testing.T) {
	t.Parallel()

	accBalanceA := models.Account{
		Id: "kgdf4gf4gfdgf554gsfag4g",
	}

	controller := NewController(nil)

	testShowBalance := map[string]struct {
		accBalanceId models.Account
		want         int
	}{
		"Status OK": {accBalanceA, http.StatusOK},
		//"Status Bad": {personAccountB.Id, http.StatusBadRequest},
	}

	for name, tt := range testShowBalance {

		path := fmt.Sprintf("/accounts/%s/balance", tt.accBalanceId.Id)

		request := httptest.NewRequest(http.MethodGet, path, nil)
		response := httptest.NewRecorder()

		controller.HandlerShowBalance(response, request)

		respondeCode := response.Result().StatusCode

		if respondeCode != tt.want {
			t.Errorf("%s: got %v, want %v", name, respondeCode, tt.want)
		}

	}
}

// TestHandlerShowAccounts, teste do handler para listagem de conta
func TestHandlerShowAccounts(t *testing.T) {
	t.Parallel()

	controller := NewController(nil)

	response := httptest.NewRecorder()
	request := httptest.NewRequest(http.MethodGet, "/accounts", nil)

	controller.HandlerShowAccounts(response, request)

	if response.Result().StatusCode != http.StatusOK {
		t.Fail()
	}

}
