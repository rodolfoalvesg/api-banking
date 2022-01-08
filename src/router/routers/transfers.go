package routers

import (
	"api/src/controllers"
	"net/http"
)

var transfersRouters = []Router{
	{
		URI:            "/transfers",
		Method:         http.MethodPost,
		Function:       controllers.AmountTransfers,
		Authentication: false,
	},
	{
		URI:            "/transfers",
		Method:         http.MethodGet,
		Function:       controllers.ShowTransfers,
		Authentication: false,
	},
}
