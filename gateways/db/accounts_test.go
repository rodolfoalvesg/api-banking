package db

import (
	"fmt"
	"testing"
	"time"

	"github.com/rodolfoalvesg/api-banking/api/domain/models"
)

func TestShowBalanceId(t *testing.T) {

	t.Run("Teste: Falha", func(t *testing.T) {
		accFakeID := FieldsToMethodsDB{}
		_, err := accFakeID.ShowBalanceID()

		want := fmt.Errorf("Conta não localizada")

		if err == want {
			t.Errorf("got %v, want %v", err, want)
		}
	})

	t.Run("Teste: OK", func(t *testing.T) {
		accFake := FieldsToMethodsDB{Accounts: models.Account{
			ID:        "erwerwer74rwe4r1rwe4r58",
			Name:      "Fulano de tal",
			CPF:       "01254565485",
			Secret:    "dsdgfgf544",
			Balance:   2000,
			CreatedAt: time.Now(),
		}}

		accFake.AddedAccount()

		accFakeID := FieldsToMethodsDB{ID: "erwerwer74rwe4r1rwe4r58"}
		balance, _ := accFakeID.ShowBalanceID()
		want := 2000

		if balance.Balance != want {
			t.Errorf("got %v, want %v", balance.Balance, want)
		}
	})

}

func TestShowAccounts(t *testing.T) {
	accListAccountsFake := FieldsToMethodsDB{}

	listAcc, _ := accListAccountsFake.ShowAccounts()

	want := 0

	if len(listAcc) != want {
		t.Errorf("got %v, want %v", len(listAcc), want)
	}
}
