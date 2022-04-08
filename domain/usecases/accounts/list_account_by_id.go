package account

import (
	"context"
	"errors"
)

func (u Usecase) ShowBalance(ctx context.Context, accID string) (int, error) {

	if len(accID) == 0 {
		return 0, errors.New("É preciso um ID válido como parâmetro")
	}

	accBalance, err := u.repo.ShowBalanceID(ctx, accID)
	if err != nil {
		return 0, err
	}

	return accBalance, nil

}
