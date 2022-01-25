package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/src/controllers"
)

var accountRouters = []Router{
	{
		URI:            "/accounts",
		Method:         http.MethodPost,
		Function:       controllers.RouteMethods.HandleCreateAccount,
		Authentication: false,
	},
	{
		URI:            "/accounts/{account_id}/balance",
		Method:         http.MethodGet,
		Function:       controllers.ShowBalance,
		Authentication: false,
	},
	{
		URI:            "/accounts",
		Method:         http.MethodGet,
		Function:       controllers.ShowAccounts,
		Authentication: false,
	},
}
