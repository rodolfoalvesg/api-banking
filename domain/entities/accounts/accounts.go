package accounts

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/models"
	"github.com/rodolfoalvesg/api-banking/api/gateways/db"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/security"
)

var (
	ErrInvalidPassword = errors.New("A senha nao atende aos requisitos")
	ErrInvalidID       = errors.New("É preciso um ID válido como parâmetro")
)

// CreateAccount, verifica as regras da conta
func CreateNewAccount(ctx context.Context, acc models.Account) (models.Account, error) {

	// Analisa se a senha atende os critérios
	if len(acc.Secret) < 8 {
		return models.Account{}, ErrInvalidPassword
	}

	//Cria um hash da senha passada
	passwdHash, err := security.SecurityHash(acc.Secret)
	if err != nil {
		return models.Account{}, err
	}

	acc.Secret = string(passwdHash)  //account.Secret, atribui ao campo de senha do modelo o HASH
	acc.ID = uuid.New().String()     //account.Id, cria um id único e atribui ao campo Id
	acc.CreatedAt = time.Now().UTC() //account.CreatedAt, data e hora

	return acc, nil
}

func ShowBalance(accID string) (int, error) {

	if len(accID) == 0 {
		return 0, ErrInvalidID
	}

	modelListId := &db.FieldsToMethodsDB{
		Id: accID,
	}

	accountPerson, err := modelListId.ShowBalanceID()
	if err != nil {
		return 0, err
	}

	responseAccount := &db.FieldsToMethodsDB{
		Balance: accountPerson.Balance,
	}

	return responseAccount.Balance, nil
}

func ShowListAccounts() ([]models.Account, error) {
	modelShowAccounts := &db.FieldsToMethodsDB{}
	accountLits, err := modelShowAccounts.ShowAccounts()
	if err != nil {
		return nil, err
	}
	return accountLits, err
}
