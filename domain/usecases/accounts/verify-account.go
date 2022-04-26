package account

import (
	"context"
	"errors"
)

var ErrConflictCPF = errors.New("CPF already exists")

//ShowBalance, caso de uso para exibir saldo de uma conta através do ID
func (u Usecase) VerifyAccount(ctx context.Context, accCPF string) error {

	account, _ := u.repo.ListAccountsByCPF(ctx, accCPF)
	if account.CPF == accCPF {
		return ErrConflictCPF
	}

	return nil

}
