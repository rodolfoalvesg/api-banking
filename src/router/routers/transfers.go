package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/src/controllers"
)

var transfersRouters = []Router{
	{
		URI:            "/transfers",
		Method:         http.MethodPost,
		Function:       controllers.HandleTransfers,
		Authentication: false,
	},
	{
		URI:            "/transfers",
		Method:         http.MethodGet,
		Function:       controllers.ShowTransfers,
		Authentication: false,
	},
}
