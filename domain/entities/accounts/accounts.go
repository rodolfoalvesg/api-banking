package accounts

import (
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/gateways/db"
	"github.com/rodolfoalvesg/api-banking/api/gateways/http/security"
)

type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

var (
	ErrInvalidPassword = errors.New("A senha nao atende aos requisitos")
	ErrInvalidID       = errors.New("É preciso um ID válido como parâmetro")
)

// CreateAccount,
func (acc Account) CreateNewAccount() (Account, error) {

	if (len(acc.Secret) < 8) || (acc.Secret == "") {
		return Account{}, ErrInvalidPassword
	}

	passwdHash, err := security.SecurityHash(acc.Secret) //Cria um hash da senha passada
	if err != nil {
		return Account{}, err
	}

	acc.Secret = string(passwdHash)  //account.Secret, atribui ao campo de senha do modelo o HASH
	acc.ID = uuid.New().String()     //account.Id, cria um id único e atribui ao campo Id
	acc.CreatedAt = time.Now().UTC() //account.CreatedAt, data e hora

	return acc, nil
}

func NewShowBalance(accID string) (int, error) {

	modelListID := &db.FieldsToMethodsDB{
		ID: accID,
	}

	accountPerson, err := modelListID.ShowBalanceID()
	if err != nil {
		return 0, err
	}

	responseAccount := &db.FieldsToMethodsDB{
		Balance: accountPerson.Balance,
	}

	return responseAccount.Balance, nil
}

func ShowListAccounts() ([]Account, error) {
	modelShowAccounts := &db.FieldsToMethodsDB{}
	accountLits, err := modelShowAccounts.ShowAccounts()
	if err != nil {
		return nil, err
	}
	return accountLits, err
}
