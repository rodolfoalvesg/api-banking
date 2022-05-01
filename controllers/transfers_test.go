package controllers

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/common/security"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
	account "github.com/rodolfoalvesg/api-banking/api/domain/usecases/accounts"
	transfer "github.com/rodolfoalvesg/api-banking/api/domain/usecases/transfers"
)

// TestCreateTransferHandler, teste para criação de registro de transferência
func TestCreateTransferHandler(t *testing.T) {
	listAccountsFake := []accounts.Account{
		{
			ID:      uuid.New().String(),
			CPF:     "12345678900",
			Balance: 15000,
		},
		{
			ID:      uuid.New().String(),
			CPF:     "12345678911",
			Balance: 2500,
		},
	}

	type TestCase struct {
		Name           string
		transferMock   transfer.UseCaseTransferMock
		accountMock    account.UseCaseMock
		bodyTransfer   interface{}
		tokenID        string
		formatToken    string
		wantStatusCode int
	}

	myTransferFakeA := transfers.Transfer{
		AccountDestinationID: listAccountsFake[1].ID,
		Amount:               2500,
	}

	myTransferFakeB := transfers.Transfer{}

	transferID := uuid.New()

	testCase := []TestCase{
		{
			Name: "Created successfully Transfer",
			transferMock: transfer.UseCaseTransferMock{
				SaveTransfer: func(t transfers.Transfer) (uuid.UUID, error) {
					return transferID, nil
				},
			},
			accountMock: account.UseCaseMock{
				ListAccountByID: func(string) (accounts.Account, error) {
					return listAccountsFake[0], nil
				},
				UpdatedAccount: func(accounts.Balance) error {
					return nil
				},
			},
			bodyTransfer:   myTransferFakeA,
			tokenID:        listAccountsFake[0].ID,
			formatToken:    "Bearer ",
			wantStatusCode: 201,
		},
		{
			Name: "Transfer does't created",
			transferMock: transfer.UseCaseTransferMock{
				SaveTransfer: func(t transfers.Transfer) (uuid.UUID, error) {
					return transferID, nil
				},
			},
			accountMock: account.UseCaseMock{
				ListAccountByID: func(string) (accounts.Account, error) {
					return listAccountsFake[0], nil
				},
				UpdatedAccount: func(accounts.Balance) error {
					return nil
				},
			},
			bodyTransfer:   "invalid",
			tokenID:        listAccountsFake[0].ID,
			formatToken:    "Bearer ",
			wantStatusCode: 400,
		},
		{
			Name: "A of the input fields does't meet the requirements",
			transferMock: transfer.UseCaseTransferMock{
				SaveTransfer: func(t transfers.Transfer) (uuid.UUID, error) {
					return transferID, nil
				},
			},
			accountMock: account.UseCaseMock{
				ListAccountByID: func(string) (accounts.Account, error) {
					return listAccountsFake[0], nil
				},
				UpdatedAccount: func(accounts.Balance) error {
					return nil
				},
			},
			bodyTransfer:   myTransferFakeB,
			tokenID:        listAccountsFake[0].ID,
			formatToken:    "Bearer ",
			wantStatusCode: 412,
		},
		{
			Name: "Destination account not found",
			transferMock: transfer.UseCaseTransferMock{
				SaveTransfer: func(t transfers.Transfer) (uuid.UUID, error) {
					return transferID, nil
				},
			},
			accountMock: account.UseCaseMock{
				ListAccountByID: func(string) (accounts.Account, error) {
					return accounts.Account{}, errors.New("Account not found")
				},
				UpdatedAccount: func(accounts.Balance) error {
					return nil
				},
			},
			bodyTransfer:   myTransferFakeA,
			tokenID:        listAccountsFake[0].ID,
			formatToken:    "Bearer ",
			wantStatusCode: 404,
		}, {
			Name: "UUID conflict",
			transferMock: transfer.UseCaseTransferMock{
				SaveTransfer: func(t transfers.Transfer) (uuid.UUID, error) {
					return uuid.UUID{}, errors.New("Conflict ID")
				},
			},
			accountMock: account.UseCaseMock{
				ListAccountByID: func(string) (accounts.Account, error) {
					return listAccountsFake[0], nil
				},
				UpdatedAccount: func(accounts.Balance) error {
					return nil
				},
			},
			bodyTransfer:   myTransferFakeA,
			tokenID:        listAccountsFake[0].ID,
			formatToken:    "Bearer ",
			wantStatusCode: 409,
		}, {
			Name: "Internal Server Error",
			transferMock: transfer.UseCaseTransferMock{
				SaveTransfer: func(t transfers.Transfer) (uuid.UUID, error) {
					return transferID, nil
				},
			},
			accountMock: account.UseCaseMock{
				ListAccountByID: func(string) (accounts.Account, error) {
					return listAccountsFake[0], nil
				},
				UpdatedAccount: func(accounts.Balance) error {
					return errors.New("Internal Server Error")
				},
			},
			bodyTransfer:   myTransferFakeA,
			tokenID:        listAccountsFake[0].ID,
			formatToken:    "Bearer ",
			wantStatusCode: 500,
		}, {
			Name: "Invalid Token",
			transferMock: transfer.UseCaseTransferMock{
				SaveTransfer: func(t transfers.Transfer) (uuid.UUID, error) {
					return transferID, nil
				},
			},
			accountMock: account.UseCaseMock{
				ListAccountByID: func(string) (accounts.Account, error) {
					return listAccountsFake[0], nil
				},
				UpdatedAccount: func(accounts.Balance) error {
					return nil
				},
			},
			bodyTransfer:   myTransferFakeA,
			tokenID:        listAccountsFake[0].ID,
			formatToken:    "",
			wantStatusCode: 401,
		},
	}

	for _, tc := range testCase {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := NewController(
				&account.UseCaseMock{
					ListAccountByID: tt.accountMock.ListAccountByID,
					UpdatedAccount:  tt.accountMock.UpdatedAccount,
				},
				&transfer.UseCaseTransferMock{
					SaveTransfer: tt.transferMock.SaveTransfer,
				},
			)

			token, _ := security.CreateToken(tt.tokenID)
			bearer := tt.formatToken + token

			path := fmt.Sprintf("/transfers")
			jsonBodyTransfer, _ := json.Marshal(tt.bodyTransfer)
			req := bytes.NewReader(jsonBodyTransfer)
			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodPost, path, req)

			r.Header.Add("Authorization", bearer)

			http.HandlerFunc(handler.CreateTransferHandler).ServeHTTP(w, r)

			if w.Code != tt.wantStatusCode {
				t.Errorf("%s: got %v, want %v", tt.Name, w.Code, tt.wantStatusCode)
			}

		})

	}
}

// TestListTransferHandler, teste para listagem de transferência
func TestListTransferHandler(t *testing.T) {
	listTransfersFake := []transfers.Transfer{}

	type TestCaseB struct {
		Name           string
		transferMock   transfer.UseCaseTransferMock
		tokenID        string
		formatToken    string
		wantStatusCode int
	}

	listAccountsFake := []accounts.Account{
		{
			ID:      uuid.New().String(),
			CPF:     "12345678900",
			Balance: 15000,
		},
		{
			ID:      uuid.New().String(),
			CPF:     "12345678911",
			Balance: 2500,
		},
	}

	testCase := []TestCaseB{
		{
			Name: "Transfer successfully listed",
			transferMock: transfer.UseCaseTransferMock{
				ListAllTransfers: func(string) ([]transfers.Transfer, error) {
					return listTransfersFake, nil
				},
			},
			tokenID:        listAccountsFake[0].ID,
			formatToken:    "Bearer ",
			wantStatusCode: 200,
		},
		{
			Name: "Internal Server Error",
			transferMock: transfer.UseCaseTransferMock{
				ListAllTransfers: func(string) ([]transfers.Transfer, error) {
					return nil, errors.New("Internal Server Error")
				},
			},
			tokenID:        listAccountsFake[0].ID,
			formatToken:    "Bearer ",
			wantStatusCode: 500,
		},
		{
			Name: "Invalid Token",
			transferMock: transfer.UseCaseTransferMock{
				ListAllTransfers: func(string) ([]transfers.Transfer, error) {
					return nil, errors.New("Internal Server Error")
				},
			},
			tokenID:        listAccountsFake[0].ID,
			formatToken:    "",
			wantStatusCode: 401,
		},
	}

	for _, tc := range testCase {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			handler := NewController(nil, &transfer.UseCaseTransferMock{
				ListAllTransfers: tt.transferMock.ListAllTransfers,
			})

			token, _ := security.CreateToken(tt.tokenID)
			bearer := tt.formatToken + token

			path := fmt.Sprintf("/transfers")

			w := httptest.NewRecorder()
			r := httptest.NewRequest(http.MethodGet, path, nil)

			r.Header.Add("Authorization", bearer)

			http.HandlerFunc(handler.ListTransferHandler).ServeHTTP(w, r)

			if w.Code != tt.wantStatusCode {
				t.Errorf("%s: got %v, want %v", tt.Name, w.Code, tt.wantStatusCode)
			}
		})

	}

}
