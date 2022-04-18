package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/controllers"
)

func RouterTransfers(c controllers.Controller) []Router {
	return []Router{
		{
			URI:            "/transfers",
			Method:         http.MethodPost,
			Function:       nil,
			Authentication: false,
		},
		{
			URI:            "/transfers",
			Method:         http.MethodGet,
			Function:       nil,
			Authentication: false,
		},
	}
}
