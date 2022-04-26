package routers

import (
	"net/http"

	"github.com/rodolfoalvesg/api-banking/api/controllers"
)

//RouterLogin, rota de login
func RouterLogin(c controllers.Controller) Router {
	return Router{
		URI:            "/login",
		Method:         http.MethodPost,
		Function:       c.LoginHandler,
		Authentication: false,
	}

}
