package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/src/controllers"
)

var accountRouters = []Router{
	{
		URI:            "/accounts",
		Method:         http.MethodPost,
		Function:       controllers.Control.HandleCreateAccount,
		Authentication: false,
	},
	{
		URI:            "/accounts/{account_id}/balance",
		Method:         http.MethodGet,
		Function:       controllers.Control.ShowBalance,
		Authentication: false,
	},
	{
		URI:            "/accounts",
		Method:         http.MethodGet,
		Function:       controllers.Control.ShowAccounts,
		Authentication: false,
	},
}
