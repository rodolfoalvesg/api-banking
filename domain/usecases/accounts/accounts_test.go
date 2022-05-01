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

//TestCreateAccount. teste de caso de uso para registro e criação de uma conta
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

//TestShowBalance, teste de caso de uso para exibir saldo de uma conta
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
		err         error
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
					return nil, errors.New("Erro")
				},
			},
			want: nil,
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

//TestVerifyAccount, teste de caso para verificar a existência de uma conta
func TestVerifyAccount(t *testing.T) {
	t.Parallel()

	type TestCase struct {
		name        string
		accCPF      string
		accountMock accounts.AccountMock
		want        accounts.Account
		err         error
	}

	myAccountFake := accounts.Account{

		ID:        uuid.New().String(),
		Name:      "Rodolfo",
		CPF:       "12345678900",
		Secret:    "12345678",
		Balance:   25550,
		CreatedAt: time.Now().UTC(),
	}

	testCase := []TestCase{
		{
			name:   "Pass if the account doesn't exist",
			accCPF: "12345678911",
			accountMock: accounts.AccountMock{
				ListAccountsByCPFAcc: func(string) (accounts.Account, error) {
					return myAccountFake, nil
				},
			},
			err: nil,
		},
		{
			name:   "Does not pass if the account exists",
			accCPF: "12345678900",
			accountMock: accounts.AccountMock{
				ListAccountsByCPFAcc: func(string) (accounts.Account, error) {
					return myAccountFake, ErrConflictCPF
				},
			},
			err: ErrConflictCPF,
		},
	}
	for _, tc := range testCase {
		tt := tc
		t.Run(tt.name, func(t *testing.T) {
			usecase := NewUsecase(accounts.AccountMock{
				ListAccountsByCPFAcc: tt.accountMock.ListAccountsByCPFAcc,
			})

			err := usecase.VerifyAccount(context.Background(), tt.accCPF)

			if err != tt.err {
				t.Errorf("%s, want %v, got %v", tt.name, tt.err, err)
			}
		})
	}
}

//TestUpdateAccount, teste de caso para atualizar saldo de uma conta
func TestUpdateAccount(t *testing.T) {
}
