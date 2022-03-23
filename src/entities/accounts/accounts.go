package accounts

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/src/db"
	"github.com/rodolfoalvesg/api-banking/api/src/models"
	"github.com/rodolfoalvesg/api-banking/api/src/security"
)

var (
	ErrInvalidPassword = errors.New("A senha nao atende aos requisitos")
	ErrInvalidId       = errors.New("A senha nao atende aos requisitos")
)

// CreateAccount,
func CreateAccount(acc models.Account) ([]models.Account, error) {

	if (len(acc.Secret) < 8) || (acc.Secret == "") {
		return nil, ErrInvalidPassword
	}

	passwdHash, err := security.SecurityHash(acc.Secret) //Cria um hash da senha passada
	if err != nil {
		return nil, err
	}

	acc.Secret = string(passwdHash)  //account.Secret, atribui ao campo de senha do modelo o HASH
	acc.Id = uuid.New().String()     //account.Id, cria um id único e atribui ao campo Id
	acc.CreatedAt = time.Now().UTC() //account.CreatedAt, data e hora
	modelAccount := &db.FieldsToMethodsDB{
		Accounts: acc,
	}

	modelAccount.AddedAccount()

	modelList := &db.FieldsToMethodsDB{}
	data, err := modelList.ShowAccounts()
	if err != nil {
		return nil, err
	}

	return data, nil
}

func ShowBalance(accId string) (int, error) {

	if len(accId) == 0 {
		return 0, ErrInvalidId
	}

	modelListId := &db.FieldsToMethodsDB{
		Id: accId,
	}

	accountPerson, err := modelListId.ShowBalanceId()
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
