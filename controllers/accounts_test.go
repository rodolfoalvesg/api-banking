package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
	"github.com/rodolfoalvesg/api-banking/api/gateways/db"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/responses"
)

// TestCreateAccount, teste do handler de criação de conta
type AccountRequest struct {
	Name    string `json:"name"`
	CPF     string `json:"cpf"`
	Secret  string `json:"secret"`
	Balance int    `json:"balance"`
}

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		Name           string
		accountMock    account.AccountMock
		bodyAcc        interface{}
		want           uuid.UUID
		wantStatusCode int
	}

	myAccount := accounts.Account{
		Name:    "Teste",
		CPF:     "12345678900",
		Secret:  "12345678",
		Balance: 25000,
	}

	accID, _ := db.NewRepository().SaveAccount(context.TODO(), myAccount)

	tests := []TestCase{
		{
			Name: "Account created successfully",
			accountMock: account.AccountMock{
				CreateM: func(acc accounts.Account) (uuid.UUID, error) {
					return accID, nil
				},
			},
			bodyAcc: AccountRequest{
				Name:    myAccount.Name,
				CPF:     myAccount.CPF,
				Secret:  myAccount.Secret,
				Balance: myAccount.Balance,
			},
			wantStatusCode: http.StatusOK,
			want:           accID,
		}, {
			Name: "Account not created",
			accountMock: account.AccountMock{
				CreateM: func(acc accounts.Account) (uuid.UUID, error) {
					return accID, nil
				},
			},
			bodyAcc:        "invalid",
			wantStatusCode: http.StatusBadRequest,
			want:           uuid.UUID{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := NewController(account.Usecase{})

			path := fmt.Sprintf("/accounts")
			jsonBodyAcc, _ := json.Marshal(tt.bodyAcc)
			req := bytes.NewReader(jsonBodyAcc)
			request := httptest.NewRequest(http.MethodPost, path, req)
			response := httptest.NewRecorder()

			http.HandlerFunc(handler.CreateAccount).ServeHTTP(response, request)

			responses.RespondJSON(response, tt.wantStatusCode, response.Code)

		})
	}

}

/*

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
*/
