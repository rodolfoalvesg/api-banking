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
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
)

//TestLoginHandler, teste de requisição para login
func TestLoginHandler(t *testing.T) {
	myAccountFake := accounts.Account{
		ID:        uuid.New().String(),
		Name:      "Fulano de Tal",
		CPF:       "12345678900",
		Secret:    "12345678",
		Balance:   525000,
		CreatedAt: time.Now().UTC(),
	}

	hash, _ := accounts.GeneratePasswdHash(context.Background(), myAccountFake)
	myAccountFake.Secret = string(hash)

	myLoginFake := account.Login{
		CPF:    "12345678900",
		Secret: "12345678",
	}

	type TestCase struct {
		Name           string
		loginMock      account.UseCaseMock
		loginBody      interface{}
		wantStatusCode int
	}

	testCase := []TestCase{
		{
			Name: "Login successfully",
			loginMock: account.UseCaseMock{
				ListAccountsByCPF: func(accCPF string) (accounts.Account, error) {
					return myAccountFake, nil
				},
			},
			loginBody:      myLoginFake,
			wantStatusCode: 200,
		},
		{
			Name: "Body Invalid",
			loginMock: account.UseCaseMock{
				ListAccountsByCPF: func(accCPF string) (accounts.Account, error) {
					return accounts.Account{}, errors.New("Error")
				},
			},
			loginBody:      "invalid",
			wantStatusCode: 422,
		},
		{
			Name: "Login Fail",
			loginMock: account.UseCaseMock{
				ListAccountsByCPF: func(accCPF string) (accounts.Account, error) {
					return accounts.Account{}, errors.New("Error")
				},
			},
			loginBody:      myLoginFake,
			wantStatusCode: 400,
		},
	}

	for _, tc := range testCase {
		tt := tc
		t.Run(tc.Name, func(t *testing.T) {
			t.Parallel()

			handler := NewController(&account.UseCaseMock{
				ListAccountsByCPF: tt.loginMock.ListAccountsByCPF,
			}, nil)

			path := fmt.Sprintf("/login")
			jsonLoginBody, _ := json.Marshal(tt.loginBody)
			req := bytes.NewReader(jsonLoginBody)

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, path, req)

			http.HandlerFunc(handler.LoginHandler).ServeHTTP(w, r)

			if w.Code != tt.wantStatusCode {
				t.Errorf("%s: got %v, want %v", tt.Name, w.Code, tt.wantStatusCode)
			}

		})
	}
}
