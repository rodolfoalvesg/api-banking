package controllers

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/rodolfoalvesg/api-banking/api/src/models"
)

func TestHandlers(t *testing.T) {
	t.Parallel()

	personAccount := models.Account{
		Name:    "Rodolfo Alves",
		Cpf:     "01225465245",
		Secret:  "12345678",
		Balance: 5000,
	}

	controller := NewController(nil)

	response := httptest.NewRecorder()

	bodyBytes, err := json.Marshal(personAccount)
	if err != nil {
		t.Fatal(err)
	}

	body := bytes.NewReader(bodyBytes)

	t.Run("Test HandlerCreateAccount", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodPost, "/accounts", body)

		controller.HandlerCreateAccount(response, request)
		if response.Result().StatusCode != http.StatusOK {
			t.Fail()
		}

	})

	t.Run("Test HandlerShowBalance", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/accounts/{account_id}/balance", body)

		controller.HandlerShowBalance(response, request)
		if response.Result().StatusCode != http.StatusOK {
			t.Fail()
		}
	})

	t.Run("Testt HandlerShowAccounts", func(t *testing.T) {
		request := httptest.NewRequest(http.MethodGet, "/accounts", body)

		controller.HandlerShowAccounts(response, request)
		if response.Result().StatusCode != http.StatusOK {
			t.Fail()
		}
	})

}
