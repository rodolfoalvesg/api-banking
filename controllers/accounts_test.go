package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/domain/models"
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

	accFake := models.Account{
		ID:     "dfsh15hjfg4hgfsdhgdsf",
		Secret: "123456789",
	}

	accListA, _ := accounts.NewCreateAccount(accFake)
	accListB := models.Account{}

	controller := NewController(nil)

	testShowBalance := map[string]struct {
		accBalanceID models.Account
		want         int
	}{
		"Status OK":  {accListA, http.StatusOK},
		"Status Bad": {accListB, http.StatusBadRequest},
	}

	for name, tt := range testShowBalance {

		path := fmt.Sprintf("/accounts/%s/balance", tt.accBalanceID.ID)

		request := httptest.NewRequest(http.MethodGet, path, nil)
		response := httptest.NewRecorder()

		vars := map[string]string{
			"account_id": tt.accBalanceID.ID,
		}

		request = mux.SetURLVars(request, vars)

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

	request := httptest.NewRequest(http.MethodGet, "/accounts", nil)
	response := httptest.NewRecorder()

	controller.HandlerShowAccounts(response, request)

	fmt.Println(response)

	if response.Result().StatusCode != http.StatusOK {
		t.Fail()
	}
}
