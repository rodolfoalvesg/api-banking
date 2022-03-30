package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/controllers"
	"github.com/rodolfoalvesg/api-banking/api/gateways/db"
)

var controller = controllers.NewController(&db.FieldsToMethodsDB{})

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
