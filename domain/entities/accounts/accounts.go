package accounts

import (
	"context"
	"errors"
	"time"

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

// CreateAccount, verifica as regras da conta
func ValidatePasswordHash(ctx context.Context, account Account) (string, error) {

	// Analisa se a senha atende os critérios
	if len(account.Secret) < 8 {
		return "", errors.New("A senha nao atende aos requisitos")
	}

	//Cria um hash da senha passada
	passwdHash, err := security.SecurityHash(account.Secret)
	if err != nil {
		return "", errors.New("Não foi possível criar o HASH")
	}

	account.Secret = string(passwdHash) //account.Secret, atribui ao campo de senha do modelo o HASH

	return account.Secret, nil
}

/*func (a Account) ShowBalance(accID string) (int, error) {

	if len(accID) == 0 {
		return 0, errors.New("É preciso um ID válido como parâmetro")
	}

	modelListId := &db.Database{
		ID: accID,
	}

	accountPerson, err := modelListId.ShowBalanceID()
	if err != nil {
		return 0, err
	}

	responseAccount := &db.Database{
		Balance: accountPerson.Balance,
	}

	return responseAccount.Balance, nil
}

func (a Account) ShowListAccounts() ([]Account, error) {
	modelShowAccounts := &db.Database{}
	accountLits, err := modelShowAccounts.ShowAccounts()
	if err != nil {
		return nil, err
	}
	return accountLits, err
}
*/
