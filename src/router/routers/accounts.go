package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/src/controllers"
)

var accountRouters = []Router{
	{
		URI:            "/accounts",
		Method:         http.MethodPost,
		Function:       controllers.CreateAccount,
		Authentication: false,
	},
	{
		URI:            "/accounts/{account_id}/ballance",
		Method:         http.MethodGet,
		Function:       controllers.ShowBallance,
		Authentication: false,
	},
	{
		URI:            "/accounts",
		Method:         http.MethodGet,
		Function:       controllers.ShowAccounts,
		Authentication: false,
	},
}
