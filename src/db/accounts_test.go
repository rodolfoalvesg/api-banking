package db

import (
	"testing"
	"time"

	"github.com/rodolfoalvesg/api-banking/api/src/models"
)

func TestShowBalanceId(t *testing.T) {
	accFake := FieldsToMethodsDB{Accounts: models.Account{
		Id:        "erwerwer74rwe4r1rwe4r58",
		Name:      "Fulano de tal",
		Cpf:       "01254565485",
		Secret:    "dsdgfgf544",
		Balance:   2000,
		CreatedAt: time.Now(),
	}}

	accFake.AddedAccount()

	accFakeID := FieldsToMethodsDB{Id: "erwerwer74rwe4r1rwe4r58"}
	balance, _ := accFakeID.ShowBalanceId()
	want := 2000

	if balance.Balance != want {
		t.Errorf("got %v, want %v", balance.Balance, want)
	}
}

func TestShowAccounts(t *testing.T) {
	accListAccountsFake := FieldsToMethodsDB{}

	listAcc, _ := accListAccountsFake.ShowAccounts()

	want := 0

	if len(listAcc) != want {
		t.Errorf("got %v, want %v", len(listAcc), want)
	}
}
