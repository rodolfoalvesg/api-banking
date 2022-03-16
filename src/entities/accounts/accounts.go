package accounts

import (
	"encoding/json"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/src/db"
	"github.com/rodolfoalvesg/api-banking/api/src/models"
	"github.com/rodolfoalvesg/api-banking/api/src/security"
)

func CreateAccount(bodyRequest []byte, err error) ([]models.Account, error) {
	var account models.Account
	err = json.Unmarshal(bodyRequest, &account)
	if err != nil {
		return nil, err
	}

	passwdHash, err := security.SecurityHash(account.Secret) //Cria um hash da senha passada
	if err != nil {
		return nil, err
	}

	account.Secret = string(passwdHash)  //account.Secret, atribui ao campo de senha do modelo o HASH
	account.Id = uuid.New().String()     //account.Id, cria um id único e atribui ao campo Id
	account.CreatedAt = time.Now().UTC() //account.CreatedAt, data e hora
	modelAccount := &db.FieldsToMethodsDB{
		Accounts: account,
	}

	modelAccount.AddedAccount()

	modelList := &db.FieldsToMethodsDB{}
	data, err := modelList.ShowAccounts()
	if err != nil {
		return nil, err
	}

	return data, nil
}
