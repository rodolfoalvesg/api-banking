package transfer_db

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
	transfer "github.com/rodolfoalvesg/api-banking/api/domain/usecases/transfers"
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
		return uuid.UUID{}, fmt.Errorf("ID already exists!")
	}

	transfer.ID = uuID.String()
	transfer.Created_At = time.Now().UTC()
	dt.dataTransfers[uuID] = transfer

	return uuID, nil
}

// ListAllTransfers, Lista todas as transferências de um usuário
func (dt *DatabaseTransfer) ListAllTransfers(_ context.Context, accID string) ([]transfers.Transfer, error) {

	listTransfers := []transfers.Transfer{}

	for _, transfer := range dt.dataTransfers {
		if transfer.Account_origin_ID == accID {
			listTransfers = append(listTransfers, transfer)
		}
	}

	return listTransfers, nil
}
