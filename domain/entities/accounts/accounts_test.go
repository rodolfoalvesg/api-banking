package accounts

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/rodolfoalvesg/api-banking/api/domain/models"
)

// TestCreateAccount, teste de criação de conta
func TestCreateAccount(t *testing.T) {
	var accountFakeA = models.Account{
		Name:    "Rodolfo Alves",
		CPF:     "12345678900",
		Secret:  "01470258",
		Balance: 25000,
	}
	var accountFakeB = models.Account{
		Secret: "01478",
	}

	t.Run("Criação de conta: OK", func(t *testing.T) {
		listAccount, _ := NewCreateAccount(accountFakeA)
		got := listAccount.Balance
		want := 25000

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}

	})

	t.Run("Criação de conta: Falha na senha", func(t *testing.T) {
		_, err := NewCreateAccount(accountFakeB)
		want := fmt.Errorf("A senha não atende aos requisitos")
		if err == want {
			t.Errorf("got %v, want %v", err, want)
		}
	})
}

// TestShowBalance, teste de busca de saldo

// TestShowListAccounts, teste de listagemd e contas
func TestShowListAccounts(t *testing.T) {

	accListA, _ := ShowListAccounts()

	accFake := models.Account{
		Name:   "Rodolfo",
		CPF:    "01470225825",
		Secret: "123456123",
	}

	acc, _ := NewCreateAccount(accFake)
	accListB := []models.Account{}
	accListB = append(accListB, acc)

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
