package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/controllers"
)

var controller = controllers.NewController(&controllers.Controller{})

var accountRouters = []Router{
	{
		URI:            "/accounts",
		Method:         http.MethodPost,
		Function:       controller.HandlerCreateAccount,
		Authentication: false,
	},
	{
		URI:            "/accounts/{account_id}/balance",
		Method:         http.MethodGet,
		Function:       controller.HandlerShowBalance,
		Authentication: false,
	},
	{
		URI:            "/accounts",
		Method:         http.MethodGet,
		Function:       controller.HandlerShowAccounts,
		Authentication: false,
	},
}
