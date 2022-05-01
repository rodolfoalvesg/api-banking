package account

import (
	"context"
	"errors"

	"github.com/rodolfoalvesg/api-banking/api/common/security"
)

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}

var (
	ErrInvalidPasswd = errors.New("validating password")
	ErrCreateToken   = errors.New("creating token")
)

//NewLogin, caso de uso para login de usuário
func (u Usecase) NewLogin(ctx context.Context, l Login) (string, error) {

	account, err := u.repo.ListAccountsByCPF(ctx, l.CPF)
	if err != nil {
		return "", err
	}

	err = security.VerifyPasswd(account.Secret, l.Secret)
	if err != nil {
		return "", ErrInvalidPasswd
	}

	token, err := security.CreateToken(account.ID)
	if err != nil {
		return "", ErrCreateToken
	}

	return token, nil
}
