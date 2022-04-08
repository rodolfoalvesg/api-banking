package account

import (
	"context"
	"fmt"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
	usecase "github.com/rodolfoalvesg/api-banking/api/domain/usecases"
)

type Usecase struct {
	repo usecase.Repository
}

//NewUsecaseAccount, cria repositório com metódos das entidades
func NewUsecase(repo usecase.Repository) *Usecase {
	return &Usecase{
		repo: repo,
	}

}

func (u Usecase) CreateAccount(ctx context.Context, account accounts.Account) (uuid.UUID, error) {

	acc, err := accounts.ValidatePasswdHash(ctx, account)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Failed to validate password hash: %w", err)
	}

	account.Secret = string(acc)

	accID, err := u.repo.SaveAccount(ctx, account)
	if err != nil {
		return uuid.UUID{}, fmt.Errorf("Failed to saving account: %w", err)
	}

	return accID, nil
}
