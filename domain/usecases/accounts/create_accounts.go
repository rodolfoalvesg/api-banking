package account

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

func (u Usecase) CreateAccount(ctx context.Context, account accounts.Account) (uuid.UUID, error) {

	acc, err := accounts.GeneratePasswdHash(ctx, account)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Failed to validate password hash: %w", err)
	}

	account.Secret = string(acc)

	fmt.Println("Aqui Usecase")

	fmt.Println(account)
	accID, err := u.repo.SaveAccount(ctx, account)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Failed to saving account: %w", err)
	}

	return accID, nil
}
