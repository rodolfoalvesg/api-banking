package accounts

import (
	"context"
	"errors"
	"time"

	"github.com/rodolfoalvesg/api-banking/api/common/security"
)

type Account struct {
	ID        string    `json:"id"`
	Name      string    `json:"name"`
	CPF       string    `json:"cpf"`
	Secret    string    `json:"secret"`
	Balance   int       `json:"balance"`
	CreatedAt time.Time `json:"created_at,omitempty"`
}

type Balance struct {
	ID      string
	Balance int
}

var (
	ErrCreateHash       = errors.New("creating hash")
	ErrInvalidPasswd    = errors.New("password must be at least 8 characters long")
	ErrInvalidLengthCPF = errors.New("CPF must be exactly 11 characters long")
)

// CreateAccount, verifica as regras da conta
func GeneratePasswdHash(ctx context.Context, account Account) ([]byte, error) {

	//Cria um hash da senha passada
	passwdHash, err := security.SecurityHash(account.Secret)
	if err != nil {
		return nil, ErrCreateHash
	}

	return passwdHash, nil
}

func ValidateCreateAccountData(account Account) error {

	// Analisa se a senha atende os critérios
	if len(account.Secret) < 8 {
		return ErrInvalidPasswd
	}

	if len(account.CPF) != 11 {
		return ErrInvalidLengthCPF
	}

	return nil
}
