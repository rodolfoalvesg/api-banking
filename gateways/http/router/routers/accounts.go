package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/controllers"
)

func RouterAccounts(c controllers.Controller) []Router {

	return []Router{
		{
			URI:            "/accounts",
			Method:         http.MethodPost,
			Function:       c.CreateAccount,
			Authentication: false,
		},
		{
			URI:            "/accounts/{account_id}/balance",
			Method:         http.MethodGet,
			Function:       nil,
			Authentication: false,
		},
		{
			URI:            "/accounts",
			Method:         http.MethodGet,
			Function:       nil,
			Authentication: false,
		},
	}
}
