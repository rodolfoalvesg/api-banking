package routers

import (
	"net/http"
)

var transfersRouters = []Router{
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
