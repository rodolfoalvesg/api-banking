package account

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/accounts"
)

type AccountMock struct {
	CreateM func(acc accounts.Account) (uuid.UUID, error)
}

func (m *AccountMock) CreateAccount(ctx context.Context, account accounts.Account) (uuid.UUID, error) {
	return m.CreateM(account)
}
