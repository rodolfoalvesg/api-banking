package controllers

import (
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/gorilla/mux"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
)

//TestCreateAccount, teste do handler para criação de conta
func TestCreateAccountHandler(t *testing.T) {
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

	myAccountC := accounts.Account{
		CPF:    "1234567890",
		Secret: "12345678",
	}

	var accID = uuid.New()
	tests := []TestCaseA{
		{
			Name: "Account created successfully",
			accountMock: account.UseCaseMock{
				SaveAccount: func(accounts.Account) (uuid.UUID, error) {
					return accID, nil
				},
				VerifyAccountByCPF: func(string) error {
					return nil
				},
			},
			bodyAcc:        myAccountA,
			wantStatusCode: http.StatusCreated,
			want:           accID,
		},
		{
			Name: "Account not created",
			accountMock: account.UseCaseMock{
				SaveAccount: func(accounts.Account) (uuid.UUID, error) {
					return uuid.UUID{}, nil
				},
			},
			bodyAcc:        "invalid",
			wantStatusCode: http.StatusBadRequest,
			want:           uuid.UUID{},
		},
		{
			Name: "Invalid password length",
			accountMock: account.UseCaseMock{
				SaveAccount: func(accounts.Account) (uuid.UUID, error) {
					return uuid.UUID{}, nil
				},
			},
			bodyAcc:        myAccountB,
			wantStatusCode: http.StatusBadRequest,
			want:           uuid.UUID{},
		},
		{
			Name: "Conflict ID",
			accountMock: account.UseCaseMock{
				SaveAccount: func(accounts.Account) (uuid.UUID, error) {
					return uuid.UUID{}, errors.New("Conflict")
				},
				VerifyAccountByCPF: func(string) error {
					return nil
				},
			},
			bodyAcc:        myAccountA,
			wantStatusCode: http.StatusConflict,
			want:           uuid.UUID{},
		},
		{
			Name: "Invalid CPF",
			accountMock: account.UseCaseMock{
				SaveAccount: func(accounts.Account) (uuid.UUID, error) {
					return uuid.UUID{}, nil
				},
			},
			bodyAcc:        myAccountC,
			wantStatusCode: http.StatusBadRequest,
			want:           accID,
		},
		{
			Name: "Account already exist",
			accountMock: account.UseCaseMock{
				SaveAccount: func(accounts.Account) (uuid.UUID, error) {
					return accID, nil
				},
				VerifyAccountByCPF: func(string) error {
					return errors.New("account already exist")
				},
			},
			bodyAcc:        myAccountA,
			wantStatusCode: http.StatusConflict,
			want:           accID,
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := NewController(&account.UseCaseMock{
				SaveAccount:        tt.accountMock.SaveAccount,
				VerifyAccountByCPF: tt.accountMock.VerifyAccountByCPF,
			}, nil)

			path := fmt.Sprintf("/accounts")
			jsonBodyAcc, _ := json.Marshal(tt.bodyAcc)
			req := bytes.NewReader(jsonBodyAcc)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, path, req)

			http.HandlerFunc(handler.CreateAccountHandler).ServeHTTP(w, r)

			if w.Code != tt.wantStatusCode {
				t.Errorf("%s: got %v, want %v", tt.Name, w.Code, tt.wantStatusCode)
			}

		})
	}
}

//TestHandlerShowBalance, teste do handler para exibição de saldo
func TestShowBalanceHandler(t *testing.T) {
	t.Parallel()

	type TestCaseB struct {
		Name           string
		accountMock    account.UseCaseMock
		accountID      string
		wantStatusCode int
	}

	balance := 15000
	accID := uuid.New().String()

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
		},
		{
			Name: "Empty ID",
			accountMock: account.UseCaseMock{
				ListBalanceByID: func(uuid.UUID) (int, error) {
					return 0, nil
				},
			},
			accountID:      uuid.Nil.String(),
			wantStatusCode: http.StatusBadRequest,
		},
		{
			Name: "Invalid ID",
			accountMock: account.UseCaseMock{
				ListBalanceByID: func(uuid.UUID) (int, error) {
					return 0, errors.New("Not Found")
				},
			},
			accountID:      uuid.New().String(),
			wantStatusCode: http.StatusNotFound,
		},
		{
			Name: "Invalid accountID format",
			accountMock: account.UseCaseMock{
				ListBalanceByID: func(uuid.UUID) (int, error) {
					return 0, errors.New("invalid accountID format")
				},
			},
			accountID:      "xxxxxxx-xxxx",
			wantStatusCode: http.StatusBadRequest,
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := NewController(&account.UseCaseMock{
				ListBalanceByID: tt.accountMock.ListBalanceByID,
			}, nil)

			path := fmt.Sprintf("/accounts/{account_id}/balance")
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, path, nil)

			vars := map[string]string{
				"account_id": tt.accountID,
			}

			r = mux.SetURLVars(r, vars)

			http.HandlerFunc(handler.ShowBalanceHandler).ServeHTTP(w, r)

			if w.Code != tt.wantStatusCode {
				t.Errorf("%s: got %v, want %v", tt.Name, w.Code, tt.wantStatusCode)
			}
		})
	}
}

// TestShowAccountsHandler, teste do handler para listagem de todas as contas
func TestShowAccountsHandler(t *testing.T) {
	t.Parallel()

	type TestCaseC struct {
		Name           string
		accountMock    account.UseCaseMock
		wantStatusCode int
	}

	baseAccounts := []accounts.Account{}

	tests := []TestCaseC{
		{
			Name: "Listed Accounts",
			accountMock: account.UseCaseMock{
				ListAllAccounts: func(context.Context) ([]accounts.Account, error) {
					return baseAccounts, nil
				},
			},
			wantStatusCode: http.StatusOK,
		}, {
			Name: "Too Many Requests",
			accountMock: account.UseCaseMock{
				ListAllAccounts: func(context.Context) ([]accounts.Account, error) {
					return nil, errors.New("Internal Server Error")
				},
			},
			wantStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tc := range tests {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()
			handler := NewController(&account.UseCaseMock{
				ListAllAccounts: tt.accountMock.ListAllAccounts,
			}, nil)

			path := fmt.Sprintf("/accounts")
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, path, nil)

			http.HandlerFunc(handler.ShowAccountsHandler).ServeHTTP(w, r)

			if w.Code != tt.wantStatusCode {
				t.Errorf("%s: got %v, want %v", tt.Name, tt.wantStatusCode, w.Code)
			}
		})
	}
}
