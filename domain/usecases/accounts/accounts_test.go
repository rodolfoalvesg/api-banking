package account

import (
	"context"
	"errors"
	"reflect"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

func TestCreateAccount(t *testing.T) {
	t.Parallel()

	var myFakeAccount = accounts.Account{
		Name:    "Rodolfo",
		CPF:     "12345678900",
		Secret:  "12345678",
		Balance: 15500,
	}

	type TestCaseA struct {
		Name        string
		accountMock accounts.AccountMock
		accFake     accounts.Account
		want        uuid.UUID
	}

	accID := uuid.New()
	testCase := []TestCaseA{
		{
			Name: "account created successfully",
			accountMock: accounts.AccountMock{
				SaveAcc: func(accounts.Account) (uuid.UUID, error) {
					return accID, nil
				},
			},
			accFake: myFakeAccount,
			want:    accID,
		},
		{
			Name: "account not created",
			accountMock: accounts.AccountMock{
				SaveAcc: func(accounts.Account) (uuid.UUID, error) {
					return uuid.UUID{}, errors.New("Conflict")
				},
			},
			accFake: myFakeAccount,
			want:    uuid.UUID{},
		},
	}

	for _, tc := range testCase {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			usecase := NewUsecase(accounts.AccountMock{
				SaveAcc: tt.accountMock.SaveAcc,
			})

			got, err := usecase.CreateAccount(context.Background(), tt.accFake)
			if got != tt.want && err != nil {
				t.Errorf("Error %s", err)
			}
		})
	}
}

func TestShowBalance(t *testing.T) {
	t.Parallel()

	myFakeAccount := accounts.Account{

		Balance: 15500,
	}

	accID := uuid.New()

	type TestCaseB struct {
		Name        string
		accountMock accounts.AccountMock
		accFakeID   uuid.UUID
		want        int
	}

	testCase := []TestCaseB{
		{
			Name: "Balance successfully listed",
			accountMock: accounts.AccountMock{
				ListBalanceByIDAcc: func(uuid.UUID) (int, error) {
					return myFakeAccount.Balance, nil
				},
			},
			accFakeID: accID,
			want:      15500,
		},
		{
			Name: "Error listing balance",
			accountMock: accounts.AccountMock{
				ListBalanceByIDAcc: func(uuid.UUID) (int, error) {
					return 0, errors.New("Error")
				},
			},
			accFakeID: accID,
			want:      0,
		},
	}

	for _, tc := range testCase {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			usecase := NewUsecase(accounts.AccountMock{
				ListBalanceByIDAcc: tt.accountMock.ListBalanceByIDAcc,
			})

			balance, err := usecase.ShowBalance(context.Background(), tt.accFakeID)
			if balance != tt.want && err != nil {
				t.Errorf("%s, got %v, want %v", tt.Name, balance, tt.want)

			}

		})
	}
}

//TestShowAccounts, teste de caso de uso para listar contas
func TestShowAccounts(t *testing.T) {
	t.Parallel()

	type TestCaseC struct {
		Name        string
		accountMock accounts.AccountMock
		want        []accounts.Account
	}

	myListAccounts := []accounts.Account{
		{
			ID:        uuid.New().String(),
			Name:      "Rodolfo",
			CPF:       "12345678900",
			Secret:    "12345678",
			Balance:   25550,
			CreatedAt: time.Now().UTC(),
		},
		{
			ID:        uuid.New().String(),
			Name:      "Rosangela",
			CPF:       "01472558",
			Secret:    "12345678",
			Balance:   50000,
			CreatedAt: time.Now().UTC(),
		},
	}

	testCase := []TestCaseC{
		{
			Name: "Accounts Listed",
			accountMock: accounts.AccountMock{
				ListAllAcc: func(context.Context) ([]accounts.Account, error) {
					return myListAccounts, nil
				},
			},
			want: myListAccounts,
		},
		{
			Name: "Error: Accounts not Listed",
			accountMock: accounts.AccountMock{
				ListAllAcc: func(context.Context) ([]accounts.Account, error) {
					return []accounts.Account{}, errors.New("Erro")
				},
			},
			want: []accounts.Account{},
		},
	}

	for _, tc := range testCase {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			usecase := NewUsecase(accounts.AccountMock{
				ListAllAcc: tt.accountMock.ListAllAcc,
			})

			listAccounts, err := usecase.ShowAccounts(context.Background())
			if !reflect.DeepEqual(listAccounts, tt.want) && err != nil {
				t.Errorf("%s, want %v, got %v", tt.Name, tt.want, listAccounts)
			}

		})
	}

}
