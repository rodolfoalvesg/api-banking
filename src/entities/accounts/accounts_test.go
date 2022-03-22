package accounts

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rodolfoalvesg/api-banking/api/src/models"
)

func TestCreateAccount(t *testing.T) {
	var accountFakeA = models.Account{
		Name:    "Rodolfo Alves",
		Cpf:     "12345678900",
		Secret:  "01470258",
		Balance: 25000,
	}
	var accountFakeB = models.Account{
		Secret: "01478",
	}

	t.Run("Criação de conta: OK", func(t *testing.T) {
		got, _ := CreateAccount(accountFakeA)

		want := []models.Account{
			got[0],
			got[1],
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("Criação de conta: Falha na senha", func(t *testing.T) {
		_, err := CreateAccount(accountFakeB)
		want := fmt.Errorf("A senha não atende aos requisitos")
		if err == want {
			t.Errorf("got %v, want %v", err, want)
		}
	})
}

func TestShowBalance(t *testing.T) {

	t.Run("Listando saldo: Id preenchido", func(t *testing.T) {
		accIdBalance := "kgdf4gf4gfdgf554gsfag4g"
		got, _ := ShowBalance(accIdBalance)
		want := 5000
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Listando saldo: Id vazio", func(t *testing.T) {
		accIdBalance := ""
		_, err := ShowBalance(accIdBalance)
		want := fmt.Errorf("O id não pode ser vazio")
		if err == want {
			t.Errorf("got %v, want %v", err, want)
		}
	})

}

func TestShowListAccounts(t *testing.T) {
	t.Run("Listando todas as contas", func(t *testing.T) {

		got, _ := ShowListAccounts()
		want := []models.Account{
			got[0],
			got[1],
		}

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

}
