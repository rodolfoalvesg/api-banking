package transfer

import (
	"context"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

//ShowTransfers, caso de uso para exibir todas as transferências realizadas por um usuário
func (u UsecaseTransfers) ShowTransfers(ctx context.Context, accID string) ([]transfers.Transfer, error) {
	transferList, err := u.repo.ListAllTransfers(ctx, accID)
	if err != nil {
		return nil, err
	}
	return transferList, nil
}
