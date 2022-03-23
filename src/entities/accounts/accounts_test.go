package accounts

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rodolfoalvesg/api-banking/api/src/models"
)

// TestCreateAccount, teste de criação de conta
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
		listAccount, _ := CreateAccount(accountFakeA)
		got := listAccount[0].Balance
		want := 25000

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

// TestShowBalance, teste de busca de saldo
func TestShowBalance(t *testing.T) {

	accFake := models.Account{
		Id:      "dfsh15hjfg4hgfsdhgdsf",
		Secret:  "123456789",
		Balance: 5000,
	}

	accListA, _ := CreateAccount(accFake)
	accListB := models.Account{}

	t.Run("Listando saldo: Id preenchido", func(t *testing.T) {
		accIdBalance := accListA[0].Id
		got, _ := ShowBalance(accIdBalance)
		want := 5000
		if got != want {
			t.Errorf("got %v, want %v", got, want)
		}
	})

	t.Run("Listando saldo: Id vazio", func(t *testing.T) {
		accIdBalance := accListB.Id
		_, err := ShowBalance(accIdBalance)
		want := fmt.Errorf("O id não pode ser vazio")
		if err == want {
			t.Errorf("got %v, want %v", err, want)
		}
	})
}

// TestShowListAccounts, teste de listagemd e contas
func TestShowListAccounts(t *testing.T) {

	accListA, _ := ShowListAccounts()

	accFake := models.Account{
		Name:   "Rodolfo",
		Cpf:    "01470225825",
		Secret: "123456123",
	}

	accListB, _ := CreateAccount(accFake)

	testListAccounts := map[string]struct {
		acc  []models.Account
		want int
	}{
		"Sem conta": {accListA, 0},
		"Com conta": {accListB, 1},
	}

	for name, tt := range testListAccounts {
		if len(tt.acc) != tt.want {
			t.Errorf("%s: got %v, want %v", name, tt.acc, tt.want)
		}
	}

}
