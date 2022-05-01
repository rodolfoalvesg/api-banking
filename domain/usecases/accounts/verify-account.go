package account

import (
	"context"
	"errors"
)

var ErrConflictCPF = errors.New("CPF already exists")

//VerifyAccount, caso de uso para verificar a existÊncia de uma conta
func (u Usecase) VerifyAccount(ctx context.Context, accCPF string) error {

	account, _ := u.repo.ListAccountsByCPF(ctx, accCPF)
	if account.CPF == accCPF {
		return ErrConflictCPF
	}

	return nil

}
