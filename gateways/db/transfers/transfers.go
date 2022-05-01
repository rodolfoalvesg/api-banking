package transfer_db

import (
	"context"
	"errors"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
	transfer "github.com/rodolfoalvesg/api-banking/api/domain/usecases/transfers"
)

var (
	ErrIDExists = errors.New("ID already exists")
)

var _ transfer.RepositoryTransfers = (*DatabaseTransfer)(nil)

type DatabaseTransfer struct {
	dataTransfers map[uuid.UUID]transfers.Transfer
}

// NewRepositoryTransfer, cria um novo repositório do banco
func NewRepositoryTransfer() *DatabaseTransfer {
	return &DatabaseTransfer{
		dataTransfers: make(map[uuid.UUID]transfers.Transfer),
	}
}

// SaveTransfer, salva a transferência entre contas
func (dt *DatabaseTransfer) SaveTransfer(ctx context.Context, transfer transfers.Transfer) (uuid.UUID, error) {

	var uuID = uuid.New()

	if _, ok := dt.dataTransfers[uuID]; ok {
		return uuid.UUID{}, ErrIDExists
	}

	transfer.ID = uuID.String()
	transfer.CreatedAt = time.Now().UTC()
	dt.dataTransfers[uuID] = transfer

	return uuID, nil
}

// ListAllTransfers, Lista todas as transferências de um usuário
func (dt *DatabaseTransfer) ListAllTransfers(_ context.Context, accID string) ([]transfers.Transfer, error) {

	listTransfers := []transfers.Transfer{}

	for _, transfer := range dt.dataTransfers {
		if transfer.AccountOriginID == accID {
			listTransfers = append(listTransfers, transfer)
		}
	}

	return listTransfers, nil
}
