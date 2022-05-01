package account

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/common/security"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

//TestNewLogin, teste caso de uso para login
func TestNewLogin(t *testing.T) {

	myLoginFakeA := Login{
		CPF:    "12345678900",
		Secret: "12345678",
	}

	myLoginFakeB := Login{
		CPF:    "12345678900",
		Secret: "123456789545454",
	}

	myAccountFake := accounts.Account{
		ID:        uuid.New().String(),
		Name:      "Fulano de Tal",
		CPF:       "12345678900",
		Secret:    "12345678",
		Balance:   525000,
		CreatedAt: time.Now().UTC(),
	}

	token, _ := security.CreateToken(myAccountFake.ID)
	hash, _ := accounts.GeneratePasswdHash(context.Background(), myAccountFake)
	myAccountFake.Secret = string(hash)

	t.Run("Successfully generated token", func(t *testing.T) {
		t.Parallel()

		accountMock := accounts.AccountMock{
			ListAccountsByCPFAcc: func(string) (accounts.Account, error) {
				return myAccountFake, nil
			},
		}

		usecase := NewUsecase(accounts.AccountMock{
			ListAccountsByCPFAcc: accountMock.ListAccountsByCPFAcc,
		})

		tk, err := usecase.NewLogin(context.Background(), myLoginFakeA)
		if err != nil {
			t.Errorf("want %v, got %v", token, tk)
		}
	})

	t.Run("Account not found", func(t *testing.T) {
		t.Parallel()

		errNotFound := errors.New("Account not found")

		accountMock := accounts.AccountMock{
			ListAccountsByCPFAcc: func(string) (accounts.Account, error) {
				return accounts.Account{}, errNotFound
			},
		}

		usecase := NewUsecase(accounts.AccountMock{
			ListAccountsByCPFAcc: accountMock.ListAccountsByCPFAcc,
		})

		_, err := usecase.NewLogin(context.Background(), myLoginFakeA)
		if err != errNotFound {
			t.Errorf("want %v, got %v", errNotFound, err)
		}
	})

	t.Run("Invalid Passwd", func(t *testing.T) {
		t.Parallel()

		accountMock := accounts.AccountMock{
			ListAccountsByCPFAcc: func(string) (accounts.Account, error) {
				return myAccountFake, nil
			},
		}

		usecase := NewUsecase(accounts.AccountMock{
			ListAccountsByCPFAcc: accountMock.ListAccountsByCPFAcc,
		})

		_, err := usecase.NewLogin(context.Background(), myLoginFakeB)
		if !errors.Is(err, ErrInvalidPasswd) {
			t.Errorf("want %v, got %v", ErrInvalidPasswd, err)
		}
	})

}
