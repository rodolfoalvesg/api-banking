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
func GeneratePasswdHash(ctx context.Context, account Account) ([]byte, error) {

	// Analisa se a senha atende os critérios
	if len(account.Secret) < 8 {
		return nil, errors.New("A senha nao atende aos requisitos")
	}

	//Cria um hash da senha passada
	passwdHash, err := security.SecurityHash(account.Secret)
	if err != nil {
		return nil, errors.New("Não foi possível criar o HASH")
	}

	return passwdHash, nil
}
