package controllers

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
)

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	type TestCaseA struct {
		Name           string
		accountMock    account.UseCaseMock
		bodyAcc        interface{}
		wantStatusCode int
		want           uuid.UUID
	}

	myAccountA := accounts.Account{
		Name:    "Teste",
		CPF:     "12345678900",
		Secret:  "12345678",
		Balance: 25000,
	}

	myAccountB := accounts.Account{}

	var accID = uuid.New()
	tests := []TestCaseA{
		{
			Name: "Account created successfully",
			accountMock: account.UseCaseMock{
				SaveAccount: func(accounts.Account) (uuid.UUID, error) {
					return accID, nil
				},
			},
			bodyAcc:        myAccountA,
			wantStatusCode: http.StatusCreated,
			want:           accID,
		}, {
			Name: "Account not created",
			accountMock: account.UseCaseMock{
				SaveAccount: func(accounts.Account) (uuid.UUID, error) {
					return accID, nil
				},
			},
			bodyAcc:        "invalid",
			wantStatusCode: http.StatusBadRequest,
			want:           uuid.UUID{},
		}, {
			Name: "Invalid password length",
			accountMock: account.UseCaseMock{
				SaveAccount: func(accounts.Account) (uuid.UUID, error) {
					return accID, nil
				},
			},
			bodyAcc:        myAccountB,
			wantStatusCode: http.StatusFailedDependency,
			want:           uuid.UUID{},
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := NewController(&account.UseCaseMock{
				SaveAccount: tt.accountMock.SaveAccount,
			})

			path := fmt.Sprintf("/accounts")
			jsonBodyAcc, _ := json.Marshal(tt.bodyAcc)
			req := bytes.NewReader(jsonBodyAcc)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, path, req)

			http.HandlerFunc(handler.CreateAccountHandler).ServeHTTP(w, r)

			if w.Code != tt.wantStatusCode {
				t.Errorf("%s: got %v, want %v", tt.Name, tt.wantStatusCode, w.Code)
			}

		})
	}
}

//TestHandlerShowBalance, teste do handler para exibição de saldo
func TestHandlerShowBalance(t *testing.T) {
	t.Parallel()

	type TestCaseB struct {
		Name           string
		accountMock    account.UseCaseMock
		accountID      uuid.UUID
		wantStatusCode int
		want           int
	}

	balance := 15000
	accID := uuid.New()

	tests := []TestCaseB{
		{
			Name: "Balance successfully listed",
			accountMock: account.UseCaseMock{
				ListBalanceByID: func(uuid.UUID) (int, error) {
					return balance, nil
				},
			},
			accountID:      accID,
			wantStatusCode: http.StatusOK,
			want:           balance,
		}, {
			Name: "Error listing balance",
			accountMock: account.UseCaseMock{
				ListBalanceByID: func(uuid.UUID) (int, error) {
					return 0, nil
				},
			},
			accountID:      uuid.UUID{},
			wantStatusCode: http.StatusBadRequest,
			want:           0,
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := NewController(&account.UseCaseMock{
				ListBalanceByID: tt.accountMock.ListBalanceByID,
			})

			path := fmt.Sprintf("/accounts/{account_id}/balance")
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, path, nil)

			vars := map[string]string{
				"account_id": tt.accountID.String(),
			}

			r = mux.SetURLVars(r, vars)

			http.HandlerFunc(handler.ShowBalanceHandler).ServeHTTP(w, r)

			if w.Code != tt.wantStatusCode {
				t.Errorf("%s: got %v, want %v", tt.Name, tt.wantStatusCode, w.Code)
			}
		})
	}
}

/*
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
