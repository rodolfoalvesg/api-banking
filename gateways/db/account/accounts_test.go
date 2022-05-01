package account_db

import (
	"context"
	"reflect"
	"testing"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

//TestSaveAccount, teste de método para salvar conta no db
func TestSaveAccount(t *testing.T) {
	t.Parallel()

	t.Run("Save Account", func(t *testing.T) {
		myAccountFake := accounts.Account{
			Name:    "Fulano de Tal",
			CPF:     "12345678900",
			Secret:  "12345678",
			Balance: 2500,
		}

		repository := NewRepository()
		got, _ := repository.SaveAccount(context.Background(), myAccountFake)
		if got == (uuid.Nil) {
			t.Errorf("got %v, want !=  %v", got, uuid.UUID{})
		}
	})

}

//TestListBalanceByID, teste de método para exibir saldo de uma conta do db
func TestListBalanceByID(t *testing.T) {
	t.Parallel()
	repository := NewRepository()
	t.Run("Balance Listed", func(t *testing.T) {

		myAccountFake := accounts.Account{
			Name:    "Fulano de Tal",
			CPF:     "12345678900",
			Secret:  "12345678",
			Balance: 2500,
		}

		accID, _ := repository.SaveAccount(context.Background(), myAccountFake)
		got, err := repository.ListBalanceByID(context.Background(), accID)
		if err != nil {
			t.Errorf("got %v, want  %v", err, got)
		}
	})

	t.Run("Balance doesn't Listed", func(t *testing.T) {

		accID := uuid.New()
		want := 0
		got, _ := repository.ListBalanceByID(context.Background(), accID)
		if got != want {
			t.Errorf("got %v, want  %v", got, want)
		}
	})
}

//TestListAllAccounts, teste de método para exibir todas as contas db
func TestListAllAccounts(t *testing.T) {
	t.Parallel()
	repository := NewRepository()

	t.Run("Accounts Listed", func(t *testing.T) {
		myAccountFake := accounts.Account{
			Name:    "Fulano de Tal",
			CPF:     "12345678900",
			Secret:  "12345678",
			Balance: 2500,
		}

		_, err := repository.SaveAccount(context.Background(), myAccountFake)
		if err != nil {
			t.Errorf("Save Account error")
		}

		got, _ := repository.ListAllAccounts(context.Background())
		want := 1
		if !reflect.DeepEqual(len(got), want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

//TestListAccountsByCPF, teste de método para exibir uma conta do db pelo CPF
func TestListAccountsByCPF(t *testing.T) {
	t.Parallel()
	repository := NewRepository()

	t.Run("Accounts Listed", func(t *testing.T) {
		myAccountFake := accounts.Account{
			Name:    "Fulano de Tal",
			CPF:     "12345678900",
			Secret:  "12345678",
			Balance: 2500,
		}

		_, err := repository.SaveAccount(context.Background(), myAccountFake)
		if err != nil {
			t.Errorf("Save Account error")
		}

		acc, err := repository.ListAccountsByCPF(context.Background(), myAccountFake.CPF)
		if err != nil {
			t.Errorf("got %v, want %v", err, acc)
		}
	})

	t.Run("Accounts doesn't Listed", func(t *testing.T) {

		want := accounts.Account{}
		got, _ := repository.ListAccountsByCPF(context.Background(), "12345678911")
		if want != got {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}

//TestListAccountByID, teste de método para exibir uma conta do db pelo ID
func TestListAccountByID(t *testing.T) {
	t.Parallel()
	myAccountFake := accounts.Account{
		Name:    "Fulano de Tal",
		CPF:     "12345678900",
		Secret:  "12345678",
		Balance: 2500,
	}

	repository := NewRepository()

	t.Run("account ID Listed", func(t *testing.T) {
		accID, err := repository.SaveAccount(context.Background(), myAccountFake)

		if err != nil {
			t.Errorf("ERROR: Save Account")
		}

		_, err = repository.ListAccountByID(context.Background(), accID.String())

		if err != nil {
			t.Errorf("Erro: %v", err)
		}

	})

	t.Run("Error invalid ID", func(t *testing.T) {
		accID := ""

		got, _ := repository.ListAccountByID(context.Background(), accID)
		want := accounts.Account{}
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})

	t.Run("Error Account Not Found", func(t *testing.T) {
		accID := uuid.New()

		got, _ := repository.ListAccountByID(context.Background(), accID.String())
		want := accounts.Account{}
		if got != want {
			t.Errorf("want %v, got %v", want, got)
		}
	})
}

//TestUpdatedAccount, teste de método para atualizar saldo de um usuário
func TestUpdatedAccount(t *testing.T) {
	t.Parallel()

	repository := NewRepository()

	t.Run("Error invalid ID", func(t *testing.T) {
		balance := accounts.Balance{
			ID:      "",
			Balance: 250,
		}

		err := repository.UpdatedAccount(context.Background(), balance)
		if err == nil {
			t.Errorf("Error, %v", err)
		}
	})

	t.Run("Error internal error", func(t *testing.T) {
		balance := accounts.Balance{
			ID:      uuid.New().String(),
			Balance: 250,
		}

		err := repository.UpdatedAccount(context.Background(), balance)
		if err == nil {
			t.Errorf("Error, %v", err)
		}
	})

	t.Run("Error internal error", func(t *testing.T) {

		myAccountFake := accounts.Account{
			Name:    "Fulano de Tal",
			CPF:     "12345678900",
			Secret:  "12345678",
			Balance: 2500,
		}

		accID, err := repository.SaveAccount(context.Background(), myAccountFake)
		if err != nil {
			t.Errorf("Save Account error")
		}

		balance := accounts.Balance{
			ID:      accID.String(),
			Balance: 250,
		}

		err = repository.UpdatedAccount(context.Background(), balance)
		if err != nil {
			t.Errorf("Error, %v", err)
		}
	})
}
