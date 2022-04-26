package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/controllers"
)

//RouterTransfers, rotas da entidade transferência
func RouterTransfers(c controllers.Controller) []Router {
	return []Router{
		{
			URI:            "/transfers",
			Method:         http.MethodPost,
			Function:       c.CreateTransferHandler,
			Authentication: true,
		},
		{
			URI:            "/transfers",
			Method:         http.MethodGet,
			Function:       c.ListTransferHandler,
			Authentication: true,
		},
	}
}
