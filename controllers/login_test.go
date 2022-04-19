package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
)

func TestLoginHandler(t *testing.T) {
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
				ListAccountsByCPF: func(accCPF string) (string, error) {
					return myLoginFake.Secret, nil
				},
			},
			loginBody:      myLoginFake,
			wantStatusCode: 200,
		},
		{
			Name: "Body Invalid",
			loginMock: account.UseCaseMock{
				ListAccountsByCPF: func(accCPF string) (string, error) {
					return "", nil
				},
			},
			loginBody:      "invalid",
			wantStatusCode: 422,
		},
		{
			Name: "Login Fail",
			loginMock: account.UseCaseMock{
				ListAccountsByCPF: func(accCPF string) (string, error) {
					return "", errors.New("Error")
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
			})

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
