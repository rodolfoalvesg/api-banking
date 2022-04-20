package transfer

import (
	"context"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

type RepositoryTransfers interface {
	SaveTransfer(context.Context, transfers.Transfer) (uuid.UUID, error)
}

type UsecaseTransfers struct {
	repo RepositoryTransfers
}

//NewUsecaseAccount, cria repositório com metódos das entidades
func NewUsecaseTransfers(repo RepositoryTransfers) *UsecaseTransfers {
	return &UsecaseTransfers{
		repo: repo,
	}

}
