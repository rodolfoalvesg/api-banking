package account

import (
	"context"
	"errors"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

var (
	ErrInsufficientFunds = errors.New("insufficient funds")
)

//UpdateAccount, caso de uso para atualização de saldo após uma transferência
func (u Usecase) UpdateAccount(ctx context.Context, t transfers.Transfer) error {
	originAccount, err := u.repo.ListAccountByID(ctx, t.AccountOriginID)
	if err != nil {
		return err
	}

	if originAccount.Balance == 0 || originAccount.Balance < int(t.Amount) {
		return ErrInsufficientFunds
	}

	destinationAccount, err := u.repo.ListAccountByID(ctx, t.AccountDestinationID)
	if err != nil {
		return err
	}

	originAccount.Balance -= int(t.Amount)
	destinationAccount.Balance += int(t.Amount)

	origin := accounts.Balance{
		ID:      originAccount.ID,
		Balance: originAccount.Balance,
	}

	err = u.repo.UpdatedAccount(ctx, origin)
	if err != nil {
		return err
	}

	destination := accounts.Balance{
		ID:      destinationAccount.ID,
		Balance: destinationAccount.Balance,
	}

	err = u.repo.UpdatedAccount(ctx, destination)
	if err != nil {
		return err
	}

	return nil
}
