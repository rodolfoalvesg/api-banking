package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/controllers"
)

//RouterAccounts, rotas da entidade Accounts
func RouterAccounts(c controllers.Controller) []Router {

	return []Router{
		{
			URI:            "/accounts",
			Method:         http.MethodPost,
			Function:       c.CreateAccountHandler,
			Authentication: false,
		},
		{
			URI:            "/accounts/{account_id}/balance",
			Method:         http.MethodGet,
			Function:       c.ShowBalanceHandler,
			Authentication: false,
		},
		{
			URI:            "/accounts",
			Method:         http.MethodGet,
			Function:       c.ShowAccountsHandler,
			Authentication: false,
		},
	}
}
