package account

import (
	"context"
	"testing"

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
	}

	for _, tc := range testCase {
		tt := tc
		t.Run(tt.Name, func(t *testing.T) {
			t.Parallel()

			usecase := NewUsecase(accounts.AccountMock{
				SaveAcc: tt.accountMock.SaveAcc,
			})

			got, err := usecase.CreateAccount(context.Background(), tt.accFake)
			if err != nil {
				t.Errorf("Error creating account %s", err)
			}

			if got != tt.want {
				t.Errorf("%s, want %v, got %v", tt.Name, tt.want, got)
			}
		})
	}
}
