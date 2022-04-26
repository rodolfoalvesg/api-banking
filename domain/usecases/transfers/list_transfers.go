package transfer

import (
	"context"
	"fmt"

	"github.com/rodolfoalvesg/api-banking/api/domain/entities/transfers"
)

//ShowTransfers, caso de uso para exibir todas as transferências realizadas por um usuário
func (u UsecaseTransfers) ShowTransfers(ctx context.Context, accID string) ([]transfers.Transfer, error) {
	TransferList, err := u.repo.ListAllTransfers(ctx, accID)
	if err != nil {
		return []transfers.Transfer{}, fmt.Errorf("Error showing transfers: %v", err)
	}
	return TransferList, nil
}
