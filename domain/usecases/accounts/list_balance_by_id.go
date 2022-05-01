package account

import (
	"context"

	"github.com/google/uuid"
)

//ShowBalance, caso de uso para exibir saldo de uma conta através do ID
func (u Usecase) ShowBalance(ctx context.Context, accID uuid.UUID) (int, error) {

	accBalance, err := u.repo.ListBalanceByID(ctx, accID)
	if err != nil {
		return 0, err
	}

	return accBalance, nil

}
